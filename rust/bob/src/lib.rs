pub fn reply(message: &str) -> &str {
    match message.trim() {
        m if is_question(&m) && is_valid(m) && m.to_uppercase() == m => {
            "Calm down, I know what I'm doing!"
        }
        m if is_valid(m) && m.to_uppercase() == m => "Whoa, chill out!",
        m if is_question(&m) => "Sure.",
        "" => "Fine. Be that way!",
        _ => "Whatever.",
    }
}

fn is_question(string: &str) -> bool {
    string.chars().last() == Some('?')
}

fn is_valid(string: &str) -> bool {
    string.chars().any(char::is_alphabetic)
}
