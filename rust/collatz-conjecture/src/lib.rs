pub fn collatz(n: u64) -> Option<u64> {
    if n == 0 {
        return None;
    }

    let mut n: u64 = n;
    let mut steps: u64 = 0;

    while n != 1 {
        if n % 2 == 0 {
            n /= 2
        } else {
            match n.overflowing_mul(3) {
                (result, false) if u64::MAX - result > 1 => n = result + 1,
                _ => return None,
            }
        }

        steps += 1;
    }
    Some(steps)
}
