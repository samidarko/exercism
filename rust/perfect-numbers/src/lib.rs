use crate::Classification::{Abundant, Deficient, Perfect};

#[derive(Debug, PartialEq, Eq)]
pub enum Classification {
    Abundant,
    Perfect,
    Deficient,
}

pub fn classify(num: u64) -> Option<Classification> {
    if num == 0 {
        return None; // could be in the match but early return avoid computing `get_aliquot_sum(num)`
    }
    match get_aliquot_sum(num) {
        aliquot_sum if aliquot_sum > num => Some(Abundant),
        aliquot_sum if aliquot_sum < num => Some(Deficient),
        _ => Some(Perfect),
    }
}

fn get_aliquot_sum(num: u64) -> u64 {
    (1..=num / 2).filter(|i| num % i == 0).sum()
}
