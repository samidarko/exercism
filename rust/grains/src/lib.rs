pub fn square(position: u32) -> u64 {
    if position < 1 || position > 64 {
        panic!("Square must be between 1 and 64");
    }
    1 << (position - 1)
}

pub fn total() -> u128 {
    // https://exercism.org/tracks/rust/exercises/grains/solutions/philip98
    // is returning `u64::max_value()` which is equal to `square(64)` and allow to keep u64 as return type
    let grains: u128 = square(64) as u128;
    grains * 2 - 1
}
