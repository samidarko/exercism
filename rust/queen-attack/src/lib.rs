#[derive(Debug)]
pub struct ChessPosition {
    rank: i32,
    file: i32,
}

#[derive(Debug)]
pub struct Queen(ChessPosition);

impl ChessPosition {
    pub fn new(rank: i32, file: i32) -> Option<Self> {
        match (rank, file) {
            (0..=7, 0..=7) => Some(ChessPosition { rank, file }),
            _ => None,
        }
    }
}

impl Queen {
    pub fn new(position: ChessPosition) -> Self {
        Self(position)
    }

    pub fn can_attack(&self, other: &Queen) -> bool {
        let rank_diff = (other.0.rank - self.0.rank).abs();
        let file_diff = (other.0.file - self.0.file).abs();

        if rank_diff == 0 || file_diff == 0 || rank_diff == file_diff {
            return true;
        }

        false
    }
}
