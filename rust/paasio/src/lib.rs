use std::io::{Read, Result, Write};

pub struct ReadStats<R> {
    wrapped: R,
    total_ops: usize,
    total_bytes: usize,
}

impl<R: Read> ReadStats<R> {
    pub fn new(wrapped: R) -> ReadStats<R> {
        Self {
            wrapped,
            total_ops: 0,
            total_bytes: 0,
        }
    }

    pub fn get_ref(&self) -> &R {
        &self.wrapped
    }

    pub fn bytes_through(&self) -> usize {
        self.total_bytes
    }

    pub fn reads(&self) -> usize {
        self.total_ops
    }
}

impl<R: Read> Read for ReadStats<R> {
    fn read(&mut self, buf: &mut [u8]) -> Result<usize> {
        self.total_ops += 1;
        self.wrapped.read(buf).map(|bytes_count| {
            self.total_bytes += bytes_count;
            bytes_count
        })
    }
}

pub struct WriteStats<W> {
    wrapped: W,
    total_ops: usize,
    total_bytes: usize,
}

impl<W: Write> WriteStats<W> {
    pub fn new(wrapped: W) -> WriteStats<W> {
        Self {
            wrapped,
            total_ops: 0,
            total_bytes: 0,
        }
    }

    pub fn get_ref(&self) -> &W {
        &self.wrapped
    }

    pub fn bytes_through(&self) -> usize {
        self.total_bytes
    }

    pub fn writes(&self) -> usize {
        self.total_ops
    }
}

impl<W: Write> Write for WriteStats<W> {
    fn write(&mut self, buf: &[u8]) -> Result<usize> {
        self.total_ops += 1;
        self.wrapped.write(buf).map(|bytes_count| {
            self.total_bytes += bytes_count;
            bytes_count
        })
    }

    fn flush(&mut self) -> Result<()> {
        self.wrapped.flush()
    }
}
