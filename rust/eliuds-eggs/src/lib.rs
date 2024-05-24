pub fn egg_count(display_value: u32) -> usize {
    let mut value = display_value as usize;
    let mut result: usize = 0;

    while value > 0 {
        result += value & 1;
        value >>= 1;
    }
    result
}
