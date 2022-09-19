use std::ops::Add;
use std::time::Duration;
use time::PrimitiveDateTime as DateTime;

// Returns a DateTime one billion seconds after start.
pub fn after(start: DateTime) -> DateTime {
    let giga_second = u64::pow(10, 9);
    start.add(Duration::from_secs(giga_second))
}
