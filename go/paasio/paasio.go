// Package paasio Please note that DB is the suffix for structs since DB are
// my name and lastname initials. Mutex is mandatory here, even on the get
// count, otherwise tests will fail randomly.
package paasio

import (
	"io"
	"sync"
)

// DBReader is the interface ReadCounter implementation
type DBReader struct {
	amountRead int64
	opsCount   int
	buffer     io.Reader
	mutex      sync.Mutex
}

func (r *DBReader) Read(what []byte) (count int, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	count, err = r.buffer.Read(what)
	if count > 0 {
		r.opsCount++
		r.amountRead += int64(count)
	}
	return
}

// ReadCount returns the amount read
func (r *DBReader) ReadCount() (n int64, nops int) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	n = r.amountRead
	nops = r.opsCount
	return
}

// SetBuffer sets the buffer
func (r *DBReader) SetBuffer(buffer io.Reader) {
	r.buffer = buffer
}

// DBWriter writer implementation
type DBWriter struct {
	amountWrote int64
	opsCount    int
	buffer      io.Writer
	mutex       sync.Mutex
}

// SetBuffer sets the buffer
func (w *DBWriter) SetBuffer(buffer io.Writer) {
	w.buffer = buffer
}

// WriteCount implementation
func (w *DBWriter) WriteCount() (n int64, nops int) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	n = w.amountWrote
	nops = w.opsCount
	return
}

// Write method
func (w *DBWriter) Write(what []byte) (l int, err error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	l, err = w.buffer.Write(what)
	if l > 0 {
		w.amountWrote += int64(l)
		w.opsCount++
	}
	return l, nil
}

// DBReaderWriter implementation
type DBReaderWriter struct {
	reader        DBReader
	writer        DBWriter
	amountRead    int64
	opsCountRead  int
	amountWrote   int64
	opsCountWrite int
	buffer        io.ReadWriter
	mutex         sync.Mutex
}

func (rw *DBReaderWriter) Read(what []byte) (count int, err error) {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	count, err = rw.buffer.Read(what)
	if count > 0 {
		rw.amountRead += int64(count)
		rw.opsCountRead++
	}
	return
}

// ReadCount returns the amount read
func (rw *DBReaderWriter) ReadCount() (n int64, nops int) {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	n = rw.amountRead
	nops = rw.opsCountRead
	return
}

// WriteCount implementation
func (rw *DBReaderWriter) WriteCount() (n int64, nops int) {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	n = rw.amountWrote
	nops = rw.opsCountWrite
	return
}

// Write method
func (rw *DBReaderWriter) Write(what []byte) (l int, err error) {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	l, err = rw.buffer.Write(what)
	if l > 0 {
		rw.amountWrote += int64(l)
		rw.opsCountWrite++
	}
	return l, nil
}

// SetBuffer sets the buffer
func (rw *DBReaderWriter) SetBuffer(buffer io.ReadWriter) {
	rw.buffer = buffer
}

// NewWriteCounter bla
func NewWriteCounter(writer io.Writer) (wc WriteCounter) {
	w := new(DBWriter)
	w.SetBuffer(writer)
	wc = w
	return
}

// NewReadCounter bli
func NewReadCounter(reader io.Reader) (rc ReadCounter) {
	r := new(DBReader)
	r.SetBuffer(reader)
	rc = r
	return
}

// NewReadWriteCounter blu
func NewReadWriteCounter(readWriter io.ReadWriter) (rw ReadWriteCounter) {
	r := new(DBReaderWriter)
	r.SetBuffer(readWriter)
	rw = r
	return
}
