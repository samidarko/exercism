/// Determines whether the supplied string is a valid ISBN number
pub fn is_valid_isbn(isbn: &str) -> bool {
    let isbn = isbn.replace("-", "");

    if isbn.len() != 10 {
        return false;
    }

    let mut factor = 10;
    let mut sum = 0;

    for (i, c) in isbn.chars().enumerate() {
        if let Some(n) = c.to_digit(10) {
            sum += n * factor;
            factor -= 1;
        } else if i == 9 && c == 'X' {
            sum += 10;
        } else {
            return false;
        }
    }

    sum % 11 == 0
}
