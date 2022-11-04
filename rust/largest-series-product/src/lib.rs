use crate::Error::{InvalidDigit, Other, SpanTooLong};

#[derive(Debug, PartialEq, Eq)]
pub enum Error {
    SpanTooLong,
    InvalidDigit(char),
    Other,
}

pub fn lsp(string_digits: &str, span: usize) -> Result<u64, Error> {
    if span == 0 {
        return Ok(1);
    }

    if span > string_digits.len() {
        return Err(SpanTooLong);
    }

    let mut digits: Vec<u32> = vec![];
    for c in string_digits.chars() {
        if let Some(digit) = c.to_digit(10) {
            digits.push(digit);
        } else {
            return Err(InvalidDigit(c));
        }
    }

    let largest_series_product: Vec<u32> = digits
        .windows(span)
        .map(|series| series.iter().product())
        .collect();

    largest_series_product
        .into_iter()
        .max()
        .map(|value| value as u64)
        .ok_or(Other)
}
