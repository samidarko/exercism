/// Check a Luhn checksum.
pub fn is_valid(code: &str) -> bool {
    let code = code.replace(" ", "");

    if code.len() < 2 {
        return false;
    }

    let mut digits_sum = 0;
    let mut is_second_digit = code.len() % 2 == 0;

    for c in code.chars() {
        if !c.is_digit(10) {
            return false;
        }
        let mut digit = c.to_digit(10).unwrap();
        if is_second_digit {
            println!("{}", digit);
            digit *= 2;
            if digit > 9 {
                digit -= 9;
            }
        }

        is_second_digit = !is_second_digit;
        digits_sum += digit;
    }

    digits_sum % 10 == 0
}
