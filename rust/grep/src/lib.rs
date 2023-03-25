use anyhow::Error;
use std::fs;

/// While using `&[&str]` to handle flags is convenient for exercise purposes,
/// and resembles the output of [`std::env::args`], in real-world projects it is
/// both more convenient and more idiomatic to contain runtime configuration in
/// a dedicated struct. Therefore, we suggest that you do so in this exercise.
///
/// In the real world, it's common to use crates such as [`clap`] or
/// [`structopt`] to handle argument parsing, and of course doing so is
/// permitted in this exercise as well, though it may be somewhat overkill.
///
/// [`clap`]: https://crates.io/crates/clap
/// [`std::env::args`]: https://doc.rust-lang.org/std/env/fn.args.html
/// [`structopt`]: https://crates.io/crates/structopt
#[derive(Debug, Default)]
pub struct Flags {
    n: bool, // Print the line numbers of each matching line.
    l: bool, // Print only the names of files that contain at least one matching line.
    i: bool, // Match line using a case-insensitive comparison.
    v: bool, // Invert the program -- collect all lines that fail to match the pattern.
    x: bool, // Only match entire lines, instead of lines that contain a match.
}

impl Flags {
    pub fn new(flags: &[&str]) -> Self {
        flags
            .iter()
            .fold(Flags::default(), |flags, flag| match *flag {
                "-n" => Self { n: true, ..flags },
                "-l" => Self { l: true, ..flags },
                "-i" => Self { i: true, ..flags },
                "-v" => Self { v: true, ..flags },
                "-x" => Self { x: true, ..flags },
                _ => flags,
            })
    }
}

fn is_match(flags: &Flags, line: &str, pattern: &str) -> bool {
    let (line, pattern): (String, String) = if flags.i {
        (line.to_lowercase(), pattern.to_lowercase())
    } else {
        (line.to_string(), pattern.to_string())
    };

    let matching = if flags.x {
        line == pattern
    } else {
        line.contains(&pattern)
    };

    if flags.v {
        !matching
    } else {
        matching
    }
}

fn format(flags: &Flags, file: &str, line: &str, number: usize, file_prefix: bool) -> String {
    let result = match flags {
        f if f.l => return file.to_string(),
        f if f.n => {
            format!("{}:{line}", number + 1)
        }
        _ => line.to_string(),
    };
    if file_prefix {
        return format!("{file}:{result}");
    }
    result
}

pub fn grep(pattern: &str, flags: &Flags, files: &[&str]) -> Result<Vec<String>, Error> {
    let mut result: Vec<String> = vec![];
    for file in files {
        let content = fs::read_to_string(file)?;
        for (number, line) in content.split('\n').enumerate() {
            if line.is_empty() {
                continue;
            }
            if is_match(flags, line, pattern) {
                let entry = format(flags, file, line, number, files.len() > 1);

                if flags.l
                    && result
                        .last()
                        .map(|last_entry| last_entry == &entry)
                        .or(Some(false))
                        .unwrap()
                {
                    continue;
                } else {
                    result.push(entry);
                }
            }
        }
    }
    Ok(result)
}
