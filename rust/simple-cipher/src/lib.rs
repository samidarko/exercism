use std::iter::repeat;

const A: u8 = 'a' as u8;
const Z: u8 = 'z' as u8;

pub fn encode(key: &str, s: &str) -> Option<String> {
    if key.is_empty() || key.chars().any(|c|c.is_uppercase() || c.is_ascii_digit()) {
        return None
    }
    let encoding = s.chars().zip(repeat(key.chars()).flat_map(|c| c)).map(|(c, k)| {
        let order = k as u8 - A;
        wrap(c as u8 + order)
    }).collect::<String>();
    Some(encoding)
}

pub fn decode(key: &str, s: &str) -> Option<String> {
    if key.is_empty() || key.chars().any(|c|c.is_uppercase() || c.is_ascii_digit()) {
        return None
    }
    let decoding = s.chars().zip(repeat(key.chars()).flat_map(|c| c)).map(|(c, k)| {
        let order = k as u8 - A;
        wrap(c as u8 - order)
    }).collect::<String>();
    Some(decoding)
}

pub fn encode_random(s: &str) -> (String, String) {
    unimplemented!(
        "Generate random key with only a-z chars and encode {}. Return tuple (key, encoded s)",
        s
    )
}

pub fn wrap(c: u8) -> char {
    let c = match c {
        c if c > Z =>  A + (c - (Z + 1)),
        c if c < A =>  Z - ((A - 1) - c),
        c => c
    };
    c as char
}

// func wrap(r rune) rune {
// 	if r > 'z' {
// 		// if passed 'z' returns to 'a'
// 		r = 'a' + (r - ('z' + 1))
// 	}
//
// 	if r < 'a' {
// 		// if passed 'a' returns to 'z'
// 		r = 'z' - (('a' - 1) - r)
// 	}
//
// 	return r
// }