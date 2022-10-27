use num::integer::gcd;

#[derive(PartialEq, Eq, Debug, Hash)]
pub enum BucketName {
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
    pub goal_bucket: BucketName,
    /// How many liters are left in the other bucket?
    pub other_bucket: u8,
}

pub struct Bucket {
    size: u8,
    quantity: u8,
    num_steps: u8,
    name: BucketName,
}

impl Bucket {
    pub fn new(size: u8, name: BucketName) -> Self {
        Self {
            size,
            quantity: 0,
            num_steps: 0,
            name,
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
    pub fn pour(&mut self, into: &mut Bucket) {
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

pub fn next_bucket(bucket_name: &BucketName) -> BucketName {
    match bucket_name {
        BucketName::One => BucketName::Two,
        BucketName::Two => BucketName::One,
    }
}

/// Solve the bucket problem
pub fn solve(
    capacity_1: u8,
    capacity_2: u8,
    goal: u8,
    start_bucket: &BucketName,
) -> Option<BucketStats> {
    if capacity_1 == 0 || capacity_2 == 0 || goal == 0 {
        return None;
    }

    if goal % (gcd(capacity_1, capacity_2)) != 0 {
        return None;
    }

    let (mut current_bucket, mut other_bucket) = match start_bucket {
        BucketName::One => (
            Bucket::new(capacity_1, BucketName::One),
            Bucket::new(capacity_2, BucketName::Two),
        ),
        BucketName::Two => (
            Bucket::new(capacity_2, BucketName::Two),
            Bucket::new(capacity_1, BucketName::One),
        ),
    };

    current_bucket.fill();

    if other_bucket.size == goal {
        other_bucket.fill()
    }

    // while current_bucket.quantity != goal && other_bucket.quantity != goal {
    //     if current_bucket.is_empty() {
    // 		current_bucket.fill();
    //     } else if other_bucket.is_full() {
    // 		other_bucket.empty();
    //     } else {
    // 		current_bucket.pour(other_bucket);
    //     }
    // }

    let moves = current_bucket.num_steps + other_bucket.num_steps;

    if current_bucket.quantity == goal {
        // return current_bucket.name, num_steps, otherBucket.quantity, nil
        return Some(BucketStats {
            moves,
            goal_bucket: current_bucket.name,
            other_bucket: other_bucket.quantity,
        });
    }

    return Some(BucketStats {
        moves,
        goal_bucket: other_bucket.name,
        other_bucket: current_bucket.quantity,
    });
}
