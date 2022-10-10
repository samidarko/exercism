pub fn number(user_number: &str) -> Option<String> {
    let mut user_number: String = user_number
        .chars()
        .filter(char::is_ascii_digit)
        .collect::<String>();

    if user_number.starts_with(|c| c == '1') {
        user_number = user_number[1..].to_string();
    }

    if user_number.len() != 10 {
        return None;
    }

    if let (Some(first), Some(second)) = (user_number.chars().nth(0), user_number.chars().nth(3)) {
        if first.to_digit(10).unwrap() < 2 || second.to_digit(10).unwrap() < 2 {
            return None;
        }
    }

    return Some(user_number);
}
