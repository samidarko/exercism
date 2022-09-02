// This stub file contains items that aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

pub fn production_rate_per_hour(speed: u8) -> f64 {
    let standard_car_production: f64 = 221.0;
    let success_rate: f64 = match speed {
        0 => 0.0,
        1..=4 => 1.0,
        5..=8 => 0.9,
        _ => 0.77,
    };
    standard_car_production * success_rate * speed as f64
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    (production_rate_per_hour(speed) / 60.0) as u32
}
