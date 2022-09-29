use std::collections::HashSet;
/// Determine whether a sentence is a pangram.
pub fn is_pangram(sentence: &str) -> bool {
    let mut chars_set: HashSet<char> = HashSet::new();
    sentence
        .to_lowercase()
        .chars()
        .filter(|c| c.is_alphabetic())
        .for_each(|c| {
            chars_set.insert(c);
        });
    chars_set.len() >= 26
}
