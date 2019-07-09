package paasio

import (
	"io"
	"sync"
)

type rwCounter struct {
	sync.Mutex
	r    io.Reader
	w    io.Writer
	n    int64
	nops int
}

// ReadCount gives of read bytes and number of operations
func (r *rwCounter) ReadCount() (n int64, nops int) {
	return r.n, r.nops
}

// WriteCount gives of read bytes and number of operations
func (r *rwCounter) WriteCount() (n int64, nops int) {
	return r.n, r.nops
}

// Read reads and records the operation
func (r *rwCounter) Read(p []byte) (int, error) {
	r.Lock()
	defer r.Unlock()
	n, err := r.r.Read(p)
	if err == nil {
		r.n += int64(n)
		r.nops++
	}
	return n, err
}

//Write writes and records the operation
func (r *rwCounter) Write(p []byte) (int, error) {
	r.Lock()
	defer r.Unlock()
	n, err := r.w.Write(p)
	if err == nil {
		r.n += int64(n)
		r.nops++
	}
	return n, err
}

//NewReadCounter returns a new ReadCounter
func NewReadCounter(r io.Reader) ReadCounter {
	return &rwCounter{r: r}
}

//NewWriteCounter returns a new WriteCounter
func NewWriteCounter(w io.Writer) WriteCounter {
	return &rwCounter{w: w}
}

//NewReadWriteCounter returns a new rwcounter
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &rwCounter{r: rw, w: rw}
}
