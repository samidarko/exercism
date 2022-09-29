use std::collections::HashMap;

pub fn find_saddle_points(input: &[Vec<u64>]) -> Vec<(usize, usize)> {
    let row_min_max: HashMap<usize, (Option<u64>, Option<u64>)> = min_max_map(input);
    let col_min_max: HashMap<usize, (Option<u64>, Option<u64>)> = min_max_map(&transpose(input));

    input
        .iter()
        .enumerate()
        .flat_map(|(row_index, row)| {
            row.iter()
                .enumerate()
                .map(move |(col_index, _)| (row_index, col_index))
        })
        .filter(|(row_index, col_index)| {
            if let (Some((_, Some(row_max))), Some((Some(col_min), _))) =
                (row_min_max.get(row_index), col_min_max.get(col_index))
            {
                let value: u64 = input[*row_index][*col_index];

                if *row_max > value || *col_min < value {
                    return false;
                }

                return true;
            }
            false
        })
        .collect::<Vec<(usize, usize)>>()
}

fn transpose(rows: &[Vec<u64>]) -> Vec<Vec<u64>> {
    let mut cols: Vec<Vec<u64>> = Vec::new();

    let col_size = rows[0].len();

    for col_index in 0..col_size {
        let mut col: Vec<u64> = Vec::new();
        for row_index in 0..rows.len() {
            col.push(rows[row_index][col_index]);
        }
        cols.push(col);
    }

    cols
}

fn min_max_map(input: &[Vec<u64>]) -> HashMap<usize, (Option<u64>, Option<u64>)> {
    input
        .iter()
        .enumerate()
        .map(|(i, row)| (i, (row.iter().min().copied(), row.iter().max().copied())))
        .collect::<HashMap<usize, (Option<u64>, Option<u64>)>>()
}
