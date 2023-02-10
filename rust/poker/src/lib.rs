/// Given a list of poker hands, return a list of those hands which win.
///
/// Note the type signature: this function should return _the same_ reference to
/// the winning hand(s) as were passed in, not reconstructed strings which happen to be equal.
///

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

pub fn winning_hands<'a>(hands: &[&'a str]) -> Vec<&'a str> {
    unimplemented!("Out of {hands:?}, which hand wins?")
}
