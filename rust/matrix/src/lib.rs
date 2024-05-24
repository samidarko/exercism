pub struct Matrix {
    n: usize,
    m: usize,
    data: Vec<Vec<u32>>,
}

impl Matrix {
    pub fn new(input: &str) -> Self {
        let input: Vec<&str> = input.split("\n").collect();
        let m = input.len();
        let data: Vec<Vec<u32>> = input
            .iter()
            .map(|s| {
                s.split(" ")
                    .map(|value| value.parse::<u32>().unwrap())
                    .collect()
            })
            .collect();

        let n = data[0].len();
        Self { data, m, n }
    }

    pub fn row(&self, row_no: usize) -> Option<Vec<u32>> {
        let row_no = row_no - 1;
        if row_no < self.m {
            return Some(self.data[row_no].clone());
        }
        None
    }

    pub fn column(&self, col_no: usize) -> Option<Vec<u32>> {
        let col_no = col_no - 1;
        if col_no < self.n {
            let col: Vec<u32> = (0..self.m).map(|i| self.data[i][col_no]).collect();
            return Some(col);
        }
        None
    }
}
