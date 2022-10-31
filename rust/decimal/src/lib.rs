use num_bigint::BigUint;

/// Type implementing arbitrary-precision decimal arithmetic
pub struct Decimal {
    // implement your type here
    numerator: BigUint,
    denominator: BigUint,
}

impl Decimal {
    pub fn try_from(input: &str) -> Option<Decimal> {
        // count the number of digits after dot
        // remove dot
        // multiple by 10^(number of digits)

        unimplemented!("Create a new decimal with a value of {}", input)
    }
}
