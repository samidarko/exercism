use std::collections::{HashMap, HashSet};
// This annotation prevents Clippy from warning us that `School` has a
// `fn new()` with no arguments, but doesn't implement the `Default` trait.
//
// Normally, it's good practice to just do what Clippy tells you, but in this
// case, we want to keep things relatively simple. The `Default` trait is not the point
// of this exercise.
#[allow(clippy::new_without_default)]
pub struct School {
    data: HashMap<u32, HashSet<String>>,
}
// <'a>
impl School {
    pub fn new() -> School {
        Self {
            data: HashMap::new(),
        }
    }

    pub fn add(&mut self, grade: u32, student: &str) {
        match self.data.get_mut(&grade) {
            Some(value) => {
                value.insert(student.to_string());
            }
            None => {
                self.data
                    .insert(grade, HashSet::from([student.to_string()]));
            }
        }
    }

    pub fn grades(&self) -> Vec<u32> {
        let mut g: Vec<u32> = vec![];
        for grade in self.data.keys() {
            g.push(*grade);
        }
        g.sort_unstable();
        g
    }

    // If `grade` returned a reference, `School` would be forced to keep a `Vec<String>`
    // internally to lend out. By returning an owned vector of owned `String`s instead,
    // the internal structure can be completely arbitrary. The tradeoff is that some data
    // must be copied each time `grade` is called.
    pub fn grade(&self, grade: u32) -> Vec<String> {
        match self.data.get(&grade) {
            Some(students) => {
                let students: HashSet<String> = students.clone();
                let mut grades: Vec<String> = students.into_iter().collect();
                grades.sort_unstable();
                grades
            },
            None => vec![],
        }
    }
}
