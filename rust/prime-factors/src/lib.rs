pub fn factors(n: u64) -> Vec<u64> {
    let mut input: u64 = n;
    let mut factor: u64 = 2;
    let mut factors: Vec<u64> = Vec::new();

    while input > 1 {
        if input % factor == 0 {
            factors.push(factor);
            input /= factor;
            continue;
        }
        factor += 1;
    }

    return factors;
}
