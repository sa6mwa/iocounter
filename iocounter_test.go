package iocounter

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestNewReadCounter(t *testing.T) {
	rc := NewReadCounter(os.Stdin)
	if _, ok := rc.reader.(io.Reader); !ok {
		t.Error("expected an io.Reader, got something else")
	}
	if rc.reader != os.Stdin {
		t.Errorf("expected os.Stdin reader, but got something else")
	}
	if rc.bytesRead != 0 {
		t.Errorf("expected 0 bytes, but got %d", rc.bytesRead)
	}
}

func TestNewWriteCounter(t *testing.T) {
	wc := NewWriteCounter(os.Stdout)
	if _, ok := wc.writer.(io.Writer); !ok {
		t.Error("expected an io.Writer, got something else")
	}
	if wc.writer != os.Stdout {
		t.Errorf("expected os.Stdout writer, but got something else")
	}
	if wc.bytesWritten != 0 {
		t.Errorf("expected 0 bytes, but got %d", wc.bytesWritten)
	}
}

func TestInt_read(t *testing.T) {
	const testVal = "hello world"
	outbuf := make([]byte, len(testVal))
	rc := NewReadCounter(bytes.NewBufferString(testVal))
	n, err := rc.Read(outbuf)
	if err != nil {
		t.Fatalf("failed to read: %v", err)
	}
	if n != len(testVal) {
		t.Errorf("expected to read %d bytes, but got %d", len(testVal), n)
	}
	if testVal != string(outbuf) {
		t.Errorf("expected to read %s, but got %s", testVal, string(outbuf))
	}
	if rc.Int() != len(testVal) {
		t.Errorf("expected %d from Int(), but got %d", len(testVal), rc.Int())
	}
}

func TestInt_write(t *testing.T) {
	const testVal = "hello world"
	buf := make([]byte, 0)
	outbuf := bytes.NewBuffer(buf)
	wc := NewWriteCounter(outbuf)
	n, err := wc.Write([]byte(testVal))
	if err != nil {
		t.Fatalf("failed to write: %v", err)
	}
	if n != len(testVal) {
		t.Errorf("expected to write %d bytes, but wrote %d", len(testVal), n)
	}
	if testVal != outbuf.String() {
		t.Errorf("expected to write %s, but got %s", testVal, outbuf.String())
	}
	if wc.Int() != len(testVal) {
		t.Errorf("expected %d from Int(), but got %d", len(testVal), wc.Int())
	}
}
