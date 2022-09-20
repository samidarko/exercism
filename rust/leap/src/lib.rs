pub fn is_leap_year(year: u64) -> bool {
    match year {
        _year if year % 4 == 0 && year % 100 == 0 && year % 400 != 0 => false,
        _year if year % 4 == 0  => true,
        _year  => false,
    }
}
