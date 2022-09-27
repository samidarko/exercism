use regex::{Captures, Regex};

// TODO even if this solution is working, it's too complex. Let's think about it again someday.

pub fn abbreviate(phrase: &str) -> String {
    phrase
        .replace("_", "")
        .replace("-", " ")
        .split_whitespace()
        .flat_map(|word| un_camel_case(word))
        .map(|word| word.chars().nth(0).unwrap().to_uppercase().to_string())
        .collect::<String>()
}

fn un_camel_case(string: &str) -> Vec<String> {
    let re = Regex::new(r"(?:[A-Z][a-z]+)").unwrap();
    let caps: Vec<Captures> = re.captures_iter(string).collect();

    if caps.len() > 1 {
        return caps
            .iter()
            .map(|c| c[0].to_string())
            .collect::<Vec<String>>();
    }

    string
        .split_whitespace()
        .map(&str::to_string)
        .collect::<Vec<String>>()
}
