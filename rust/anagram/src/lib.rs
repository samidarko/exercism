use std::collections::HashSet;

// let anagrams = possible_anagrams.iter().filter(|x| are_anagrams(word, x)).collect::<&[&str]>();
// HashSet::from(anagrams)
pub fn anagrams_for<'a>(word: &str, possible_anagrams: &[&'a str]) -> HashSet<&'a str> {
    let mut anagrams: HashSet<&'a str> = HashSet::new();
    for possible_anagram in possible_anagrams {
        if are_anagrams(word, possible_anagram) {
            anagrams.insert(possible_anagram);
        }
    }
    anagrams
}

fn are_anagrams(left: &str, right: &str) -> bool {
    if left.to_lowercase() == right.to_lowercase() {
        return false
    }
    let mut l: Vec<char> = left.to_lowercase().chars().collect();
    l.sort_unstable();
    let mut r: Vec<char> = right.to_lowercase().chars().collect();
    r.sort_unstable();
    l == r
}
