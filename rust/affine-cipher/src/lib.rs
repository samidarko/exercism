use crate::AffineCipherError::NotCoprime;
use num::integer::gcd;
use std::collections::HashMap;

/// While the problem description indicates a return status of 1 should be returned on errors,
/// it is much more common to return a `Result`, so we provide an error type for the result here.
#[derive(Debug, Eq, PartialEq)]
pub enum AffineCipherError {
    NotCoprime(i32),
}

const ALPHABET: [char; 26] = [
    'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's',
    't', 'u', 'v', 'w', 'x', 'y', 'z',
];

/// Encodes the plaintext using the affine cipher with key (`a`, `b`). Note that, rather than
/// returning a return code, the more common convention in Rust is to return a `Result`.
pub fn encode(plaintext: &str, a: i32, b: i32) -> Result<String, AffineCipherError> {
    let m = ALPHABET.len() as i32;
    if !is_coprime(a, m) {
        return Err(NotCoprime(a));
    }

    let plaintext: String = plaintext
        .to_lowercase()
        .chars()
        .filter(|c| c.is_alphanumeric())
        .collect();

    let letter_index: HashMap<char, i32> = HashMap::from_iter(ALPHABET.into_iter().zip(0..));

    let encoding = plaintext
        .chars()
        .map(|c| match letter_index.get(&c) {
            Some(x) => {
                let index = (a * x + b) % m;
                ALPHABET[index as usize]
            }
            None => c,
        })
        .collect::<Vec<_>>();

    let encoding = encoding
        .chunks(5)
        .map(|s| s.iter().collect())
        .collect::<Vec<String>>();

    Ok(encoding.join(" "))
}

/// Decodes the ciphertext using the affine cipher with key (`a`, `b`). Note that, rather than
/// returning a return code, the more common convention in Rust is to return a `Result`.
pub fn decode(ciphertext: &str, a: i32, b: i32) -> Result<String, AffineCipherError> {
    let mmi = find_mmi(a)?;
    let m = ALPHABET.len() as i32;
    let letter_index: HashMap<char, i32> = HashMap::from_iter(ALPHABET.into_iter().zip(0..));

    let decoding = ciphertext
        .replace(" ", "")
        .chars()
        .map(|c| {
            if c.is_ascii_digit() {
                return c;
            }
            let y = *letter_index.get(&c).unwrap() as i32;

            let index = mmi * (y - b) % m;
            if index >= 0 {
                ALPHABET[index as usize]
            } else {
                ALPHABET[(m + index) as usize]
            }
        })
        .collect::<String>();
    Ok(decoding)
}

fn is_coprime(a: i32, b: i32) -> bool {
    gcd(a, b) == 1 && gcd(b, a) == 1
}

pub fn find_mmi(a: i32) -> Result<i32, AffineCipherError> {
    let m = ALPHABET.len() as i32;
    if !is_coprime(a, m) {
        return Err(NotCoprime(a));
    }

    for n in 1.. {
        if (a * n) % m == 1 {
            return Ok(n);
        }
    }
    Ok(0)
}
