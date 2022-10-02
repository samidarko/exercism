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

// pub fn map<F, T>(input: Vec<T>, function: F) -> Vec<T>
//     where
//         F: Fn(T) -> T,
//         T: Clone + Copy
// {
//     match input.first() {
//         Some(value) => vec![function(*value)]
//             .iter()
//             .copied()
//             .chain(map(input[1..].to_vec(), function))
//             .collect(),
//         None => vec![],
//     }
// }
