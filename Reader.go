// Package alpinocorpus provides a reader and a writer for Alpino corpora
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
	"unsafe"
)

func init() {
	C.exsltRegisterAll()
}

type Entries struct {
	it           _Ctype_alpinocorpus_iter
	r            *Reader
	opened       bool
	has_contents bool
	name         string
	contents     string
}

func (it *Entries) Next() bool {
	if !it.opened {
		return false
	}
	if C.alpinocorpus_iter_end(it.r.c, it.it) != 0 {
		return false
	}
	it.name = C.GoString(C.alpinocorpus_iter_value(it.it))
	if it.has_contents {
		it.contents = C.GoString(C.alpinocorpus_iter_contents(it.r.c, it.it))
	}
	C.alpinocorpus_iter_next(it.r.c, it.it)
	return true
}

func (it *Entries) Name() string {
	return it.name
}

func (it *Entries) Contents() string {
	if it.has_contents {
		return it.contents
	}
	if it.opened && it.name != "" {
		it.contents, _ = it.r.Get(it.name)
	}
	return it.contents
}

func (it *Entries) Close() {
	if it.opened {
		C.alpinocorpus_iter_destroy(it.it)
		it.opened = false
	}
}

type Reader struct {
	corpusname string
	opened     bool
	c          _Ctype_alpinocorpus_reader
}

// NewReader() opens an Alpino corpus for reading
func NewReader(filename string) (*Reader, error) {
	recursive := false
	fi, err := os.Stat(filename)
	if err == nil && fi.Mode().IsDir() {
		i, err := filepath.Glob(filename + "/*.xml")
		if err == nil && len(i) == 0 {
			recursive = true
		}
	}

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
	return &r, nil
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
//
// Example usage:
//
//     entries, error := reader.GetAll()
//     if error != nil {
//         log.Fatal(error)
//     }
//     for entries.Next() {
//         fmt.Println(entry.Name(), ": ", entry.Contents())
//     }
//     entries.Close()
//
// Don't forget to call Close() at the end!
func (r *Reader) GetAll() (*Entries, error) {
	if e := r.isopen(); e != nil {
		return nil, e
	}
	i := C.alpinocorpus_entry_iter(r.c)
	if i == nil {
		return nil, errors.New("Unable to get iterator")
	}
	it := Entries{it: i, r: r, opened: true, has_contents: false}
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
//
// See GetAll() for an example of how to use the result
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
	it := Entries{it: i, r: r, opened: true, has_contents: false}
	return &it, nil
}

// QueryMod() gives access to all entries that match a certain query, using a stylesheet to modify the output
//
// See GetAll() for an example of how to use the result
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
		it := Entries{it: i, r: r, opened: true, has_contents: true}
		return &it, nil
	}

	i := C.alpinocorpus_query_stylesheet_iter(r.c, csQ, csS, nil, 0)
	it := Entries{it: i, r: r, opened: true, has_contents: true}
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
