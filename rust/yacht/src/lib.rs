use std::collections::HashMap;

#[derive(Clone, Copy)]
pub enum Category {
    Ones,
    Twos,
    Threes,
    Fours,
    Fives,
    Sixes,
    FullHouse,
    FourOfAKind,
    LittleStraight,
    BigStraight,
    Choice,
    Yacht,
}

type Dice = [u8; 5];
pub fn score(dice: Dice, category: Category) -> u8 {
    use Category::*;
    let mut sum: u8 = 0;
    let mut dice_count: HashMap<u8, u8> = HashMap::new();

    for d in dice {
        *dice_count.entry(d).or_insert(0) += 1;
        sum += d;
    }

    match category {
        Yacht if dice_count.len() == 1 => 50,
        FullHouse if dice_count.len() == 2 => {
            let mut result = 0;
            for (k, v) in dice_count {
                if v > 3 {
                    return 0;
                }
                result += k * v;
            }
            result
        }
        FourOfAKind => {
            for (k, v) in dice_count {
                if v >= 4 {
                    return k * 4;
                }
            }
            0
        }
        LittleStraight if dice_count.len() == 5 => {
            dice_count.get(&6).map(|_| 0).or(Some(30)).unwrap()
        }
        BigStraight if dice_count.len() == 5 => dice_count.get(&1).map(|_| 0).or(Some(30)).unwrap(),
        Choice => sum,
        Ones | Twos | Threes | Fours | Fives | Sixes => dice_count
            .get(&get_category_value(category))
            .map(|value| value * get_category_value(category))
            .or(Some(0))
            .unwrap(),
        _ => 0,
    }
}

fn get_category_value(category: Category) -> u8 {
    use Category::*;
    match category {
        Ones => 1,
        Twos => 2,
        Threes => 3,
        Fours => 4,
        Fives => 5,
        Sixes => 6,
        _ => 0,
    }
}
