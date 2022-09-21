use std::collections::HashSet;

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &[&'a str]) -> HashSet<&'a str> {
    let result: Vec<&str> = possible_anagrams.iter()
        .filter(|a| are_anagrams(word, **a))
        .map(|x| *x)
        .collect();

    HashSet::from_iter(result)
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
