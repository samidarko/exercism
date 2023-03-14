use regex::{Captures, Regex};

pub fn answer(command: &str) -> Option<i32> {
    let command = command
        .replace("What is ", "")
        .replace(" by", "")
        .replace(" to the", "");

    let power_re = Regex::new(r"(st|nd|rd|th) power").unwrap();
    let command = power_re.replace_all(&command, "").to_string();

    let re = Regex::new(r"(?:[+-]?\d+|[a-zA-Z]+)").unwrap();
    let tokens: Vec<Captures> = re.captures_iter(&command).collect();

    if tokens.is_empty() {
        return None;
    }

    let tokens = tokens
        .iter()
        .map(|c| c[0].to_string())
        .collect::<Vec<String>>();

    let mut result: i32;
    let mut operation: &str = "";

    match tokens.first().unwrap().parse::<i32>() {
        Ok(n) => result = n,
        _ => return None,
    }

    for (i, s) in tokens[1..].into_iter().enumerate() {
        if i % 2 == 0 {
            match s.as_str() {
                "plus" | "minus" | "multiplied" | "divided" | "raised" => operation = s,
                _ => return None,
            }
        } else {
            let number = match s.parse::<i32>() {
                Ok(n) => n,
                _ => return None,
            };

            match operation {
                "plus" => result += number,
                "minus" => result -= number,
                "multiplied" => result *= number,
                "divided" => result /= number,
                "raised" => result = result.pow(number as u32),
                _ => unreachable!(),
            }

            operation = "";
        }
    }

    if !operation.is_empty() {
        return None;
    }

    Some(result)
}
