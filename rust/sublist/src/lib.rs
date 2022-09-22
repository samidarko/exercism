#[derive(Debug, PartialEq, Eq)]
pub enum Comparison {
    Equal,
    Sublist,
    Superlist,
    Unequal,
}

pub fn sublist<T: PartialEq>(first_list: &[T], second_list: &[T]) -> Comparison {
    match (first_list, second_list) {
        (f, s) if f.len() == s.len() && f == s => Comparison::Equal,
        (f, s) if f.len() < s.len() && is_sublist(f, s) => Comparison::Sublist,
        (f, s) if f.len() > s.len() && is_sublist(s, f) => Comparison::Superlist,
        _ => Comparison::Unequal,
    }
}

fn is_sublist<T: PartialEq>(first_list: &[T], second_list: &[T]) -> bool {
    first_list.is_empty()
        || second_list
            .windows(first_list.len())
            .any(|s| s == first_list)
}
