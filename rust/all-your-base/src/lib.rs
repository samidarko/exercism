#[derive(Debug, PartialEq, Eq)]
pub enum Error {
    InvalidInputBase,
    InvalidOutputBase,
    InvalidDigit(u32),
}

///
/// Convert a number between two bases.
///
/// A number is any slice of digits.
/// A digit is any unsigned integer (e.g. u8, u16, u32, u64, or usize).
/// Bases are specified as unsigned integers.
///
/// Return an `Err(.)` if the conversion is impossible.
/// The tests do not test for specific values inside the `Err(.)`.
///
///
/// You are allowed to change the function signature as long as all test still pass.
///
///
/// Example:
/// Input
///   number: &[4, 2]
///   from_base: 10
///   to_base: 2
/// Result
///   Ok(vec![1, 0, 1, 0, 1, 0])
///
/// The example corresponds to converting the number 42 from decimal
/// which is equivalent to 101010 in binary.
///
///
/// Notes:
///  * The empty slice ( "[]" ) is equal to the number 0.
///  * Never output leading 0 digits, unless the input number is 0, in which the output must be `[0]`.
///    However, your function must be able to process input with leading 0 digits.
///
pub fn convert(numbers: &[u32], from_base: u32, to_base: u32) -> Result<Vec<u32>, Error> {
    if from_base < 2 {
        return Err(Error::InvalidInputBase);
    }
    if to_base < 2 {
        return Err(Error::InvalidOutputBase);
    }
    let mut sum: u32 = 0;

    for number in numbers {
        sum += number;
        if *number >= from_base {
            return Err(Error::InvalidDigit(*number));
        }
    }

    if sum == 0 {
        return Ok(vec![0]);
    }

    // convert to base 10
    let mut ten_base: u32 = 0;
    let mut power: u32 = numbers.len() as u32;

    for number in numbers {
        power -= 1;
        ten_base += number * from_base.pow(power);
    }

    let mut result: Vec<u32> = vec![];

    while ten_base > 0 {
        let remainder = ten_base % to_base;
        ten_base = ten_base / to_base;
        result.push(remainder);
    }

    result.reverse();
    Ok(result)
}
