package palindrome

import (
	"fmt"
	"math"
	"strconv"
)

// Product type
type Product struct {
	Palindrome     int
	Factorizations [][2]int
}

// UpdateFactorization  and prevent from duplicates
func (p *Product) UpdateFactorization(factorization [2]int) {
	for _, f := range p.Factorizations {
		if f[0] == factorization[1] && f[1] == factorization[0] {
			// duplicate
			return
		}
	}
	p.Factorizations = append(p.Factorizations, factorization)
}

func Products(fmin, fmax int) (Product, Product, error) {
	if fmin > fmax {
		return Product{}, Product{}, fmt.Errorf("fmin > fmax")
	}

	minProduct := Product{math.MaxInt, make([][2]int, 0)}
	maxProduct := Product{math.MinInt, make([][2]int, 0)}

	for i := fmin; i <= fmax; i++ {
		for j := fmin; j <= fmax; j++ {

			product := i * j
			productString := strconv.Itoa(product)

			if productString == Reverse(productString) {
				// this is a Palindrome
				switch {
				case product < minProduct.Palindrome:
					minProduct = Product{
						Palindrome:     product,
						Factorizations: [][2]int{{i, j}},
					}
				case product == minProduct.Palindrome:
					minProduct.UpdateFactorization([2]int{i, j})
				case product > maxProduct.Palindrome:
					maxProduct = Product{
						Palindrome:     product,
						Factorizations: [][2]int{{i, j}},
					}
				case product == maxProduct.Palindrome:
					maxProduct.UpdateFactorization([2]int{i, j})
				}
			}
		}
	}

	if minProduct.Palindrome == math.MaxInt || maxProduct.Palindrome == math.MinInt {
		return Product{}, Product{}, fmt.Errorf("no palindromes")
	}

	return minProduct, maxProduct, nil
}

// Reverse a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
