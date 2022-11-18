/// `Palindrome` is a newtype which only exists when the contained value is a palindrome number in base ten.
///
/// A struct with a single field which is used to constrain behavior like this is called a "newtype", and its use is
/// often referred to as the "newtype pattern". This is a fairly common pattern in Rust.
#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub struct Palindrome(u64);

impl Palindrome {
    /// Create a `Palindrome` only if `value` is in fact a palindrome when represented in base ten. Otherwise, `None`.
    pub fn new(value: u64) -> Option<Palindrome> {
        let s = value.to_string();
        if s == s.chars().rev().collect::<String>() {
            return Some(Palindrome(value));
        }
        None
    }

    /// Get the value of this palindrome.
    pub fn into_inner(self) -> u64 {
        self.0
    }
}

pub fn palindrome_products(min: u64, max: u64) -> Option<(Palindrome, Palindrome)> {
    if min > max {
        return None;
    }

    let mut min_palindrome: u64 = u64::MAX;
    let mut max_palindrome: u64 = u64::MIN;

    for i in min..=max {
        for j in min..=max {
            let product = i * j;

            if let Some(palindrome) = Palindrome::new(product) {
                if palindrome.into_inner() < min_palindrome {
                    min_palindrome = palindrome.into_inner();
                }
                if palindrome.into_inner() > max_palindrome {
                    max_palindrome = palindrome.into_inner();
                }
            }
        }
    }

    if min_palindrome == u64::MAX || max_palindrome == u64::MIN {
        return None;
    }

    return Some((
        Palindrome::new(min_palindrome).unwrap(),
        Palindrome::new(max_palindrome).unwrap(),
    ));
}
