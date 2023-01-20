use crate::Direction::*;
use std::ops::Not;

#[derive(Copy, Clone, PartialEq)]
pub struct Position {
    row: isize,
    col: isize,
}

impl Position {
    pub fn new(row: isize, col: isize) -> Self {
        Self { row, col }
    }

    pub fn right(&self) -> Self {
        Self::new(self.row, self.col + 1)
    }

    pub fn left(&self) -> Self {
        Self::new(self.row, self.col - 1)
    }

    pub fn up(&self) -> Self {
        Self::new(self.row - 1, self.col)
    }

    pub fn down(&self) -> Self {
        Self::new(self.row + 1, self.col)
    }
}

#[derive(PartialEq)]
pub enum Direction {
    Up,
    Down,
    Left,
    Right,
    Idle,
}

pub fn count(diagram: &[&str]) -> u32 {
    if diagram.is_empty() || diagram[0].is_empty() {
        return 0;
    }

    let rows_count = diagram.len();
    let cols_count = diagram[0].len();
    let mut count = 0u32;

    for row in 0..rows_count {
        for col in 0..cols_count {
            let c = diagram[row].as_bytes()[col] as char;
            if c == '+' {
                let current_position =
                    Position::new(row.try_into().unwrap(), col.try_into().unwrap());
                let initial_position = Position::new(current_position.row, current_position.col);
                count += explore(diagram, current_position, initial_position, Idle);
            }
        }
    }

    count
}

pub fn explore(
    diagram: &[&str],
    current_position: Position,
    initial_position: Position,
    direction: Direction,
) -> u32 {
    let row_inbounds = 0 <= current_position.row && current_position.row < diagram.len() as isize;
    let col_inbounds =
        0 <= current_position.col && current_position.col < diagram[0].len() as isize;
    if row_inbounds.not() || col_inbounds.not() {
        return 0;
    }

    if direction != Idle && current_position == initial_position {
        return 1;
    }

    let c =
        diagram[current_position.row as usize].as_bytes()[current_position.col as usize] as char;
    match (&c, &direction) {
        ('+', Right) => {
            explore(diagram, current_position.down(), initial_position, Down)
                + explore(diagram, current_position.right(), initial_position, Right)
        }
        ('+', Left) => {
            explore(diagram, current_position.up(), initial_position, Up)
                + explore(diagram, current_position.left(), initial_position, Left)
        }
        ('+', Down) => {
            explore(diagram, current_position.left(), initial_position, Left)
                + explore(diagram, current_position.down(), initial_position, Down)
        }
        ('|' | '+', Up) => explore(diagram, current_position.up(), initial_position, Up),
        ('-', Left) => explore(diagram, current_position.left(), initial_position, Left),
        ('|', Down) => explore(diagram, current_position.down(), initial_position, Down),
        ('-', Right) | ('+', Idle) => {
            explore(diagram, current_position.right(), initial_position, Right)
        }
        _ => 0,
    }
}
