use std::fmt::{Display, Formatter, Result};
use std::iter::repeat;

pub struct Roman(u32);

impl Display for Roman {
    fn fmt(&self, f: &mut Formatter<'_>) -> Result {
        write!(f, "{}", process(self.0))
    }
}

impl From<u32> for Roman {
    fn from(num: u32) -> Self {
        Self(num)
    }
}

fn get_roman(n: u32) -> &'static str {
    match n {
        1 => "I",
        2 => "II",
        3 => "III",
        4 => "IV",
        5 => "V",
        6 => "VI",
        7 => "VII",
        8 => "VIII",
        9 => "IX",
        10 => "X",
        40 => "XL",
        50 => "L",
        90 => "XC",
        100 => "C",
        400 => "CD",
        500 => "D",
        900 => "CM",
        1000 => "M",
        _ => unreachable!(),
    }
}

fn process(value: u32) -> String {
    let mut value = value;
    let mut output = String::new();
    match value {
        0 => return String::new(),
        _ if value >= 1000 => {
            let n = value / 1000;
            output.push_str(repeat(get_roman(1000)).take(n as usize).collect::<String>().as_str());
            value -= 1000 * n;
        }
        _ if value >= 900 => {
            output.push_str(get_roman(900));
            value -= 900;
        }
        _ if value >= 500 => {
            output.push_str(get_roman(500));
            value -= 500;
        }
        _ if value >= 400 => {
            output.push_str(get_roman(400));
            value -= 400;
        }
        _ if value / 100 > 0 => {
            let n = value / 100;
            output.push_str(repeat(get_roman(100)).take(n as usize).collect::<String>().as_str());
            value -= 100 * n;
        }
        _ if value >= 90 => {
            output.push_str(get_roman(90));
            value -= 90;
        }
        _ if value >= 50 => {
            output.push_str(get_roman(50));
            value -= 50;
        }
        _ if value >= 40 => {
            output.push_str(get_roman(40));
            value -= 40;
        }
        _ if value / 10 > 0 => {
            let n = value / 10;
            output.push_str(repeat(get_roman(10)).take(n as usize).collect::<String>().as_str());
            value -= 10 * n;
        }
        _ => {
            output.push_str(get_roman(value));
            return output;
        }
    }

    output.push_str(process(value).as_str());
    output
}
