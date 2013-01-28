package alpinocorpus

/*
#cgo pkg-config: alpinocorpus libexslt
#include <stdlib.h>
#include <string.h>
#include <AlpinoCorpus/capi.h>
#include <libexslt/exslt.h>
*/
import "C"

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"unsafe"
)

func init() {
	C.exsltRegisterAll()
}

type KeyValue struct {
	Key, Value string
}

type Entries struct {
	it           _Ctype_alpinocorpus_iter
	r            *Reader
	opened       bool
	has_contents bool
	interrupt    chan bool
}

func (it *Entries) Keys() <-chan string {
	ch := make(chan string)
	go func() {
	KeysLoop:
		for {
			if !it.opened {
				break
			}
			if C.alpinocorpus_iter_has_next(it.r.c, it.it) == 0 {
				break
			}
			ent := C.alpinocorpus_iter_next(it.r.c, it.it)
			key := C.GoString(C.alpinocorpus_entry_name(ent))
			C.alpinocorpus_entry_free(ent)
			select {
			case ch <- key:
			case <-it.interrupt:
				break KeysLoop
			}
		}
		it.close()
		close(ch)
	}()
	return ch
}

func (it *Entries) Values() <-chan string {
	ch := make(chan string)
	go func() {
	ValuesLoop:
		for {
			if !it.opened {
				break
			}
			if C.alpinocorpus_iter_has_next(it.r.c, it.it) == 0 {
				break
			}
			ent := C.alpinocorpus_iter_next(it.r.c, it.it)
			if it.has_contents {
				value := C.GoString(C.alpinocorpus_entry_contents(ent))
				C.alpinocorpus_entry_free(ent)
				select {
				case ch <- value:
				case <-it.interrupt:
					break ValuesLoop
				}
			} else {
				name := C.GoString(C.alpinocorpus_entry_name(ent))
				C.alpinocorpus_entry_free(ent)
				if name != "" {
					c, e := it.r.Get(name)
					if e == nil {
						select {
						case ch <- c:
						case <-it.interrupt:
							break ValuesLoop
						}
					}
				}
			}
		}
		it.close()
		close(ch)
	}()
	return ch
}

func (it *Entries) KeysValues() <-chan KeyValue {
	ch := make(chan KeyValue)
	go func() {
	KeysValuesLoop:
		for {
			var name, cont string
			if !it.opened {
				break
			}
			if C.alpinocorpus_iter_has_next(it.r.c, it.it) == 0 {
				break
			}
			ent := C.alpinocorpus_iter_next(it.r.c, it.it)
			name = C.GoString(C.alpinocorpus_entry_name(ent))
			if it.has_contents {
				cont = C.GoString(C.alpinocorpus_entry_contents(ent))
			} else {
				cont, _ = it.r.Get(name)
			}
			C.alpinocorpus_entry_free(ent)
			select {
			case ch <- KeyValue{Key: name, Value: cont}:
			case <-it.interrupt:
				break KeysValuesLoop
			}
		}
		it.close()
		close(ch)
	}()
	return ch
}

func (it *Entries) Break() {
	if it.opened {
		it.interrupt <- true
	}
}

func (it *Entries) close() {
	if it.opened {
		C.alpinocorpus_iter_destroy(it.it)
		close(it.interrupt)
		it.opened = false
	}
}

type Reader struct {
	corpusname  string
	opened      bool
	c           _Ctype_alpinocorpus_reader
	entrieslist [](*Entries)
}

func newReader(filename string, recursive bool) (*Reader, error) {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	r := Reader{corpusname: filename, opened: false}
	if recursive {
		r.c = C.alpinocorpus_open_recursive(cs)
	} else {
		r.c = C.alpinocorpus_open(cs)
	}
	if r.c == nil {
		return &r, errors.New("Unable to open corpus " + filename)
	}
	r.opened = true

	runtime.SetFinalizer(&r, (*Reader).Close)

	return &r, nil
}

// NewReader() opens an Alpino corpus for reading
//
// If filename is a directory, guess what kind of corpus it is based on content of directory
func NewReader(filename string) (*Reader, error) {
	recursive := false
	fi, err := os.Stat(filename)
	if err == nil && fi.Mode().IsDir() {
		i, err := filepath.Glob(filename + "/*.xml")
		if err == nil && len(i) == 0 {
			recursive = true
		}
	}
	return newReader(filename, recursive)
}

// NewReaderRecursive() opens an Alpino corpus for reading from a directory, recursively
func NewReaderRecursive(dirname string) (*Reader, error) {
	return newReader(dirname, true)
}

// NewReaderNonRecursive() opens an Alpino corpus for reading from a directory, non-recursively
//
// (actually, it opens any non-recursive corpus)
func NewReaderNonRecursive(dirname string) (*Reader, error) {
	return newReader(dirname, false)
}

// Name() returns the canonical name of the corpus
func (r *Reader) Name() string {
	if !r.opened {
		return ""
	}
	return C.GoString(C.alpinocorpus_name(r.c))
}

// Len() returns the number of entries in the corpus
func (r *Reader) Len() int {
	if !r.opened {
		return 0
	}
	return int(C.alpinocorpus_size(r.c))
}

