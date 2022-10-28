use std::iter::repeat;

pub fn get_diamond(c: char) -> Vec<String> {
    let mut diamond: Vec<String> = vec![];
    let char_position = c as u8 - 'A' as u8; // nth char in alphabet (0 based)
    let width = (2 * char_position + 1) as usize;
    if width == 1 {
        return vec![c.to_string()];
    }
    let middle = char_position as usize;
    let mut current_char = 'A' as u8;
    let (mut left, mut right) = (middle, middle);

    for position in 0..width {
        let mut row = repeat(' ' as u8).take(width).collect::<Vec<u8>>();
        row[left] = current_char;
        row[right] = current_char;
        if position < middle {
            current_char += 1;
            left -= 1;
            right += 1
        } else {
            current_char -= 1;
            left += 1;
            right -= 1
        }
        diamond.push(row.into_iter().map(|c| c as char).collect::<String>())
    }

    diamond
}
