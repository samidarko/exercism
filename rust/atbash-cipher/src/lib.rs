use std::collections::HashMap;
/// "Encipher" with the Atbash cipher.
pub fn encode(plain: &str) -> String {
    let mapping = ('a'..='z')
        .zip(('a'..='z').rev())
        .collect::<HashMap<char, char>>();

    plain
        .chars()
        .filter(char::is_ascii_alphanumeric)
        .map(|c| match mapping.get(&c.to_ascii_lowercase()) {
            Some(encoded_char) => *encoded_char,
            None => c,
        })
        .collect::<Vec<char>>()
        .chunks(5)
        .map(String::from_iter)
        .collect::<Vec<String>>()
        .join(" ")
}

/// "Decipher" with the Atbash cipher.
pub fn decode(cipher: &str) -> String {
    let mapping = ('a'..='z')
        .rev()
        .zip('a'..='z')
        .collect::<HashMap<char, char>>();

    cipher
        .chars()
        .filter(char::is_ascii_alphanumeric)
        .map(|c| match mapping.get(&c.to_ascii_lowercase()) {
            Some(decoded_char) => *decoded_char,
            None => c,
        })
        .collect::<String>()
}
