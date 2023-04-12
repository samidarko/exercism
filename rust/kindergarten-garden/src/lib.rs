pub fn plants(diagram: &str, student: &str) -> Vec<&'static str> {
    let matrix = diagram
        .lines()
        .map(|line| line.chars().collect::<Vec<_>>())
        .collect::<Vec<_>>();
    let initial_position = get_initial_position(student);
    vec![
        get_plant_name(matrix[0][initial_position]),
        get_plant_name(matrix[0][initial_position + 1]),
        get_plant_name(matrix[1][initial_position]),
        get_plant_name(matrix[1][initial_position + 1]),
    ]
}

fn get_plant_name(c: char) -> &'static str {
    match c {
        'C' => "clover",
        'G' => "grass",
        'R' => "radishes",
        'V' => "violets",
        _ => unreachable!(),
    }
}

fn get_initial_position(student: &str) -> usize {
    let offset = 'A' as usize;
    let c = student.chars().next().unwrap();
    ((c as usize) - offset) * 2
}
