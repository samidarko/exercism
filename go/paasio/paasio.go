package paasio

import (
	"io"
	"sync"
)

var mu sync.RWMutex

// readCounter type
type readCounter struct {
	reader     io.Reader
	totalOps   int
	totalBytes int64
}

// writeCounter type
type writeCounter struct {
	writer     io.Writer
	totalOps   int
	totalBytes int64
}

// writeCounter type
type readWriterCounter struct {
	readCounter
	writeCounter
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
		readCounter{
			readWriter, 0, 0,
		},
		writeCounter{
			readWriter, 0, 0,
		},
	}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	totalBytes, err := rc.reader.Read(p)
	mu.Lock()
	rc.totalOps++
	rc.totalBytes += int64(totalBytes)
	mu.Unlock()
	return totalBytes, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	mu.RLock()
	defer mu.RUnlock()
	return rc.totalBytes, rc.totalOps
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	totalBytes, err := wc.writer.Write(p)
	mu.Lock()
	wc.totalOps++
	wc.totalBytes += int64(totalBytes)
	mu.Unlock()
	return totalBytes, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	mu.RLock()
	defer mu.RUnlock()
	return wc.totalBytes, wc.totalOps
}
