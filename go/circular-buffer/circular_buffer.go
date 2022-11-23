package circular

import "fmt"

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Buffer type here.
type Buffer struct {
	data  []byte
	start int
	end   int
	count int
}

func NewBuffer(size int) *Buffer {
	return &Buffer{
		data: make([]byte, size, size),
	}
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.IsEmpty() {
		return 0, fmt.Errorf("empty buffer")
	}
	datum := b.data[b.start]
	b.start++
	b.start %= len(b.data)
	b.count--
	return datum, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.IsFull() {
		return fmt.Errorf("full buffer")
	}
	b.data[b.end] = c
	b.end++
	b.end %= len(b.data)
	b.count++
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if !b.IsFull() {
		_ = b.WriteByte(c)
		return
	}
	b.data[b.start] = c
	b.start++
	b.start %= len(b.data)
}

func (b *Buffer) Reset() {
	b.start = 0
	b.end = 0
	b.count = 0
}

func (b *Buffer) IsFull() bool {
	return b.count == len(b.data)
}

func (b *Buffer) IsEmpty() bool {
	return b.count == 0
}
