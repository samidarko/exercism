pub fn encode(mut n: u64) -> String {
    let mut output: Vec<String> = vec![];

    for i in [
        1_000_000_000_000_000_000,
        1_000_000_000_000_000,
        1_000_000_000_000,
        1_000_000_000,
        1_000_000,
        1000,
        100,
    ] {
        let factor = n / i;
        if factor > 0 {
            let encoded = encode(factor);
            output.push(format!("{} {}", encoded, to_string(i)));
            n -= factor * i;
            if n == 0 {
                return output.join(" ");
            }
        }
    }

    if n > 20 {
        let tens = n / 10;
        let (first_digit, second_digit) = (tens * 10, n - tens * 10);
        output.push(format!(
            "{}-{}",
            to_string(first_digit),
            to_string(second_digit)
        ));
        return output.join(" ");
    }

    output.push(to_string(n));

    output.join(" ")
}

fn to_string(n: u64) -> String {
    let s = match n {
        0 => "zero",
        1 => "one",
        2 => "two",
        3 => "three",
        4 => "four",
        5 => "five",
        6 => "six",
        7 => "seven",
        8 => "eight",
        9 => "nine",
        10 => "ten",
        11 => "eleven",
        12 => "twelve",
        13 => "thirteen",
        14 => "fourteen",
        15 => "fifteen",
        16 => "sixteen",
        17 => "seventeen",
        18 => "eighteen",
        19 => "nineteen",
        20 => "twenty",
        30 => "thirty",
        40 => "forty",
        50 => "fifty",
        60 => "sixty",
        70 => "seventy",
        80 => "eighty",
        90 => "ninety",
        100 => "hundred",
        1000 => "thousand",
        1_000_000 => "million",
        1_000_000_000 => "billion",
        1_000_000_000_000 => "trillion",
        1_000_000_000_000_000 => "quadrillion",
        1_000_000_000_000_000_000 => "quintillion",
        _ => panic!("invalid number {}", n),
    };
    s.to_string()
}
