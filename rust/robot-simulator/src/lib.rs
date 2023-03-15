// The code below is a stub. Just enough to satisfy the compiler.
// In order to pass the tests you can add-to or change any of this code.

use crate::Direction::{East, North, South, West};

#[derive(PartialEq, Eq, Debug, Copy, Clone)]
pub enum Direction {
    North,
    East,
    South,
    West,
}

#[derive(Copy, Clone, Debug)]
pub struct Robot {
    x: i32,
    y: i32,
    direction: Direction,
}

impl Robot {
    pub fn new(x: i32, y: i32, direction: Direction) -> Self {
        Self { x, y, direction }
    }

    #[must_use]
    pub fn turn_right(self) -> Self {
        let direction = match self.direction {
            North => East,
            East => South,
            South => West,
            West => North,
        };
        Self::new(self.x, self.y, direction)
    }

    #[must_use]
    pub fn turn_left(self) -> Self {
        let direction = match self.direction {
            North => West,
            West => South,
            South => East,
            East => North,
        };
        Self::new(self.x, self.y, direction)
    }

    #[must_use]
    pub fn advance(self) -> Self {
        let (x, y) = match self.direction {
            North => (self.x, self.y + 1),
            South => (self.x, self.y - 1),
            East => (self.x + 1, self.y),
            West => (self.x - 1, self.y),
        };
        Self::new(x, y, self.direction)
    }

    #[must_use]
    pub fn instructions(self, instructions: &str) -> Self {
        instructions.chars().fold(self, |robot, instruction| {
            return match instruction {
                'A' => robot.advance(),
                'R' => robot.turn_right(),
                'L' => robot.turn_left(),
                _ => unreachable!(),
            };
        })
    }

    pub fn position(&self) -> (i32, i32) {
        (self.x, self.y)
    }

    pub fn direction(&self) -> &Direction {
        &self.direction
    }
}
