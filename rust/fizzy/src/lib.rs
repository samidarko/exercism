use std::fmt::Display;
use std::ops::Rem;

/// A Matcher is a single rule of fizzbuzz: given a function on T, should
pub type MatchFn<T> = Box<dyn Fn(T) -> bool>;

pub struct Matcher<T> {
    matcher: MatchFn<T>,
    subs: String,
}

impl<T> Matcher<T> {
    pub fn new<F, S>(matcher: F, subs: S) -> Matcher<T>
    where
        F: Fn(T) -> bool + 'static,
        S: AsRef<str>,
    {
        Matcher {
            matcher: Box::new(matcher),
            subs: subs.as_ref().to_string(),
        }
    }
}
/// a word be substituted in? If yes, which word?

/// A Fizzy is a set of matchers, which may be applied to an iterator.
///
/// Strictly speaking, it's usually more idiomatic to use `iter.map()` than to
/// consume an iterator with an `apply` method. Given a Fizzy instance, it's
/// pretty straightforward to construct a closure which applies it to all
/// elements of the iterator. However, we're using the `apply` pattern
/// here because it's a simpler interface for students to implement.
///
/// Also, it's a good excuse to try out using impl trait.
pub struct Fizzy<T> {
    matchers: Vec<Matcher<T>>,
}

impl<T: Copy> Fizzy<T>
where
    T: Display,
{
    pub fn new() -> Fizzy<T> {
        Self { matchers: vec![] }
    }

    // feel free to change the signature to `mut self` if you like
    #[must_use]
    pub fn add_matcher(mut self, matcher: Matcher<T>) -> Self {
        self.matchers.push(matcher);
        self
    }

    pub fn apply_to(&self, item: T) -> String {
        let Fizzy { ref matchers } = self;
        let mut out = String::new();
        for matcher in matchers {
            if (matcher.matcher)(item) {
                out += &matcher.subs;
            }
        }
        if out.is_empty() {
            out = item.to_string()
        }
        out
    }

    /// map this fizzy onto every element of an iterator, returning a new iterator
    pub fn apply<I: Iterator<Item = T>>(self, iter: I) -> impl Iterator<Item = String> {
        // unimplemented!() doesn't actually work, here; () is not an Iterator
        // that said, this is probably not the actual implementation you desire
        iter.map(move |item| self.apply_to(item))
    }
}

/// convenience function: return a Fizzy which applies the standard fizz-buzz rules
pub fn fizz_buzz<T>() -> Fizzy<T>
where
    T: Copy + Default + Display + From<u8> + PartialEq + Rem<Output = T> + 'static,
{
    let three: T = 3.into();
    let five: T = 5.into();

    Fizzy::new()
        .add_matcher(Matcher::new(move |n| n % three == T::default(), "fizz"))
        .add_matcher(Matcher::new(move |n| n % five == T::default(), "buzz"))
}
