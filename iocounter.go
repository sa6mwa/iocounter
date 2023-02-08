// Copyright (c) SA6MWA Michel Blomgren
//
// iocounter wraps io.Writer or io.Reader and counts the number of bytes written
// or read to/from them (License == MIT).
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package iocounter

import (
	"fmt"
	"io"
)

type ReadCounter struct {
	reader    io.Reader
	bytesRead int
}

type WriteCounter struct {
	writer       io.Writer
	bytesWritten int
}

func NewReadCounter(r io.Reader) *ReadCounter {
	return &ReadCounter{
		reader:    r,
		bytesRead: 0,
	}
}

func NewWriteCounter(w io.Writer) *WriteCounter {
	return &WriteCounter{
		writer:       w,
		bytesWritten: 0,
	}
}

// Int returns number of bytes read as an int.
func (r *ReadCounter) Int() int {
	return r.bytesRead
}

// String returns number of bytes read as a string.
func (r *ReadCounter) String() string {
	return fmt.Sprintf("%d", r.bytesRead)
}

// Int returns number of bytes written as an int.
func (w *WriteCounter) Int() int {
	return w.bytesWritten
}

// String return number of bytes written as a string.
func (w *WriteCounter) String() string {
	return fmt.Sprintf("%d", w.bytesWritten)
}

// Read implements io.Reader and increments the read counter.
func (r *ReadCounter) Read(b []byte) (n int, err error) {
	n, err = r.reader.Read(b)
	r.bytesRead += n
	return
}

// Write implements io.Writer and increments the write counter.
func (w *WriteCounter) Write(b []byte) (n int, err error) {
	n, err = w.writer.Write(b)
	w.bytesWritten += n
	return
}
