use std::io::{Read, Result, Write};

pub struct ReadStats<R> {
    wrapped: R,
    total_ops: usize,
    total_bytes: usize,
}

impl<R: Read> ReadStats<R> {
    // _wrapped is ignored because R is not bounded on Debug or Display and therefore
    // can't be passed through format!(). For actual implementation you will likely
    // wish to remove the leading underscore so the variable is not ignored.
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
        match self.wrapped.read(buf) {
            Ok(size) => {
                self.total_ops += 1;
                self.total_bytes += size;
                Ok(size)
            }
            err => {
                self.total_ops += 1;
                err
            }
        }
    }
}

pub struct WriteStats<W> {
    wrapped: W,
    total_ops: usize,
    total_bytes: usize,
}

impl<W: Write> WriteStats<W> {
    // _wrapped is ignored because W is not bounded on Debug or Display and therefore
    // can't be passed through format!(). For actual implementation you will likely
    // wish to remove the leading underscore so the variable is not ignored.
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
        match self.wrapped.write(buf) {
            Ok(size) => {
                self.total_ops += 1;
                self.total_bytes += size;
                Ok(size)
            }
            err => {
                self.total_ops += 1;
                err
            }
        }
    }

    fn flush(&mut self) -> Result<()> {
        let result = self.wrapped.flush();
        self.total_ops += 1;
        result
    }
}
