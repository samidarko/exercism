use num::integer::gcd;

#[derive(PartialEq, Eq, Debug, Hash)]
pub enum Bucket {
    One,
    Two,
}

/// A struct to hold your results in.
#[derive(PartialEq, Eq, Debug)]
pub struct BucketStats {
    /// The total number of "moves" it should take to reach the desired number of liters, including
    /// the first fill.
    pub moves: u8,
    /// Which bucket should end up with the desired number of liters? (Either "one" or "two")
    pub goal_bucket: Bucket,
    /// How many liters are left in the other bucket?
    pub other_bucket: u8,
}

pub struct SmartBucket {
    size: u8,
    quantity: u8,
    num_steps: u8,
    name: Bucket,
}

impl SmartBucket {
    pub fn new(size: u8, name: Bucket) -> Self {
        Self {
            size,
            name,
            quantity: 0,
            num_steps: 0,
        }
    }
    pub fn empty(&mut self) {
        self.quantity = 0;
        self.num_steps += 1;
    }
    pub fn fill(&mut self) {
        self.quantity = self.size;
        self.num_steps += 1;
    }
    pub fn pour(&mut self, into: &mut SmartBucket) {
        while self.quantity > 0 && into.quantity < into.size {
            self.quantity -= 1;
            into.quantity += 1;
        }
        self.num_steps += 1;
    }
    pub fn is_empty(&self) -> bool {
        self.quantity == 0
    }
    pub fn is_full(&self) -> bool {
        self.quantity == self.size
    }
}

pub fn next_bucket(bucket_name: &Bucket) -> Bucket {
    match bucket_name {
        Bucket::One => Bucket::Two,
        Bucket::Two => Bucket::One,
    }
}

/// Solve the bucket problem
pub fn solve(
    capacity_1: u8,
    capacity_2: u8,
    goal: u8,
    start_bucket: &Bucket,
) -> Option<BucketStats> {
    if capacity_1 == 0 || capacity_2 == 0 || goal == 0 {
        return None;
    }

    if goal % (gcd(capacity_1, capacity_2)) != 0 {
        return None;
    }

    let (mut current_bucket, mut other_bucket) = match start_bucket {
        Bucket::One => (
            SmartBucket::new(capacity_1, Bucket::One),
            SmartBucket::new(capacity_2, Bucket::Two),
        ),
        Bucket::Two => (
            SmartBucket::new(capacity_2, Bucket::Two),
            SmartBucket::new(capacity_1, Bucket::One),
        ),
    };

    current_bucket.fill();

    if other_bucket.size == goal {
        other_bucket.fill()
    }

    while current_bucket.quantity != goal && other_bucket.quantity != goal {
        if current_bucket.is_empty() {
            current_bucket.fill();
        } else if other_bucket.is_full() {
            other_bucket.empty();
        } else {
            current_bucket.pour(&mut other_bucket);
        }
    }

    let moves = current_bucket.num_steps + other_bucket.num_steps;

    if current_bucket.quantity == goal {
        return Some(BucketStats {
            moves,
            goal_bucket: current_bucket.name,
            other_bucket: other_bucket.quantity,
        });
    }

    Some(BucketStats {
        moves,
        goal_bucket: other_bucket.name,
        other_bucket: current_bucket.quantity,
    })
}
