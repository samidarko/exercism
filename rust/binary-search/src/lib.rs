use std::cmp::Ordering;

pub fn find<T>(array: &[T], key: T) -> Option<usize>
where
    T: Ord,
{
    if array.is_empty() {
        return None;
    }
    let (mut left, mut right) = (0usize, array.len() - 1);
    while left <= right {
        let index = (left + right) / 2;
        match array[index].cmp(&key) {
            Ordering::Less => left = index + 1,
            Ordering::Greater => {
                let (value, overflowing) = index.overflowing_sub(1);
                if overflowing {
                    return None;
                }
                right = value;
            }
            Ordering::Equal => return Some(index),
        }
    }
    None
}
