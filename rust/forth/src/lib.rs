use crate::Error::{DivisionByZero, InvalidWord, StackUnderflow, UnknownWord};
use std::collections::HashMap;

pub type Value = i32;
pub type Result = std::result::Result<(), Error>;

pub struct Forth {
    stack: Vec<Value>,
    heap: HashMap<String, Vec<String>>,
}

#[derive(Debug, PartialEq, Eq)]
pub enum Error {
    DivisionByZero,
    StackUnderflow,
    UnknownWord,
    InvalidWord,
}

// TODO add this &str
pub fn is_number(s: &str) -> bool {
    s.chars().all(|c| c.is_ascii_digit())
}

impl Forth {
    pub fn new() -> Forth {
        Self {
            stack: vec![],
            heap: HashMap::new(),
        }
    }

    pub fn stack(&self) -> &[Value] {
        &self.stack
    }

    pub fn eval(&mut self, input: &str) -> Result {
        let mut fields = self.parse_input(input);

        while !&fields.is_empty() {
            let token = fields[0].clone();
            fields.remove(0);

            if let Ok(value) = token.parse::<Value>() {
                let _ = &self.stack.push(value);
                continue;
            }

            let stack_size = self.stack.len();
            let fields_size = fields.len();

            match token.as_str() {
                ":" if fields_size >= 3
                    && !is_number(&fields[0])
                    && &fields[fields_size - 1] == ";" =>
                {
                    // TODO see if we can do without cloning
                    let stop = fields.len() - 1;
                    let _ = &self
                        .heap
                        .insert(fields[0].clone(), fields[1..stop].to_vec().clone());
                    fields = vec![]
                }
                ":" => return Err(InvalidWord),
                "drop" if stack_size > 0 => {
                    let _ = &self.stack.pop();
                }
                "dup" if stack_size > 0 => {
                    let _ = &self.stack.push(*self.stack.last().unwrap());
                }
                "over" if stack_size > 1 => {
                    let (b, a) = (&self.stack.pop().unwrap(), &self.stack.pop().unwrap());
                    let _ = &self.stack.push(*a);
                    let _ = &self.stack.push(*b);
                    let _ = &self.stack.push(*a);
                }
                "swap" if stack_size > 1 => {
                    let (b, a) = (&self.stack.pop().unwrap(), &self.stack.pop().unwrap());
                    let _ = &self.stack.push(*b);
                    let _ = &self.stack.push(*a);
                }
                "+" if stack_size > 1 => {
                    let (b, a) = (&self.stack.pop().unwrap(), &self.stack.pop().unwrap());
                    let _ = &self.stack.push(a + b);
                }
                "-" if stack_size > 1 => {
                    let (b, a) = (&self.stack.pop().unwrap(), &self.stack.pop().unwrap());
                    let _ = &self.stack.push(a - b);
                }
                "*" if stack_size > 1 => {
                    let (b, a) = (&self.stack.pop().unwrap(), &self.stack.pop().unwrap());
                    let _ = &self.stack.push(a * b);
                }
                "/" if stack_size > 1 => {
                    let (b, a) = (&self.stack.pop().unwrap(), &self.stack.pop().unwrap());
                    if *b == 0 {
                        return Err(DivisionByZero);
                    }
                    let _ = &self.stack.push(a / b);
                }
                _ if token.len() == 1 => return Err(StackUnderflow),
                "drop" | "dup" | "over" | "swap" => return Err(StackUnderflow),
                _ => return Err(UnknownWord),
            }
        }

        Ok(())
    }

    fn parse_input(&self, input: &str) -> Vec<String> {
        let mut result: Vec<String> = vec![];
        let line = input.to_lowercase();
        let tokens = line.split(" ").collect::<Vec<&str>>();

        for (i, token) in tokens.iter().enumerate() {
            if tokens[0] == ":" && i == 1 {
                result.push(token.to_string());
                continue;
            }

            if let Some(values) = &self.heap.get(&token.to_string()) {
                // TODO use concat
                for value in values.into_iter() {
                    result.push(value.to_string());
                }
            } else {
                result.push(token.to_string());
            }
        }

        result
    }
}
