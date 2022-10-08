use std::collections::HashMap;

/// "Encipher" with the Atbash cipher.
pub fn encode(plain: &str) -> String {
    let range = ('a'..='z').collect();
    map_characters(plain, range)
        .chunks(5)
        .map(String::from_iter)
        .collect::<Vec<String>>()
        .join(" ")
}

/// "Decipher" with the Atbash cipher.
pub fn decode(cipher: &str) -> String {
    let range = ('a'..='z').rev().collect();
    map_characters(cipher, range).iter().collect::<String>()
}

fn map_characters(input: &str, range: Vec<char>) -> Vec<char> {
    let mapping = range
        .iter()
        .copied()
        .zip(range.iter().rev().copied())
        .collect::<HashMap<char, char>>();

    input
        .chars()
        .filter(char::is_ascii_alphanumeric)
        .map(|c| mapping.get(&c.to_ascii_lowercase()).unwrap_or(&c).clone())
        .collect()
}
