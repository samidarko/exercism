use std::cmp::Eq;
use std::collections::HashMap;
use std::fmt::Debug;
use std::hash::Hash;

#[derive(Debug, PartialEq, Eq)]
pub struct CustomSet<T>
where
    T: Eq + Hash,
{
    elements: HashMap<T, bool>,
}

impl<T> CustomSet<T>
where
    T: Eq + Hash + Copy,
{
    pub fn new(input: &[T]) -> Self {
        let mut set: CustomSet<T> = Self {
            elements: HashMap::new(),
        };
        input.iter().for_each(|element| set.add(*element));
        set
    }

    pub fn contains(&self, element: &T) -> bool {
        self.elements.contains_key(&element)
    }

    pub fn add(&mut self, element: T) {
        self.elements.insert(element, true);
    }

    pub fn get_elements(&self) -> Vec<T> {
        self.elements.keys().copied().collect::<Vec<_>>()
    }

    pub fn is_subset(&self, other: &Self) -> bool {
        self.get_elements()
            .iter()
            .map(|element| other.contains(element))
            .all(|v| v)
    }

    pub fn len(&self) -> usize {
        self.elements.len()
    }

    pub fn is_empty(&self) -> bool {
        self.len() == 0
    }

    pub fn is_disjoint(&self, other: &Self) -> bool {
        self.intersection(other).is_empty()
    }

    #[must_use]
    pub fn intersection(&self, other: &Self) -> Self {
        let mut set: CustomSet<T> = CustomSet::new(&[]);
        let mut s1 = self;
        let mut s2 = other;
        if s2.len() < s1.len() {
            (s1, s2) = (s2, s1);
        }

        s1.get_elements().iter().for_each(|element| {
            if s2.contains(element) {
                set.add(*element);
            }
        });
        set
    }

    #[must_use]
    pub fn difference(&self, other: &Self) -> Self {
        let mut set: CustomSet<T> = CustomSet::new(&[]);
        self.get_elements().iter().for_each(|element| {
            if !other.contains(element) {
                set.add(*element);
            }
        });
        set
    }

    #[must_use]
    pub fn union(&self, other: &Self) -> Self {
        let mut set: CustomSet<T> = CustomSet::new(&[]);
        self.get_elements()
            .iter()
            .for_each(|element| set.add(*element));
        other
            .get_elements()
            .iter()
            .for_each(|element| set.add(*element));
        set
    }
}
