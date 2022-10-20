use std::collections::HashMap;
use std::ops::Not;

const EXCLUDED_CHARS: [char; 7] = ['!', '&', '@', '$', '%', '^', '&'];

/// Count occurrences of words.
pub fn word_count(words: &str) -> HashMap<String, u32> {
    let mut result: HashMap<String, u32> = HashMap::new();
    for elements in words.split_whitespace() {
        elements.split(",").for_each(|word| {
            let word = sanitize(word);
            if word.is_empty().not() {
                result
                    .entry(word)
                    .and_modify(|count| *count += 1)
                    .or_insert(1);
            }
        })
    }
    result
}

fn sanitize(word: &str) -> String {
    let word = word
        .chars()
        .filter(|c| EXCLUDED_CHARS.contains(c).not())
        .collect::<String>();
    return word
        .trim_matches(|c| c == '\'' || c == '.' || c == ':')
        .to_string()
        .to_lowercase();
}
