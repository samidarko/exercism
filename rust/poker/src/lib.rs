use crate::Category::{
    Flush, FourOfKind, FullHouse, HighCard, OnePair, Straight, StraightFlush, ThreeOfKind, TwoPair,
};
use crate::Suit::{Club, Diamond, Heart, Spade};
use std::cmp::Ordering;
use std::collections::HashMap;

/// Given a list of poker hands, return a list of those hands which win.
///
/// Note the type signature: this function should return _the same_ reference to
/// the winning hand(s) as were passed in, not reconstructed strings which happen to be equal.
///

pub fn get_suit(suit: &str) -> Suit {
    match suit {
        "C" => Club,
        "D" => Diamond,
        "H" => Heart,
        "S" => Spade,
        _ => unreachable!(),
    }
}

pub fn get_rank(rank: &str) -> u32 {
    match rank {
        "A" => 14,
        "K" => 13,
        "Q" => 12,
        "J" => 11,
        "10" => 10,
        _ => rank.parse().unwrap(),
    }
}

#[derive(PartialEq, PartialOrd, Copy, Clone)]
pub enum Category {
    HighCard = 0,
    OnePair = 1,
    TwoPair = 2,
    ThreeOfKind = 3,
    Straight = 4,
    Flush = 5,
    FullHouse = 6,
    FourOfKind = 7,
    StraightFlush = 8,
}

#[derive(PartialEq, Copy, Clone)]
pub enum Suit {
    Club,
    Diamond,
    Heart,
    Spade,
}

pub struct Card {
    rank: u32,
    suit: Suit,
}

impl Card {
    pub fn new(input: &str) -> Self {
        let rank = get_rank(&input[0..input.len() - 1]);
        let suit = get_suit(&input[input.len() - 1..input.len()]);
        Self { rank, suit }
    }
}

pub fn is_flush(cards: &[Card]) -> bool {
    let mut suit: Option<Suit> = None;
    for card in cards {
        if suit == None || Some(card.suit) == suit {
            suit = Some(card.suit);
        } else {
            return false;
        }
    }
    return true;
}

pub fn is_straight(cards: &[Card]) -> bool {
    let mut ranks = cards.iter().map(|card| card.rank).collect::<Vec<u32>>();
    ranks.sort_unstable();
    for i in 1..ranks.len() {
        if ranks[i - 1] == 5 && ranks[i] == 14 {
            // aces can start a straight (A 2 3 4 5)
            continue;
        }
        if ranks[i - 1] + 1 != ranks[i] {
            return false;
        }
    }
    true
}

pub fn get_category_and_ranks(cards: &[Card]) -> (Category, Vec<u32>) {
    let mut category = if is_flush(cards) { Flush } else { HighCard };

    let mut ranks = cards.iter().map(|card| card.rank).collect::<Vec<u32>>();

    if is_straight(cards) {
        ranks.sort_unstable_by(|a, b| b.cmp(a));
        if ranks[0] == 14 && ranks[1] < 13 {
            // even though an ace is usually high, a 5-high straight flush is the lowest-scoring straight flush
            ranks.remove(0);
            ranks.push(1);
        }
        if category == Flush {
            return (StraightFlush, ranks);
        }
        return (Straight, ranks);
    }

    let mut rank_count: HashMap<u32, u32> = HashMap::new();
    ranks.into_iter().for_each(|rank| {
        rank_count
            .entry(rank)
            .and_modify(|rank| *rank += 1)
            .or_insert(1);
    });

    let mut rank_groups = rank_count.into_iter().collect::<Vec<_>>();
    rank_groups.sort_unstable_by(|(a_rank, a_count), (b_rank, b_count)| {
        if a_count == b_count {
            return b_rank.cmp(a_rank);
        }
        return b_count.cmp(a_count);
    });

    ranks = vec![];

    for (rank, count) in rank_groups {
        match count {
            4 => category = FourOfKind,
            3 => category = ThreeOfKind,
            2 if category == ThreeOfKind => category = FullHouse,
            2 if category == OnePair => category = TwoPair,
            2 if category == HighCard => category = OnePair,
            _ => {}
        }

        ranks.push(rank);
    }

    (category, ranks)
}

pub struct Hand<'a> {
    input: &'a str,
    category: Category,
    ranks: Vec<u32>,
}

impl<'a> Hand<'a> {
    pub fn new(input: &'a str) -> Self {
        let cards: Vec<Card> = input.split(" ").map(|s| Card::new(s)).collect();
        let (category, ranks) = get_category_and_ranks(&cards);
        Self {
            input,
            category,
            ranks,
        }
    }
}

impl<'a> PartialEq<Self> for Hand<'a> {
    fn eq(&self, other: &Self) -> bool {
        self.category == other.category
            && self
                .ranks
                .iter()
                .zip(other.ranks.iter())
                .all(|(a, b)| a == b)
    }
}

impl<'a> PartialOrd for Hand<'a> {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        if self.category == other.category {
            for i in 0..self.ranks.len() {
                if &self.ranks[i] != &other.ranks[i] {
                    return self.ranks[i].partial_cmp(&other.ranks[i]);
                }
            }
            return Some(Ordering::Equal);
        }

        self.category.partial_cmp(&other.category)
    }
}

pub fn winning_hands<'a>(hands: &[&'a str]) -> Vec<&'a str> {
    let all_hands = hands
        .iter()
        .map(|hand| Hand::new(hand))
        .collect::<Vec<Hand>>();

    let mut best_hand = &all_hands[0];
    let mut result: Vec<&'a str> = vec![best_hand.input];

    for hand in &all_hands[1..] {
        if best_hand == hand {
            result.push(hand.input);
        }

        if best_hand < hand {
            result = vec![hand.input];
            best_hand = &hand;
        }
    }

    result
}
