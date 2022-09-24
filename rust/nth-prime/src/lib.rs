use std::ops::Not;

pub fn nth(n: u32) -> u32 {
    if let Some(last_prime) = (2..).filter(|x| is_prime(*x)).nth(n as usize) {
        last_prime
    } else {
        2
    }
}

fn is_prime(n: u32) -> bool {
    let m = n / 2;
    (2..=m).any(|i| n % i == 0).not()
}