// Close() closes the corpus
func (r *Reader) Close() {
	if r.opened {
		C.alpinocorpus_close(r.c)
		for _, e := range r.entrieslist {
			e.close()
		}
		r.opened = false
	}
}

// Get() returns the contents of an entry given its label
func (r *Reader) Get(entry string) (string, error) {
	if e := r.isopen(); e != nil {
		return "", e
	}

	cs := C.CString(entry)
	defer C.free(unsafe.Pointer(cs))
	s := C.alpinocorpus_read(r.c, cs)
	if s == nil {
		return "", errors.New("Entry not found: " + entry)
	}
	return C.GoString(s), nil
}

// GetAll() gives access to all entries in the corpus
func (r *Reader) GetAll() (*Entries, error) {
	if e := r.isopen(); e != nil {
		return nil, e
	}
	i := C.alpinocorpus_entry_iter(r.c)
	if i == nil {
		return nil, errors.New("Unable to get iterator")
	}
	it := Entries{it: i, r: r, opened: true, has_contents: false, interrupt: make(chan bool)}
	r.entrieslist = append(r.entrieslist, &it)
	return &it, nil
}

// ValidQuery() checks if a query is well-formed
func (r *Reader) ValidQuery(query string) bool {
	if !r.opened {
		return false
	}
	cs := C.CString(query)
	defer C.free(unsafe.Pointer(cs))
	if int(C.alpinocorpus_is_valid_query(r.c, cs)) == 0 {
		return false
	}
	return true
}

// Query() gives access to the names of all entries that match a certain query
func (r *Reader) Query(query string) (*Entries, error) {
	if e := r.isopen(); e != nil {
		return nil, e
	}

	cs := C.CString(query)
	defer C.free(unsafe.Pointer(cs))
	i := C.alpinocorpus_query_iter(r.c, cs)

	if i == nil {
		return nil, errors.New("Unable to get iterator")
	}
	it := Entries{it: i, r: r, opened: true, has_contents: false, interrupt: make(chan bool)}
	r.entrieslist = append(r.entrieslist, &it)
	return &it, nil
}

// QueryMod() gives access to all entries that match a certain query, using a stylesheet to modify the output
func (r *Reader) QueryMod(query, markerQuery, markerAttr, markerValue, stylesheet string) (*Entries, error) {
	if e := r.isopen(); e != nil {
		return nil, e
	}

	marker := (markerQuery != "" && markerAttr != "" && markerValue != "")

	if query != "" {
		if !r.ValidQuery(query) {
			return nil, errors.New("Invalid query: " + query)
		}
	}

	if query == "" && len(stylesheet) == 0 {
		return r.GetAll()
	}

	if len(stylesheet) == 0 {
		return r.Query(query)
	}

	csQ := C.CString(query)
	defer C.free(unsafe.Pointer(csQ))

	csS := C.CString(stylesheet)
	defer C.free(unsafe.Pointer(csS))

	if marker {
		if !r.ValidQuery(markerQuery) {
			return nil, errors.New("Invalid marker query: " + markerQuery)
		}
		csMQ := C.CString(markerQuery)
		csMA := C.CString(markerAttr)
		csMV := C.CString(markerValue)
		defer C.free(unsafe.Pointer(csMQ))
		defer C.free(unsafe.Pointer(csMA))
		defer C.free(unsafe.Pointer(csMV))
		i := C.alpinocorpus_query_stylesheet_marker_iter(r.c, csQ, csS, csMQ, csMA, csMV)
		it := Entries{it: i, r: r, opened: true, has_contents: true, interrupt: make(chan bool)}
		r.entrieslist = append(r.entrieslist, &it)
		return &it, nil
	}

	i := C.alpinocorpus_query_stylesheet_iter(r.c, csQ, csS, nil, 0)
	it := Entries{it: i, r: r, opened: true, has_contents: true, interrupt: make(chan bool)}
	r.entrieslist = append(r.entrieslist, &it)
	return &it, nil

}

// GetMod() returns the contents of an entry given its label, with some items marked
func (r *Reader) GetMod(entry, markerQuery, markerAttr, markerValue string) (string, error) {
	if e := r.isopen(); e != nil {
		return "", e
	}

	if !r.ValidQuery(markerQuery) {
		return "", errors.New("Invalid marker query: " + markerQuery)
	}

	cs := C.CString(entry)
	csMQ := C.CString(markerQuery)
	csMA := C.CString(markerAttr)
	csMV := C.CString(markerValue)
	defer C.free(unsafe.Pointer(cs))
	defer C.free(unsafe.Pointer(csMQ))
	defer C.free(unsafe.Pointer(csMA))
	defer C.free(unsafe.Pointer(csMV))

	s := C.alpinocorpus_read_mark_query(r.c, cs, csMQ, csMA, csMV)
	if s == nil {
		return "", errors.New("Entry not found: " + entry)
	}
	return C.GoString(s), nil
}

func (r *Reader) isopen() error {
	if r.opened {
		return nil
	}
	return errors.New("Corpus is closed: " + r.corpusname)
}
