pub fn square(position: u32) -> u64 {
    if position < 1 || position > 64 {
        panic!("Square must be between 1 and 64");
    }
    1 << (position - 1)
}

pub fn total() -> u128 {
    let grains: u128 = square(64) as u128;
    grains * 2 - 1
}
