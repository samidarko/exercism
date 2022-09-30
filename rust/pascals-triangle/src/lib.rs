pub struct PascalsTriangle(usize);

impl PascalsTriangle {
    pub fn new(row_count: usize) -> Self {
        Self(row_count)
    }

    pub fn rows(&self) -> Vec<Vec<u32>> {
        let mut triangle: Vec<Vec<u32>> = vec![];
        for i in 0..self.0 {
            let mut row = vec![1; i + 1];
            let mut j: usize = 1;
            while j < i {
                row[j] = triangle[i - 1][j - 1] + triangle[i - 1][j];
                j += 1;
            }
            triangle.push(row);
        }
        triangle
    }
}
