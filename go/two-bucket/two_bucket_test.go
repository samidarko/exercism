package twobucket

import "testing"

func runTestCase(t *testing.T, tc bucketTestCase) {
	g, m, other, e := Solve(tc.bucketOne, tc.bucketTwo, tc.goal, tc.startBucket)
	var _ error = e
	if e == nil {
		if tc.errorExpected {
			t.Fatalf("FAIL: %s\nSolve(%d, %d, %d, %q)\nExpected error\nActual: %q, %d, %d",
				tc.description,
				tc.bucketOne, tc.bucketTwo, tc.goal, tc.startBucket,
				g, m, other)
		}
		if g != tc.goalBucket || m != tc.moves || other != tc.otherBucket {
			t.Fatalf("FAIL: %s\nSolve(%d, %d, %d, %q)\nExpected: %q, %d, %d\nActual: %q, %d, %d",
				tc.description,
				tc.bucketOne, tc.bucketTwo, tc.goal, tc.startBucket,
				tc.goalBucket, tc.moves, tc.otherBucket,
				g, m, other)
		}
	} else if !tc.errorExpected {
		t.Fatalf("FAIL: %s\nSolve(%d, %d, %d, %q)\nExpected: %q, %d, %d\nGot Error %q",
			tc.description,
			tc.bucketOne, tc.bucketTwo, tc.goal, tc.startBucket,
			tc.goalBucket, tc.moves, tc.otherBucket,
			e)
	}
	t.Logf("PASS: %s", tc.description)
}

func TestSolve(t *testing.T) {
	for _, tc := range append(testCases, errorTestCases...) {
		runTestCase(t, tc)
	}
}

func BenchmarkSolve(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range append(testCases, errorTestCases...) {
			Solve(tc.bucketOne, tc.bucketTwo, tc.goal, tc.startBucket)
		}
	}
}

var errorTestCases = []bucketTestCase{
	{
		"Invalid first bucket size",
		0, 5, 5, "one", "one", 1, 0, true,
	},
	{
		"Invalid second bucket size",
		3, 0, 3, "one", "one", 1, 0, true,
	},
	{
		"Invalid goal amount",
		1, 1, 0, "one", "one", 0, 1, true,
	},
	{
		"Invalid start bucket name",
		3, 5, 1, "three", "one", 4, 5, true,
	},
}

func TestNewBucket(t *testing.T) {
	bucket := NewBucket(5)

	if bucket.size != 5 {
		t.Fatalf("bucket should be equal to 5")
	}

	if bucket.quantity != 0 {
		t.Fatalf("quantity should be equal to 0")
	}

	bucket.Fill()

	if bucket.quantity != 5 {
		t.Fatalf("quantity should be equal to 5")
	}

	bucket.Empty()

	if bucket.quantity != 0 {
		t.Fatalf("quantity should be equal to 0")
	}
}

func TestBucket_Pour(t *testing.T) {

	a, b := Bucket{size: 3, quantity: 2}, Bucket{size: 5, quantity: 2}
	a.Pour(&b)

	if a.quantity != 0 && b.quantity != 4 {
		t.Fatalf("A quantity is %d but should be equal to 0, B quantity is %d but should be equal to 4", a.quantity, b.quantity)
	}

	a, b = Bucket{size: 3, quantity: 0}, Bucket{size: 5, quantity: 2}
	a.Pour(&b)

	if a.quantity != 0 && b.quantity != 2 {
		t.Fatalf("A quantity is %d but should be equal to 0, B quantity is %d but should be equal to 2", a.quantity, b.quantity)
	}

	a, b = Bucket{size: 3, quantity: 3}, Bucket{size: 5, quantity: 5}
	a.Pour(&b)

	// bucket b is full
	if a.quantity != 3 && b.quantity != 5 {
		t.Fatalf("A quantity is %d but should be equal to 3, B quantity is %d but should be equal to 5", a.quantity, b.quantity)
	}

	a, b = Bucket{size: 3, quantity: 2}, Bucket{size: 5, quantity: 4}
	a.Pour(&b)

	if a.quantity != 1 && b.quantity != 5 {
		t.Fatalf("A quantity is %d but should be equal to 1, B quantity is %d but should be equal to 5", a.quantity, b.quantity)
	}

}
