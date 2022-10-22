use std::cmp::Ordering;

pub fn find<T, U>(array: U, key: T) -> Option<usize>
where
    T: Ord,
    U: AsRef<[T]>,
{
    let array = array.as_ref();
    if array.len() == 0 {
        return None;
    }
    let (mut left, mut right) = (0usize, array.len() - 1);
    while left <= right {
        let index = (left + right) / 2;
        match key.cmp(&array[index]) {
            Ordering::Equal => return Some(index),
            Ordering::Greater => left = index + 1,
            Ordering::Less => {
                let (value, overflowing) = index.overflowing_sub(1);
                if overflowing {
                    return None;
                }
                right = value;
            }
        }
    }
    None
}
