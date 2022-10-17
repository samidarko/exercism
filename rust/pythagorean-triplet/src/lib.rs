use std::collections::HashSet;

pub fn find(sum: u32) -> HashSet<[u32; 3]> {
    (1..=sum)
        .flat_map(|a| (a + 1..=sum - a).map(move |b| [a, b, sum - (a + b)]))
        .filter(|[a, b, c]| (a.pow(2) + b.pow(2)) == c.pow(2) && (a + b + c) == sum)
        .collect()
}
