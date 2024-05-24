pub fn egg_count(display_value: u32) -> usize {
    let mut value = display_value;
    let mut result: usize = 0;

    while value > 0 {
        if value & 1 == 1 {
            result += 1;
        }
        value = value >> 1;
    }
    result
}
