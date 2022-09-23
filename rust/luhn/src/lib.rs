/// Check a Luhn checksum.
pub fn is_valid(code: &str) -> bool {
    let code = code.replace(" ", "");

    if code.len() < 2 {
        return false;
    }

    // if a char is not a digit then stop here
    if code.chars().any(|c| !c.is_digit(10)) {
        return false;
    }

    let (_, digits_sum) = code.chars().fold(
        (code.len() % 2 == 0, 0),
        |(is_second_digit, digit_sum), c| {
            let digit = c.to_digit(10).unwrap();
            match digit * 2 {
                doubled_digit if is_second_digit && doubled_digit > 9 => {
                    (!is_second_digit, digit_sum + doubled_digit - 9)
                }
                doubled_digit if is_second_digit => (!is_second_digit, digit_sum + doubled_digit),
                _ => (!is_second_digit, digit_sum + digit),
            }
        },
    );

    digits_sum % 10 == 0
}
