use num_bigint::{BigUint, ToBigUint};
use num_traits::One;

#[derive(Eq, PartialEq)]
pub struct Decimal {
    numerator: BigUint,
    denominator: BigUint,
}

impl Decimal {
    pub fn try_from(input: &str) -> Option<Decimal> {
        match input.split_once('.') {
            Some((left, right)) if right.trim_end_matches('0').is_empty() => Some(Self {
                numerator: left.parse::<BigUint>().unwrap(),
                denominator: One::one(),
            }),
            Some((left, right)) => {
                let denominator = 10.to_biguint().unwrap().pow(right.len() as u32);
                let numerator = left.parse::<BigUint>().unwrap();
                Some(Self {
                    numerator: numerator * &denominator,
                    denominator,
                })
            },
            _ => None,
        }
    }
}
