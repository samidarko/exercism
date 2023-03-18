use std::iter::{Chain, Cycle, Rev};
use std::ops::Range;

const FILL: char = '.';

pub struct RailFence {
    rails: usize,
}

impl RailFence {
    pub fn new(rails: u32) -> RailFence {
        Self {
            rails: rails as usize,
        }
    }

    pub fn encode(&self, text: &str) -> String {
        let grid = self.encrypted_grid(text.replace(' ', "").as_str());

        grid.iter()
            .map(|rail| rail.iter().filter(|c| *c != &FILL).collect::<String>())
            .collect::<String>()
    }

    pub fn decode(&self, cipher: &str) -> String {
        let mut grid = self.encrypted_grid(cipher);
        let cipher: Vec<char> = cipher.chars().collect();

        let mut char_index: usize = 0;
        for rail in grid.iter_mut() {
            for c in rail.iter_mut() {
                if *c != FILL {
                    *c = cipher[char_index];
                    char_index += 1;
                }
            }
        }

        (0..cipher.len())
            .zip(self.get_ys())
            .map(|(x, y)| grid[y][x])
            .collect()
    }

    fn encrypted_grid(&self, text: &str) -> Vec<Vec<char>> {
        let mut grid: Vec<Vec<char>> = vec![vec![FILL; text.len()]; self.rails];
        for ((x, c), y) in (text.chars().enumerate()).zip(self.get_ys()) {
            grid[y][x] = c;
        }
        grid
    }

    fn get_ys(&self) -> Cycle<Chain<Range<usize>, Rev<Range<usize>>>> {
        (0..self.rails).chain((1..self.rails - 1).rev()).cycle()
    }
}
