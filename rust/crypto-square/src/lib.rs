pub fn encrypt(input: &str) -> String {
    let input: Vec<char> = input
        .chars()
        .filter(char::is_ascii_alphanumeric)
        .map(|c| c.to_ascii_lowercase())
        .collect::<Vec<char>>();

    let sqrt = (input.len() as f64).sqrt();
    let rows: usize = sqrt.floor() as usize;
    let cols: usize = sqrt.ceil() as usize;

    let mut output: Vec<char> = vec![];

    for i in 0..cols {
        for j in 0..rows {
            let offset = i + j * cols;
            if offset < input.len() {
                output.push(input[offset]);
            } else {
                output.push(' ');
            }
        }
        if i < cols - 1 {
            output.push(' ');
        }
    }

    output.iter().collect()
}
