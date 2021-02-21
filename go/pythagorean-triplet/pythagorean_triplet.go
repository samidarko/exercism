package pythagorean

type Triplet [3]int

func Range(min, max int) []Triplet {
	var triplets []Triplet

	for a := min; a <= max; a++ {
		for b := a + 1; b <= max; b++ {
			for c := b + 1; c <= max; c++ {
				if (a*a + b*b) == c*c {
					triplets = append(triplets, Triplet{a, b, c})
				}
			}
		}
	}

	return triplets
}

func Sum(p int) []Triplet {
	var c int
	var triplets []Triplet

	for a := 1; a <= p; a++ {
		for b := a + 1; b <= p; b++ {
			c = p - a - b
			if (a*a+b*b) == c*c && (a+b+c) == p {
				triplets = append(triplets, Triplet{a, b, c})
			}
		}
	}

	return triplets
}
