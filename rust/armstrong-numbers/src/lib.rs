pub fn is_armstrong_number(num: u32) -> bool {
    let digits: Vec<u32> = num
        .to_string()
        .chars()
        .map(|c| c.to_digit(10).unwrap())
        .collect();

    let power: u32 = digits.len() as u32;

    digits.iter().map(|d| d.pow(power)).sum::<u32>() == num
}
