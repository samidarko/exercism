pub fn sum_of_multiples(limit: u32, factors: &[u32]) -> u32 {
    let is_multiple = |n: &u32| -> bool {
        factors.iter().any(|f| *f != 0 && n % *f == 0)
    };
    (1..limit).filter(is_multiple).sum()
}

