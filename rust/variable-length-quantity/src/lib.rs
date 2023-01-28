use crate::Error::{IncompleteNumber, Overflow};

#[derive(Debug, PartialEq, Eq)]
pub enum Error {
    IncompleteNumber,
    Overflow,
}

/// Convert a list of numbers to a stream of bytes encoded with variable length encoding.
pub fn to_bytes(values: &[u32]) -> Vec<u8> {
    values.iter().flat_map(|value| encode(*value)).collect()
}

/// Given a stream of bytes, extract all numbers which are encoded in there.
pub fn from_bytes(bytes: &[u8]) -> Result<Vec<u32>, Error> {
    let mut num = 0u32;
    let mut is_last = false;
    let mut result: Vec<u32> = vec![];

    for byte in bytes {
        let mut byte = *byte;

        is_last = byte & 128 == 0;

        if !is_last {
            byte -= 128;
        }

        let value = byte as u32;

        if (32 - format!("{:b}", num).len()) < format!("{:b}", value).len() {
            return Err(Overflow);
        }

        num <<= 7;

        num += value;

        if is_last {
            result.push(num);
            num = 0;
        }
    }

    if !is_last {
        return Err(IncompleteNumber);
    }

    Ok(result)
}

pub fn encode(mut num: u32) -> Vec<u8> {
    let mut is_first = true;
    let mut result: Vec<u8> = vec![];

    loop {
        let mut byte: u8 = (num & 0x7f) as u8;
        num >>= 7;
        if is_first {
            is_first = false;
        } else {
            byte |= 0x80; // set the strongest bit to 1
        }
        result.push(byte);
        if num == 0 {
            break;
        }
    }

    result.reverse();

    result
}
