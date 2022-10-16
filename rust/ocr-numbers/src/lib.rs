// The code below is a stub. Just enough to satisfy the compiler.
// In order to pass the tests you can add-to or change any of this code.

#[derive(Debug, PartialEq, Eq)]
pub enum Error {
    InvalidRowCount(usize),
    InvalidColumnCount(usize),
}

pub fn convert(input: &str) -> Result<String, Error> {
    let mut output: Vec<String> = vec![];
    let input: Vec<&str> = input
        .split("\n")
        .collect();
    let lines: Vec<&[&str]> = input.chunks(4).collect();

    for line in lines {
        if line.len() != 4 {
            return Err(Error::InvalidRowCount(line.len()));
        }
        let mut digits: Vec<String> = vec![];
        for offset in (0..=line[0].len() - 3).step_by(3) {
            let mut digit: Vec<String> = vec![];
            for i in 0..3 {
                if line[i].len() % 3 != 0 {
                    return Err(Error::InvalidColumnCount(line[i].len()));
                }
                if let Some(value) = line[i].get(offset..offset + 3) {
                    digit.push(value.to_string());
                }
            }
            digits.push(recognize_digit(digit.join("")));
        }
        output.push(digits.join(""));
    }

    Ok(output.join(","))
}

fn recognize_digit(digit: String) -> String {
    let value = match digit.as_str() {
        " _ | ||_|" => "0",
        "     |  |" => "1",
        " _  _||_ " => "2",
        " _  _| _|" => "3",
        "   |_|  |" => "4",
        " _ |_  _|" => "5",
        " _ |_ |_|" => "6",
        " _   |  |" => "7",
        " _ |_||_|" => "8",
        " _ |_| _|" => "9",
        _ => "?",
    };
    value.to_string()
}
