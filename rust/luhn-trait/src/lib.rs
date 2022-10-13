use std::ops::ControlFlow;

pub trait Luhn {
    fn valid_luhn(&self) -> bool;
}

/// Here is the example of how to implement custom Luhn trait
/// for the &str type. Naturally, you can implement this trait
/// by hand for the every other type presented in the test suite,
/// but your solution will fail if a new type is presented.
/// Perhaps there exists a better solution for this problem?
impl<T: ToString> Luhn for T {
    fn valid_luhn(&self) -> bool {
        is_valid(self.to_string())
    }
}

pub fn is_valid(code: String) -> bool {
    let code = code.replace(" ", "");

    if code.len() < 2 {
        return false;
    }

    let result = code.chars().try_fold(
        (code.len() % 2 == 0, 0),
        |(is_second_digit, digits_sum), c| {
            if let Some(digit) = c.to_digit(10) {
                match digit * 2 {
                    doubled_digit if is_second_digit && doubled_digit > 9 => {
                        ControlFlow::Continue((!is_second_digit, digits_sum + doubled_digit - 9))
                    }
                    doubled_digit if is_second_digit => {
                        ControlFlow::Continue((!is_second_digit, digits_sum + doubled_digit))
                    }
                    _ => ControlFlow::Continue((!is_second_digit, digits_sum + digit)),
                }
            } else {
                // c is not a digit
                ControlFlow::Break((is_second_digit, digits_sum))
            }
        },
    );

    match result {
        ControlFlow::Continue((_, digits_sum)) => digits_sum % 10 == 0,
        ControlFlow::Break(_) => false,
    }
}
