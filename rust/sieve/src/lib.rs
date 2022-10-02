#[derive(Copy, Clone)]
pub struct Number {
    value: u64,
    is_prime: bool,
}

pub fn primes_up_to(upper_bound: u64) -> Vec<u64> {
    let mut numbers: Vec<Number> = vec![];
    for i in 0..upper_bound - 1 {
        numbers.push(Number {
            value: i + 2,
            is_prime: true,
        })
    }

    let mut primes: Vec<u64> = vec![];

    for i in 0..numbers.len() {
        if numbers[i].is_prime {
            primes.push(numbers[i].value);
            let mut n: usize = (numbers[i].value + numbers[i].value) as usize;
            while n <= numbers.len() + 1 {
                numbers[n - 2].is_prime = false;
                n += numbers[i].value as usize;
            }
        }
    }

    primes
}
