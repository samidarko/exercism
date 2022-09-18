#[derive(Debug)]
pub enum CalculatorInput {
    Add,
    Subtract,
    Multiply,
    Divide,
    Value(i32),
}

pub fn evaluate(inputs: &[CalculatorInput]) -> Option<i32> {
    let stack: &mut Vec<i32> = &mut Vec::new();
    fn get_values(s: &mut Vec<i32>) -> Option<(i32, i32)> {
        if s.len() > 1 {
            let b = s.pop().unwrap();
            let a = s.pop().unwrap();
            Some((a, b))
        } else {
            None
        }
    }
    for input in inputs {
        match input {
            CalculatorInput::Value(value) => stack.push(*value),
            CalculatorInput::Add => {
                if let Some((a, b)) = get_values(stack) {
                    stack.push(a + b)
                } else {
                    return None
                }
            },
            CalculatorInput::Subtract => {
                if let Some((a, b)) = get_values(stack) {
                    stack.push(a - b)
                } else {
                    return None
                }
            },
            CalculatorInput::Multiply => {
                if let Some((a, b)) = get_values(stack) {
                    stack.push(a * b)
                } else {
                    return None
                }
            },
            CalculatorInput::Divide => {
                if let Some((a, b)) = get_values(stack) {
                    stack.push(a / b)
                } else {
                    return None
                }
            },
        }
    }

    if stack.len() > 1 {
       None
    } else {
        stack.pop()
    }
}
