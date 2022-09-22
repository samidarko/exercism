use std::collections::HashMap;

pub fn brackets_are_balanced(string: &str) -> bool {
    let opening_elements: HashMap<char, char> = HashMap::from([(']', '['), ('}', '{'), (')', '(')]);

    let stack: &mut Vec<char> = &mut Vec::new();

    for c in string.chars() {
        match c {
            '[' | '{' | '(' => stack.push(c),
            ']' | '}' | ')' => match (stack.pop(), opening_elements.get(&c)) {
                (Some(last_bracket), Some(opening_element)) if *opening_element != last_bracket => return false,
                (None, _) => return false,
                _ => {},
            },
            _ => {},
        }
    }

    stack.is_empty()
}
