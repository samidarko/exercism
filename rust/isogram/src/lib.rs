use std::collections::HashSet;

pub fn check(candidate: &str) -> bool {
    let mut chars_set: HashSet<char> = HashSet::new();

    for c in candidate.to_lowercase().chars() {
        if c == '-' || c == ' ' {
            continue;
        }

        if !chars_set.insert(c) {
            return false;
        }
    }

    true
}
