use std::collections::HashMap;

pub fn brackets_are_balanced(string: &str) -> bool {
    let opening_elements: HashMap<char, char> = HashMap::from([(']', '['), ('}', '{'), (')', '(')]);

    let stack: &mut Vec<char> = &mut Vec::new();

    for c in string.chars() {
        match c {
            '[' | '{' | '(' => stack.push(c),
            ']' | '}' | ')' => match stack.pop() {
                Some(last_bracket) => match opening_elements.get(&c) {
                    Some(opening_element) if *opening_element != last_bracket => return false,
                    _ => {}
                },
                None => return false,
            },
            _ => {}
        }
    }

    stack.is_empty()
}
