package twobucket

import (
	"errors"
	"fmt"
)

type Bucket struct {
	size     int
	quantity int
	numSteps int
}

func NewBucket(size int) Bucket {
	return Bucket{size: size}
}

func (b *Bucket) Empty() {
	b.quantity = 0
	b.numSteps++
}

func (b *Bucket) Fill() {
	b.quantity = b.size
	b.numSteps++
}

func (b *Bucket) Pour(into *Bucket) {
	for b.quantity > 0 && into.quantity < into.size {
		b.quantity--
		into.quantity++
	}
	b.numSteps++
}

func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (string, int, int, error) {
	if sizeBucketOne <= 0 || sizeBucketTwo <= 0 || goalAmount <= 0 {
		return "", 0, 0, errors.New("negative or zero not accepted for bucket size or goal amount")
	}
	if startBucket != "one" && startBucket != "two" {
		return "", 0, 0, fmt.Errorf(`start bucket can only be "one" or "two" and you gave "%s"`, startBucket)
	}

	if goalAmount%(gcd(sizeBucketOne, sizeBucketTwo)) != 0 {
		return "", 0, 0, errors.New("cannot be solved")
	}

	buckets := map[string]Bucket{
		"one": NewBucket(sizeBucketOne),
		"two": NewBucket(sizeBucketTwo),
	}
	currentBucketName := startBucket
	currentBucket := buckets[currentBucketName]
	currentBucket.Fill()

	otherBucket := buckets[nextBucket(currentBucketName)]
	if otherBucket.size == goalAmount {
		otherBucket.Fill()
	}

	for currentBucket.quantity != goalAmount && otherBucket.quantity != goalAmount {
		switch {
		case currentBucket.quantity == 0:
			currentBucket.Fill()
		case otherBucket.quantity == otherBucket.size:
			otherBucket.Empty()
		default:
			currentBucket.Pour(&otherBucket)
		}
	}

	numSteps := currentBucket.numSteps + otherBucket.numSteps
	if currentBucket.quantity == goalAmount {
		return currentBucketName, numSteps, otherBucket.quantity, nil
	}

	return nextBucket(currentBucketName), numSteps, currentBucket.quantity, nil
}

func nextBucket(bucketName string) string {
	if bucketName == "one" {
		return "two"
	}
	return "one"
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
