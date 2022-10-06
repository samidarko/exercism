use std::collections::HashMap;

pub struct CodonsInfo<'a> {
    codon_name: HashMap<&'a str, &'a str>,
}

impl<'a> CodonsInfo<'a> {
    pub fn name_for(&self, codon: &str) -> Option<&'a str> {
        self.codon_name.get(codon).copied()
    }

    pub fn of_rna(&self, rna: &str) -> Option<Vec<&'a str>> {
        if rna.len() < 9 {
            return None;
        }
        let codons = rna
            .chars()
            .take(9)
            .collect::<Vec<char>>()
            .chunks(3)
            .map(String::from_iter)
            .collect::<Vec<String>>();

        let mut codon_names: Vec<&'a str> = vec![];

        codons.iter().for_each(|codon| match self.name_for(&codon) {
            Some(name) if name != "stop codon" => codon_names.push(name),
            _ => (),
        });

        if codon_names.is_empty() {
            None
        } else {
            Some(codon_names)
        }
    }
}

pub fn parse<'a>(pairs: Vec<(&'a str, &'a str)>) -> CodonsInfo<'a> {
    let codon_name = pairs.iter().copied().collect::<HashMap<&'a str, &'a str>>();
    CodonsInfo { codon_name }
}
