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

    for input in inputs {
        match input {
            CalculatorInput::Value(value) => stack.push(*value),
            CalculatorInput::Add if stack.len() > 1 => {
                let b = stack.pop().unwrap();
                let a = stack.pop().unwrap();
                stack.push(a + b)
            },
            CalculatorInput::Subtract if stack.len() > 1 => {
                let b = stack.pop().unwrap();
                let a = stack.pop().unwrap();
                stack.push(a - b)
            },
            CalculatorInput::Multiply if stack.len() > 1 => {
                let b = stack.pop().unwrap();
                let a = stack.pop().unwrap();
                stack.push(a * b)
            },
            CalculatorInput::Divide if stack.len() > 1 => {
                let b = stack.pop().unwrap();
                let a = stack.pop().unwrap();
                stack.push(a / b)
            },
            _ => return None
        }
    }

    if stack.len() > 1 {
       return None
    }

    stack.pop()
}
