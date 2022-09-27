pub fn series(digits: &str, len: usize) -> Vec<String> {
    let mut substrings: Vec<String> = Vec::new();
    let mut offset = 0;
    let mut limit;

    loop {
        limit = offset + len;
        if limit > digits.len() {
            break;
        }
        substrings.push(digits[offset..limit].to_string());
        offset += 1;
    }

    substrings
}
