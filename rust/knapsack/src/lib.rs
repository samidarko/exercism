use std::cmp::max;

pub struct Item {
    pub weight: usize,
    pub value: usize,
}

pub fn maximum_value(max_weight: usize, items: Vec<Item>) -> u32 {
    let mut dp = vec![0usize; max_weight + 1];

    for i in 1..=items.len() {
        for w in (1..=max_weight).rev() {
            if items[i - 1].weight <= w {
                dp[w] = max(dp[w], dp[w - items[i - 1].weight] + items[i - 1].value)
            }
        }
    }

    dp[max_weight] as u32
}
