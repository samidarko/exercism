use two_bucket::{solve, BucketName, BucketStats};

#[test]
fn test_case_1() {
    assert_eq!(
        solve(3, 5, 1, &BucketName::One),
        Some(BucketStats {
            moves: 4,
            goal_bucket: BucketName::One,
            other_bucket: 5,
        })
    );
}

#[test]
fn test_case_2() {
    assert_eq!(
        solve(3, 5, 1, &BucketName::Two),
        Some(BucketStats {
            moves: 8,
            goal_bucket: BucketName::Two,
            other_bucket: 3,
        })
    );
}

#[test]
fn test_case_3() {
    assert_eq!(
        solve(7, 11, 2, &BucketName::One),
        Some(BucketStats {
            moves: 14,
            goal_bucket: BucketName::One,
            other_bucket: 11,
        })
    );
}

#[test]
fn test_case_4() {
    assert_eq!(
        solve(7, 11, 2, &BucketName::Two),
        Some(BucketStats {
            moves: 18,
            goal_bucket: BucketName::Two,
            other_bucket: 7,
        })
    );
}

#[test]
fn goal_equal_to_start_bucket() {
    assert_eq!(
        solve(1, 3, 3, &BucketName::Two),
        Some(BucketStats {
            moves: 1,
            goal_bucket: BucketName::Two,
            other_bucket: 0,
        })
    );
}

#[test]
fn goal_equal_to_other_bucket() {
    assert_eq!(
        solve(2, 3, 3, &BucketName::One),
        Some(BucketStats {
            moves: 2,
            goal_bucket: BucketName::Two,
            other_bucket: 2,
        })
    );
}

#[test]
fn not_possible_to_reach_the_goal() {
    assert_eq!(solve(6, 15, 5, &BucketName::One), None);
}

#[test]
fn with_same_buckets_but_different_goal_then_it_is_possible() {
    assert_eq!(
        solve(6, 15, 9, &BucketName::One),
        Some(BucketStats {
            moves: 10,
            goal_bucket: BucketName::Two,
            other_bucket: 0,
        })
    );
}
