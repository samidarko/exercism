#[macro_use]
extern crate lazy_static;
use rand::Rng;
use std::collections::HashSet;
use std::sync::Mutex;

lazy_static! {
    static ref USED_NAMES: Mutex<HashSet<String>> = Mutex::new(HashSet::new());
}

const ALPHABET: [char; 26] = [
    'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S',
    'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
];

pub struct Robot {
    name: String,
}

impl Robot {
    pub fn new() -> Self {
        Self {
            name: get_random_name(),
        }
    }

    pub fn name(&self) -> &str {
        return self.name.as_str();
    }

    pub fn reset_name(&mut self) {
        self.name = get_random_name();
    }
}

fn get_random_name() -> String {
    let mut rng = rand::thread_rng();
    let mut used_names = USED_NAMES.lock().unwrap();
    loop {
        let first_char = rng.gen_range(0..26);
        let second_char = rng.gen_range(0..26);
        let number = rng.gen_range(0..1_000);
        let name = format!(
            "{}{}{:03}",
            ALPHABET[first_char], ALPHABET[second_char], number
        );
        if used_names.insert(name.clone()) {
            return name;
        }
    }
}
