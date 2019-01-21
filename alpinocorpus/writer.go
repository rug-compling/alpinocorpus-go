package alpinocorpus

/*
Functions that need free:
alpinocorpus_read
alpinocorpus_read_mark_queries
alpinocorpus_read_mark_query
alpinocorpus_name
alpinocorpus_write
alpinocorpus_write_corpus
*/

/*
#cgo pkg-config: alpinocorpus
#include <stdlib.h>
#include <AlpinoCorpus/capi.h>
*/
import "C"

import (
	"errors"
	"runtime"
	"sync"
	"unsafe"
)

type Writer struct {
	corpusname string
	opened     bool
	c          C.alpinocorpus_writer
	mu         sync.Mutex
}

type writerType int

const (
	Compact writerType = iota
	Dbxml
)

var (
	wType = map[writerType]string{
		Compact: "COMPACT_CORPUS_WRITER",
		Dbxml:   "DBXML_CORPUS_WRITER"}
)

// NewWriter() opens an Alpino corpus for writing.
// The corpus is of type Dbxml.
func NewWriter(filename string, overwrite bool) (*Writer, error) {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	ct := C.CString(wType[Dbxml])
	defer C.free(unsafe.Pointer(ct))
	ov := 0
	if overwrite {
		ov = 1
	}
	w := Writer{corpusname: filename, opened: false}
	w.c = C.alpinocorpus_writer_open(cs, C.int(ov), ct)
	if w.c == nil {
		return &w, errors.New("Unable to open corpus " + filename)
	}
	w.opened = true
	runtime.SetFinalizer(&w, (*Writer).Close)
	return &w, nil
}

// NewWriterType() opens an Alpino corpus for writing.
// The type of corpus is specified in the third argument.
// Currently, the only valid types are Compact and Dbxml.
func NewWriterType(filename string, overwrite bool, writertype writerType) (*Writer, error) {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	ct := C.CString(wType[writertype])
	defer C.free(unsafe.Pointer(ct))
	ov := 0
	if overwrite {
		ov = 1
	}
	w := Writer{corpusname: filename, opened: false}
	w.c = C.alpinocorpus_writer_open(cs, C.int(ov), ct)
	if w.c == nil {
		return &w, errors.New("Unable to open corpus " + filename)
	}
	w.opened = true
	runtime.SetFinalizer(&w, (*Writer).Close)
	return &w, nil
}

func (w *Writer) Write(name, contents string) error {
	if e := w.isopen(); e != nil {
		return e
	}
	csN := C.CString(name)
	csC := C.CString(contents)
	defer C.free(unsafe.Pointer(csN))
	defer C.free(unsafe.Pointer(csC))
	e := C.alpinocorpus_write(w.c, csN, csC)
	if e != nil {
		defer C.free(unsafe.Pointer(e))
		return errors.New(C.GoString(e))
	}
	return nil
}

func (w *Writer) WriteCorpus(r *Reader, failsafe bool) error {
	if e := w.isopen(); e != nil {
		return e
	}
	fs := 0
	if failsafe {
		fs = 1
	}
	e := C.alpinocorpus_write_corpus(w.c, r.c, C.int(fs))
	if e != nil {
		defer C.free(unsafe.Pointer(e))
		return errors.New(C.GoString(e))
	}
	return nil
}

// Close() closes the corpus
func (w *Writer) Close() {
	w.mu.Lock()
	if w.opened {
		w.opened = false
		C.alpinocorpus_writer_close(w.c)
	}
	w.mu.Unlock()
}

func (w *Writer) isopen() error {
	if w.opened {
		return nil
	}
	return errors.New("Corpus is closed: " + w.corpusname)
}

// Check whether a particular writer type is available.
// Currently, the only valid types are Compact and Dbxml.
func WriterAvailable(writertype writerType) bool {
	ct := C.CString(wType[writertype])
	defer C.free(unsafe.Pointer(ct))
	if int(C.alpinocorpus_writer_available(ct)) == 0 {
		return false
	}
	return true
}
