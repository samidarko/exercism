// This stub file contains items that aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

use std::collections::HashMap;

fn word_count(sentence: &[&str]) -> HashMap<String, i32> {
    let mut words_count: HashMap<String, i32> = HashMap::new();
    for word in sentence.iter() {
        words_count.entry(word.to_string()).and_modify(|count| *count += 1).or_insert(1);
    }
    words_count
}

pub fn can_construct_note(magazine: &[&str], note: &[&str]) -> bool {
    let mut magazine_words: HashMap<String, i32> = word_count(magazine);
    let mut note_words: HashMap<String, i32> = word_count(note);

    for (key, val) in note_words.iter() {
        if magazine_words.get(key).unwrap_or_else(|| &0) < val {
            return false
        }
    }

    true
}
