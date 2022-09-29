use std::collections::HashMap;

pub fn check(candidate: &str) -> bool {
    let mut chars_map: HashMap<char, bool> = HashMap::new();

    for c in candidate.to_lowercase().chars() {
        if c == '-' || c == ' ' {
            continue;
        }

        if let Some(_) = chars_map.get(&c) {
            return false;
        }

        chars_map.insert(c, true);
    }

    true
}
