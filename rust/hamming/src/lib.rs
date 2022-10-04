/// Return the Hamming distance between the strings,
/// or None if the lengths are mismatched.
pub fn hamming_distance(s1: &str, s2: &str) -> Option<usize> {
    if s1.len() != s2.len() {
        return None;
    }
    let errors_count = s1
        .chars()
        .zip(s2.chars())
        .map(|(c1, c2)| if c1 != c2 { 1 } else { 0 })
        .sum::<usize>();
    Some(errors_count)
}
