const SPACE: u8 = b' ';
const ZERO: u8 = b'0';
const MINE: u8 = b'*';

pub fn annotate(minefield: &[&str]) -> Vec<String> {
    let minefield: Vec<&[u8]> = minefield.into_iter().map(|s| s.as_bytes()).collect();

    let row_max = minefield.len();
    if row_max == 0 {
        return vec![];
    }
    let col_max = minefield[0].len();

    let mut result: Vec<String> = vec![];

    for row_index in 0..minefield.len() {
        let mut s = String::new();
        for col_index in 0..minefield[row_index].len() {
            let char = minefield[row_index][col_index];

            if char == MINE {
                s.push(MINE as char);
                continue;
            }

            if char == SPACE {
                let mut mines_count: u8 = ZERO;

                for position in get_surrounding_positions(row_max, col_max, row_index, col_index) {
                    if minefield[position.row][position.col] == MINE {
                        mines_count += 1;
                    }
                }

                if mines_count > ZERO {
                    s.push(mines_count as char);
                } else {
                    s.push(SPACE as char);
                }
            }
        }

        result.push(s);
    }

    result
}

pub struct Position {
    row: usize,
    col: usize,
}

pub fn get_surrounding_positions(
    row_max: usize,
    col_max: usize,
    row_index: usize,
    col_index: usize,
) -> Vec<Position> {
    let mut positions: Vec<Position> = vec![];

    let moves: Vec<(i8, i8)> = vec![
        (-1, -1),
        (-1, 0),
        (-1, 1),
        (0, -1),
        (0, 1),
        (1, -1),
        (1, 0),
        (1, 1),
    ];

    fn calculate_move(index: usize, inc: i8) -> (usize, bool) {
        if inc < 0 {
            return index.overflowing_sub(1);
        }

        (index + inc as usize, false)
    }

    for (row_inc, col_inc) in moves {
        let (row, overflow) = calculate_move(row_index, row_inc);
        if overflow {
            continue;
        }
        let (col, overflow) = calculate_move(col_index, col_inc);
        if overflow {
            continue;
        }

        if row < row_max && col < col_max {
            positions.push(Position { row, col })
        }
    }

    positions
}
