use std::cmp::Ordering;
use num_bigint::{BigInt, ToBigInt};
use num_traits::One;
use num_integer::Integer;
use std::ops::{Add, Mul, Sub};

#[derive(Eq, Ord, Debug)]
pub struct Decimal {
    numerator: BigInt,
    denominator: BigInt,
}

impl PartialEq for Decimal {
    fn eq(&self, other: &Self) -> bool {
        let (_denominator, self_numerator, other_numerator) = self.get_common_denominator(other);
        self_numerator == other_numerator
    }
}

impl PartialOrd for Decimal {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        let (_denominator, self_numerator, other_numerator) = self.get_common_denominator(other);
        self_numerator.partial_cmp(&other_numerator)
    }
}

impl Add for Decimal {
    type Output = Self;

    fn add(self, other: Self) -> Self {
        let (denominator, self_numerator, other_numerator) = self.get_common_denominator(&other);
        Self {
            numerator: self_numerator + other_numerator,
            denominator
        }
    }
}

impl Sub for Decimal {
    type Output = Self;

    fn sub(self, other: Self) -> Self::Output {
        let (denominator, self_numerator, other_numerator) = self.get_common_denominator(&other);
        Self {
            numerator: self_numerator - other_numerator,
            denominator
        }
    }
}

impl Mul for Decimal {
    type Output = Self;

    fn mul(self, other: Self) -> Self::Output {
        Self {
            numerator: self.numerator * other.numerator,
            denominator: self.denominator * other.denominator
        }
    }
}

impl Decimal {
    pub fn try_from(input: &str) -> Option<Decimal> {
        match input.split_once('.') {
            Some((left, right)) if right.trim_end_matches('0').is_empty() => Some(Self {
                numerator: left.parse::<BigInt>().unwrap(),
                denominator: One::one(),
            }),
            Some((_left, right)) => {
                let denominator = 10.to_bigint().unwrap().pow(right.len() as u32);
                let numerator = input.replace(".", "").parse::<BigInt>().unwrap();
                Some(Self {
                    numerator,
                    denominator,
                })
            },
            _ => Self::try_from(format!("{}.0", input).as_str()),
        }
    }

    fn get_common_denominator(&self, other: &Self) -> (BigInt, BigInt, BigInt) {
        let denominator = self.denominator.lcm(&other.denominator);
        let self_numerator = &self.numerator * (&denominator / &self.denominator);
        let other_numerator = &other.numerator * (&denominator / &other.denominator);
        (denominator, self_numerator, other_numerator)
    }
}
