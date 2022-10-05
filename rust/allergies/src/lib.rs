use self::Allergen::*;

pub struct Allergies {
    score: u32,
}

#[derive(Debug, PartialEq, Eq, Copy, Clone)]
pub enum Allergen {
    Eggs,
    Peanuts,
    Shellfish,
    Strawberries,
    Tomatoes,
    Chocolate,
    Pollen,
    Cats,
}

const ALLERGENS: [Allergen; 8] = [
    Eggs,
    Peanuts,
    Shellfish,
    Strawberries,
    Tomatoes,
    Chocolate,
    Pollen,
    Cats,
];

fn get_allergen_code(allergen: &Allergen) -> u32 {
    use Allergen::*;
    match allergen {
        Eggs => 1,
        Peanuts => 2,
        Shellfish => 4,
        Strawberries => 8,
        Tomatoes => 16,
        Chocolate => 32,
        Pollen => 64,
        Cats => 128,
    }
}

impl Allergies {
    pub fn new(score: u32) -> Self {
        Self { score }
    }

    pub fn is_allergic_to(&self, allergen: &Allergen) -> bool {
        let code = get_allergen_code(allergen);
        self.score & code == code
    }

    pub fn allergies(&self) -> Vec<Allergen> {
        ALLERGENS
            .iter()
            .copied()
            .filter(|a| self.is_allergic_to(a))
            .collect()
    }
}
