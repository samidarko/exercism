use num::Zero;
use std::cmp::PartialOrd;
use std::ops::Add;

pub struct Triangle<T> {
    a: T,
    b: T,
    c: T,
}

impl<T> Triangle<T>
where
    T: Add<Output = T> + PartialOrd + Copy + Zero,
{
    pub fn build(sides: [T; 3]) -> Option<Triangle<T>> {
        let [a, b, c] = sides;
        let triangle = Self { a, b, c };

        match triangle.is_a_triangle() {
            true => Some(triangle),
            false => None,
        }
    }

    pub fn is_equilateral(&self) -> bool {
        self.a == self.b && self.b == self.c && self.c == self.a
    }

    pub fn is_scalene(&self) -> bool {
        self.a != self.b && self.b != self.c && self.c != self.a
    }

    pub fn is_isosceles(&self) -> bool {
        self.a == self.b || self.b == self.c || self.c == self.a
    }

    pub fn is_a_triangle(&self) -> bool {
        (self.a > T::zero() && self.b > T::zero() && self.c > T::zero())
            && (self.a + self.b >= self.c)
            && (self.b + self.c >= self.a)
            && (self.c + self.a >= self.b)
    }
}
