use std::ops::FnMut;

/// What should the type of _function be?
pub fn map<F, T, U>(input: Vec<T>, mut function: F) -> Vec<U>
where
    F: FnMut(T) -> U,
{
    let mut v: Vec<U> = vec![];
    for i in input {
        v.push(function(i));
    }
    v
}
