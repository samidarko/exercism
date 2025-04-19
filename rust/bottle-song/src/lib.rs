pub fn recite(start_bottles: u32, take_down: u32) -> String {
    let mut result = String::new();
    for i in 0..take_down {
        let bottle_count = start_bottles - i;
        let first_sentence = format!(
            "{} green bottle{} hanging on the wall,\n",
            digit_to_string(bottle_count, true),
            plural(bottle_count)
        );
        result.push_str(&first_sentence.clone());
        result.push_str(&first_sentence);
        result.push_str("And if one green bottle should accidentally fall,\n");
        result.push_str(&format!(
            "There'll be {} green bottle{} hanging on the wall.\n\n",
            digit_to_string(bottle_count - 1, false),
            plural(bottle_count - 1)
        ))
    }
    result
}

fn plural(i: u32) -> String {
    let result = if i == 1 { "" } else { "s" };
    result.to_string()
}

fn digit_to_string(i: u32, capitalize: bool) -> String {
    let result = match i {
        10 => "Ten".to_string(),
        9 => "Nine".to_string(),
        8 => "Eight".to_string(),
        7 => "Seven".to_string(),
        6 => "Six".to_string(),
        5 => "Five".to_string(),
        4 => "Four".to_string(),
        3 => "Three".to_string(),
        2 => "Two".to_string(),
        1 => "One".to_string(),
        0 => "No".to_string(),
        _ => unreachable!(),
    };
    if capitalize {
        return result
    }
    result.to_lowercase()
}
