pub fn spiral_matrix(size: u32) -> Vec<Vec<u32>> {
    let dimension = size as usize;
    let mut matrix = vec![vec![0; dimension]; dimension];

    let (mut row, mut col): (i32, i32) = (0, 0);
    let (mut row_inc, mut col_inc): (i32, i32) = (0, 1);

    for value in 1u32..=(size * size) {
        matrix[row as usize][col as usize] = value;
        let (next_row, next_col) = (row - row_inc, col + col_inc);
        if next_row >= size as i32
            || next_col >= size as i32
            || next_row.is_negative()
            || next_col.is_negative()
            || matrix[next_row as usize][next_col as usize] > 0
        {
            (row_inc, col_inc) = (-col_inc, row_inc);
        }

        row -= row_inc;
        col += col_inc;
    }

    matrix
}
