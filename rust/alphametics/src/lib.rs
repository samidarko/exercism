use std::collections::{HashMap, HashSet};

pub fn solve(input: &str) -> Option<HashMap<char, u8>> {
    let parts: Vec<&str> = input.split(" == ").collect();
    if parts.len() != 2 {
        return None;
    }

    let terms: Vec<&str> = parts[0].split(" + ").collect();
    let result = parts[1];

    // Collect unique letters and track first letters
    let mut letters = HashSet::new();
    let mut first_letters = HashSet::new();

    for word in terms.iter().chain(std::iter::once(&result)) {
        if let Some(first_char) = word.chars().next() {
            first_letters.insert(first_char);
        }
        for ch in word.chars() {
            letters.insert(ch);
        }
    }

    let mut letters: Vec<char> = letters.into_iter().collect();
    letters.sort_unstable();

    // Generate permutations and check solutions
    let digits: Vec<u8> = (0..=9).collect();
    let perms = permutations(&digits, letters.len());

    for perm in perms {
        // Check for leading zeros
        if first_letters.iter().any(|&letter| {
            let idx = letters.iter().position(|&c| c == letter).unwrap();
            perm[idx] == 0
        }) {
            continue;
        }

        // Create letter to digit mapping
        let letters_map: HashMap<char, u8> = letters
            .iter()
            .zip(perm.iter().copied())
            .map(|(&letter, digit)| (letter, digit))
            .collect();

        // Check if the equation holds
        let sum: u64 = terms
            .iter()
            .map(|term| word_value(term, &letters_map))
            .sum();

        if sum == word_value(result, &letters_map) {
            return Some(letters_map);
        }
    }

    None
}

fn word_value(word: &str, letters_map: &HashMap<char, u8>) -> u64 {
    word.chars()
        .fold(0, |acc, ch| acc * 10 + letters_map[&ch] as u64)
}

fn permutations(digits: &[u8], n: usize) -> Vec<Vec<u8>> {
    let mut result = Vec::new();
    let mut perm = vec![0; n];
    let mut used = vec![false; digits.len()];

    fn generate(
        pos: usize,
        digits: &[u8],
        perm: &mut [u8],
        used: &mut [bool],
        result: &mut Vec<Vec<u8>>,
    ) {
        if pos == perm.len() {
            result.push(perm.to_vec());
            return;
        }

        for i in 0..digits.len() {
            if !used[i] {
                used[i] = true;
                perm[pos] = digits[i];
                generate(pos + 1, digits, perm, used, result);
                used[i] = false;
            }
        }
    }

    generate(0, digits, &mut perm, &mut used, &mut result);
    result
}
