pub fn factors(n: u64) -> Vec<u64> {
    let mut n: u64 = n;
    let mut factor: u64 = 2;
    let mut factors: Vec<u64> = Vec::new();

    while n > 1 {
        if n % factor == 0 {
            factors.push(factor);
            n /= factor;
            continue;
        }
        factor += 1;
    }

    return factors;
}
