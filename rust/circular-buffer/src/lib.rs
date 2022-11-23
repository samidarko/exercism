use crate::Error::{EmptyBuffer, FullBuffer};

pub struct CircularBuffer<T> {
    data: Vec<T>,
    start: usize,
    end: usize,
    count: usize,
}

#[derive(Debug, PartialEq, Eq)]
pub enum Error {
    EmptyBuffer,
    FullBuffer,
}

impl<T: Default + Clone> CircularBuffer<T> {
    pub fn new(capacity: usize) -> Self {
        Self {
            data: vec![T::default(); capacity],
            start: 0,
            end: 0,
            count: 0,
        }
    }

    pub fn write(&mut self, element: T) -> Result<(), Error> {
        if self.is_full() {
            return Err(FullBuffer);
        }

        self.data[self.end] = element;
        self.end += 1;
        self.end %= self.data.len();
        self.count += 1;

        Ok(())
    }

    pub fn read(&mut self) -> Result<T, Error> {
        if self.is_empty() {
            return Err(EmptyBuffer);
        }

        let datum = self.data[self.start].clone();
        self.start += 1;
        self.start %= self.data.len();
        self.count -= 1;

        Ok(datum)
    }

    pub fn clear(&mut self) {
        self.data = vec![T::default(); self.data.len()];
        self.start = 0;
        self.end = 0;
        self.count = 0;
    }

    pub fn overwrite(&mut self, element: T) {
        if !self.is_full() {
            let _ = self.write(element);
            return;
        }

        self.data[self.start] = element;
        self.start += 1;
        self.start %= self.data.len();
    }

    pub fn is_empty(&self) -> bool {
        self.count == 0
    }

    pub fn is_full(&self) -> bool {
        self.count == self.data.len()
    }
}
