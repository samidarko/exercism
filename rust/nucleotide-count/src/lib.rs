use std::collections::HashMap;

pub fn count(nucleotide: char, dna: &str) -> Result<usize, char> {
    match nucleotide_counts(dna) {
        Ok(counts) => match counts.get(&nucleotide) {
            Some(count) => Ok(*count),
            None => Err(nucleotide),
        },
        Err(c) => Err(c),
    }
}

pub fn nucleotide_counts(dna: &str) -> Result<HashMap<char, usize>, char> {
    let mut counts: HashMap<char, usize> = HashMap::from([('A', 0), ('C', 0), ('G', 0), ('T', 0)]);

    for c in dna.chars() {
        match counts.get(&c) {
            Some(count) => {
                counts.insert(c, *count + 1);
            }
            None => return Err(c),
        }
    }

    Ok(counts)
}
