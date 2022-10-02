use std::ops::Fn;

/// What should the type of _function be?
pub fn map<F>(input: Vec<i32>, function: F) -> Vec<i32>
where
    F: Fn(i32) -> i32,
{
    match input.first() {
        Some(value) => vec![function(*value)]
            .iter()
            .copied()
            .chain(map(input[1..].to_vec(), function))
            .collect(),
        None => vec![],
    }
}
