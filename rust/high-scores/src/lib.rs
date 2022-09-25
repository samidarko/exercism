#[derive(Debug)]
pub struct HighScores<'a> {
    scores: &'a [u32],
}

impl<'a> HighScores<'a> {
    pub fn new(scores: &'a [u32]) -> Self {
        Self { scores }
    }

    pub fn scores(&self) -> &[u32] {
        self.scores
    }

    pub fn latest(&self) -> Option<u32> {
        match self.scores.last() {
            Some(score) => Some(*score),
            None => None,
        }
    }

    pub fn personal_best(&self) -> Option<u32> {
        match self.scores.iter().max() {
            Some(score) => Some(*score),
            None => None,
        }
    }

    pub fn personal_top_three(&self) -> Vec<u32> {
        let mut scores: Vec<u32> = self.scores.to_vec();
        scores.sort_unstable();
        scores.reverse();
        scores.iter().take(3).map(|x| *x).collect::<Vec<u32>>()
    }
}
