/// Yields each item of a and then each item of b
pub fn append<I, J>(a: I, b: J) -> impl Iterator<Item = I::Item>
where
    I: Iterator,
    J: Iterator<Item = I::Item>,
{
    let mut result = vec![];
    for x in a {
        result.push(x);
    }
    for x in b {
        result.push(x);
    }
    result.into_iter()
}

/// Combines all items in all nested iterators inside into one flattened iterator
pub fn concat<I>(nested_iter: I) -> impl Iterator<Item = <I::Item as Iterator>::Item>
where
    I: Iterator,
    I::Item: Iterator,
{
    let mut result = vec![];
    for x in nested_iter {
        for y in x {
            result.push(y);
        }
    }
    result.into_iter()
}

/// Returns an iterator of all items in iter for which `predicate(item)` is true
pub fn filter<I, F>(iter: I, predicate: F) -> impl Iterator<Item = I::Item>
where
    I: Iterator,
    F: Fn(&I::Item) -> bool,
{
    foldl(iter, vec![], |mut acc, el| {
        if predicate(&el) {
            acc.push(el);
        }
        acc
    })
    .into_iter()
}

pub fn length<I: Iterator>(iter: I) -> usize {
    foldl(iter, 0, |acc, _el| acc + 1)
}

/// Returns an iterator of the results of applying `function(item)` on all iter items
pub fn map<I, F, U>(iter: I, function: F) -> impl Iterator<Item = U>
where
    I: Iterator,
    F: Fn(I::Item) -> U,
{
    foldl(iter, vec![], |mut acc, el| {
        acc.push(function(el));
        acc
    })
    .into_iter()
}

pub fn foldl<I, F, U>(mut iter: I, initial: U, function: F) -> U
where
    I: Iterator,
    F: Fn(U, I::Item) -> U,
{
    match iter.next() {
        Some(first) => foldl(iter, function(initial, first), function),
        None => initial,
    }
}

pub fn foldr<I, F, U>(mut iter: I, initial: U, function: F) -> U
where
    I: DoubleEndedIterator,
    F: Fn(U, I::Item) -> U,
{
    let mut result = initial;
    while let Some(item) = iter.next_back() {
        result = function(result, item);
    }
    result
}

/// Returns an iterator with all the original items, but in reverse order
pub fn reverse<I: DoubleEndedIterator>(iter: I) -> impl Iterator<Item = I::Item> {
    let mut iter = iter.into_iter();
    let mut result = vec![];

    while let Some(item) = iter.next_back() {
        result.push(item);
    }

    result.into_iter()
}
