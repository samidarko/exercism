extern crate core;

#[derive(Debug, PartialEq, Eq)]
pub struct Dna(String);

#[derive(Debug, PartialEq, Eq)]
pub struct Rna(String);

const RNA_NUCLEOTIDES: [char; 4] = ['A', 'C', 'G', 'U'];
const DNA_NUCLEOTIDES: [char; 4] = ['A', 'C', 'G', 'T'];

impl Dna {
    pub fn new(dna: &str) -> Result<Dna, usize> {
        for (i, c) in dna.chars().enumerate() {
            if !DNA_NUCLEOTIDES.contains(&c) {
                return Err(i);
            }
        }
        Ok(Dna(dna.to_string()))
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
                _ => panic!("wrong char {}", c),
            })
            .collect::<String>();
        Rna(rna.to_string())
    }
}

impl Rna {
    pub fn new(rna: &str) -> Result<Rna, usize> {
        for (i, c) in rna.chars().enumerate() {
            if !RNA_NUCLEOTIDES.contains(&c) {
                return Err(i);
            }
        }
        Ok(Rna(rna.to_string()))
    }
}
