extern crate core;

use std::intrinsics::unreachable;

#[derive(Debug, PartialEq, Eq)]
pub struct Dna(String);

#[derive(Debug, PartialEq, Eq)]
pub struct Rna(String);

const RNA_NUCLEOTIDES: [char; 4] = ['A', 'C', 'G', 'U'];
const DNA_NUCLEOTIDES: [char; 4] = ['A', 'C', 'G', 'T'];

impl Dna {
    pub fn new(dna: &str) -> Result<Dna, usize> {
        validate(dna, &DNA_NUCLEOTIDES).map(|s| Dna(s))
    }

    pub fn into_rna(self) -> Rna {
        let rna = self
            .0
            .chars()
            .map(|c| match c {
                'A' => 'U',
                'C' => 'G',
                'G' => 'C',
                'T' => 'A',
                _ => unreachable!(),
            })
            .collect::<String>();
        Rna(rna.to_string())
    }
}

impl Rna {
    pub fn new(rna: &str) -> Result<Rna, usize> {
        validate(rna, &RNA_NUCLEOTIDES).map(|s| Rna(s))
    }
}

fn validate(input: &str, nucleotides: &[char; 4]) -> Result<String, usize> {
    for (i, c) in input.chars().enumerate() {
        if !nucleotides.contains(&c) {
            return Err(i);
        }
    }
    Ok(input.to_string())
}
