use crate::Error::*;
use crate::Tabulation::*;

const LAST_GAME: usize = 9;

#[derive(Debug, PartialEq, Eq)]
pub enum Error {
    NotEnoughPinsLeft,
    GameComplete,
}

#[derive(Copy, Clone, PartialEq)]
enum Tabulation {
    OpenFrame,
    Spare,
    Strike,
}

#[derive(Clone)]
pub struct Frame {
    throws: Vec<u16>,
    tabulation: Tabulation,
}

impl Frame {
    pub fn add_pins(&mut self, pins: u16) {
        self.throws.push(pins);
        match (self.total_pins_count(), self.throws_count()) {
            (10, 1) => {
                self.tabulation = Strike;
            }
            (10, 2) => {
                self.tabulation = Spare;
            }
            _ => (),
        }
    }

    pub fn total_pins_count(&self) -> u16 {
        self.throws.iter().sum()
    }

    pub fn throws_count(&self) -> usize {
        self.throws.len()
    }
}

pub struct BowlingGame {
    frames: Vec<Frame>,
    current: usize,
}

impl BowlingGame {
    pub fn new() -> Self {
        let frames = (0..=LAST_GAME)
            .map(|_| Frame {
                throws: vec![],
                tabulation: OpenFrame,
            })
            .collect();
        Self { frames, current: 0 }
    }

    pub fn current_frame(&mut self) -> Option<&mut Frame> {
        self.frames.get_mut(self.current)
    }

    pub fn roll(&mut self, pins: u16) -> Result<(), Error> {
        if pins > 10 {
            return Err(NotEnoughPinsLeft);
        }
        if self.current > LAST_GAME {
            return Err(GameComplete);
        }

        self.current_frame().map(|frame| frame.add_pins(pins));

        let current = self.current;

        if let Some(current_frame) = self.frames.get(self.current) {
            if current < LAST_GAME && current_frame.total_pins_count() > 10 {
                return Err(NotEnoughPinsLeft);
            }

            if current > 8
                && current_frame.tabulation == Strike
                && current_frame.throws_count() == 3
            {
                if let Some(second_throw) = current_frame.throws.get(1) {
                    if *second_throw < 10 && current_frame.total_pins_count() > 20 {
                        return Err(NotEnoughPinsLeft);
                    }
                }
            }
        }

        if let Some(frame) = self.current_frame() {
            if (frame.tabulation == OpenFrame && frame.throws_count() == 2)
                || ((frame.tabulation == Strike || frame.tabulation == Spare)
                    && current < LAST_GAME)
                || frame.throws_count() == 3
            {
                self.current += 1;
            }
        }

        Ok(())
    }

    pub fn frame_score(&self, i: usize) -> u16 {
        if let Some(frame) = self.frames.get(i) {
            let mut score = frame.total_pins_count();
            if i == LAST_GAME {
                return score;
            }
            return match frame.tabulation {
                Strike => {
                    self.frames.get(i + 1).map(|next_frame| {
                        if next_frame.throws.len() > 1 {
                            score += next_frame.throws.iter().take(2).sum::<u16>();
                        } else {
                            self.frames.get(i + 2).map(|next_next_frame| {
                                score += next_frame.throws.first().unwrap_or(&0)
                                    + next_next_frame.throws.first().unwrap_or(&0)
                            });
                        }
                    });
                    score
                }
                Spare => {
                    self.frames.get(i + 1).map(|next_frame| {
                        score += next_frame.throws.first().unwrap_or(&0);
                    });
                    score
                }
                OpenFrame => score,
            };
        }

        0
    }

    pub fn score(&self) -> Option<u16> {
        if self.current < LAST_GAME {
            return None;
        }
        if let Some(last_frame) = self.frames.get(LAST_GAME) {
            match last_frame.tabulation {
                Strike | Spare if last_frame.throws.len() < 3 => return None,
                OpenFrame if last_frame.throws.len() < 2 => return None,
                _ => (),
            }
        }
        Some((0..self.frames.len()).map(|i| self.frame_score(i)).sum())
    }
}
