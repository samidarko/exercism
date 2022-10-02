#[derive(Copy, Clone)]
pub struct Number {
    value: u64,
    is_prime: bool,
}

pub fn primes_up_to(upper_bound: u64) -> Vec<u64> {
    let mut numbers: Vec<Number> = (0..upper_bound - 1)
        .map(|value| Number {
            value: value + 2,
            is_prime: true,
        })
        .collect();

    let mut primes: Vec<u64> = vec![];

    for i in 0..numbers.len() {
        if numbers[i].is_prime {
            primes.push(numbers[i].value);
            let mut n: usize = (numbers[i].value * 2) as usize;
            while n <= numbers.len() + 1 {
                numbers[n - 2].is_prime = false;
                n += numbers[i].value as usize;
            }
        }
    }

    primes
}
