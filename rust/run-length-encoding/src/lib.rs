use core::iter::repeat;
use std::ops::Not;

pub fn encode(source: &str) -> String {
    let mut encoding: Vec<String> = vec![];
    let mut repetition: usize = 0;
    let mut character: Option<char> = None;

    for c in source.chars() {
        if Some(c) != character {
            if let Some(repeated) = character {
                encoding.push(encoder(repeated, repetition));
            }
            repetition = 1;
            character = Some(c);
        } else {
            repetition += 1;
        }
    }

    if let Some(repeated) = character {
        encoding.push(encoder(repeated, repetition));
    }
    encoding.join("")
}

fn encoder(c: char, repetition: usize) -> String {
    match repetition {
        0 => "".to_string(),
        1 => c.to_string(),
        n => format!("{}{}", n, c),
    }
}

pub fn decode(mut source: &str) -> String {
    let mut decoding = String::new();
    while source.is_empty().not() {
        let repetition = source
            .chars()
            .take_while(|c| c.is_ascii_digit())
            .collect::<String>();
        let repetition_length = repetition.len();
        let c: char = source.chars().nth(repetition_length).unwrap();
        decoding.push_str(decoder(c, repetition).as_str());
        source = &source[repetition_length + 1..];
    }
    decoding
}

fn decoder(c: char, repetition: String) -> String {
    if repetition.is_empty() {
        return c.to_string();
    }
    let repetition: usize = repetition.parse::<usize>().unwrap();
    repeat(c).take(repetition).collect::<String>()
}
