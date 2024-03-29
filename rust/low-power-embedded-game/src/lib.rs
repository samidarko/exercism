// This stub file contains items that aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

pub fn divmod(dividend: i16, divisor: i16) -> (i16, i16) {
    let quotient = dividend / divisor;
    let remainder = dividend - divisor * quotient;
    (quotient, remainder)
}

pub fn is_even(value: &usize) -> bool {
    match divmod(*value as i16, 2) {
        (_, 0) => true,
        _ => false,
    }
}

pub fn evens<T>(iter: impl Iterator<Item=T>) -> impl Iterator<Item=T> {
    iter.enumerate()
        .filter(|(pos, _)| is_even(pos))
        .map(|(_, value)| value)
}

pub struct Position(pub i16, pub i16);

impl Position {
    pub fn manhattan(&self) -> i16 {
        return self.0.abs() + self.1.abs()
    }
}
