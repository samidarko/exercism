package paasio

import (
	"io"
	"sync"
)

// readCounter type
type readCounter struct {
	reader     io.Reader
	totalOps   int
	totalBytes int64
	sync.RWMutex
}

// writeCounter type
type writeCounter struct {
	writer     io.Writer
	totalOps   int
	totalBytes int64
	sync.RWMutex
}

// writeCounter type
type readWriterCounter struct {
	ReadCounter
	WriteCounter
}

// For the return of the function NewReadWriteCounter, you must also define a type that satisfies the ReadWriteCounter interface.

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer: writer}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader: reader}
}

func NewReadWriteCounter(readWriter io.ReadWriter) ReadWriteCounter {
	return &readWriterCounter{
		NewReadCounter(readWriter),
		NewWriteCounter(readWriter),
	}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	totalBytes, err := rc.reader.Read(p)
	rc.Lock()
	rc.totalOps++
	rc.totalBytes += int64(totalBytes)
	rc.Unlock()
	return totalBytes, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.RLock()
	defer rc.RUnlock()
	return rc.totalBytes, rc.totalOps
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	totalBytes, err := wc.writer.Write(p)
	wc.Lock()
	wc.totalOps++
	wc.totalBytes += int64(totalBytes)
	wc.Unlock()
	return totalBytes, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.RLock()
	defer wc.RUnlock()
	return wc.totalBytes, wc.totalOps
}
