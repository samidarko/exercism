use rand::{Rng, SeedableRng};
use std::time::SystemTime;
// use rand_core;

// fn all_names<'a>() -> Vec<String> {
//     let mut names: Vec<String> = vec![];
//     for c in 'A'..='Z' {
//         for d in 'A'..='Z' {
//             for n in 0..1_000 {
//                 names.push(format!("{}{}{:03}", c, d, n));
//             }
//         }
//     }
//     names
// }

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

    // pub fn all_names(self) -> Vec<String> {
    //     let mut names: Vec<String> = vec![];
    //     for c in 'A'..='Z' {
    //         for d in 'A'..='Z' {
    //             for n in 0..1_000 {
    //                 names.push(format!("{}{}{:03}", c, d, n));
    //             }
    //         }
    //     }
    //     names
    // }

    pub fn name(&self) -> &str {
        return self.name.as_str();
    }

    pub fn reset_name(&mut self) {
        self.name = get_random_name();
    }
}

fn get_random_name() -> String {
    let unix_ts: u64 = SystemTime::now()
        .duration_since(SystemTime::UNIX_EPOCH)
        .unwrap()
        .as_nanos() as u64;
    let mut rng = rand_chacha::ChaCha8Rng::seed_from_u64(unix_ts);
    let first_char = rng.gen_range(0..26);
    let second_char = rng.gen_range(0..26);
    let number = rng.gen_range(0..1_000);
    format!(
        "{}{}{:03}",
        ALPHABET[first_char], ALPHABET[second_char], number
    )
}

// fn name_generator() -> impl Fn() -> String {
//
//     let mut names: Vec<String> = vec![];
//     let mut index: usize = 0;
//
//     let function: fn() -> String = || {
//
//         match names.get(index) {
//             Some(name) => name.clone(),
//             None => panic!("no more names"),
//         }
//     };
//     function
// }
