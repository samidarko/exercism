const ALPHABET: [char; 26] = [
    'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's',
    't', 'u', 'v', 'w', 'x', 'y', 'z',
];

pub fn rotate(input: &str, key: i8) -> String {
    input
        .chars()
        .map(|c| match c.to_ascii_lowercase() {
            _ if c.is_alphabetic() => ALPHABET
                .into_iter()
                .position(|letter| letter == c.to_ascii_lowercase())
                .map(|i| {
                    let position = i as i8 + key;
                    let rotated = match key {
                        k if k >= 0 && position >= 26 => ALPHABET[(position - 26) as usize],
                        _ if position < 0 => ALPHABET[(26 + position) as usize],
                        _ => ALPHABET[position as usize],
                    };
                    if c.is_uppercase() {
                        rotated.to_ascii_uppercase()
                    } else {
                        rotated
                    }
                })
                .unwrap(),
            _ => c,
        })
        .collect::<String>()
}
