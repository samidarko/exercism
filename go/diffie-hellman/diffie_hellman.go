package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

// PrivateKey returns a Private Key
func PrivateKey(p *big.Int) *big.Int {
	x := new(big.Int)
	max := x.Sub(p, big.NewInt(2))
	result := x.Add(x.Rand(rand.New(rand.NewSource(time.Now().UnixNano())), max), big.NewInt(2))
	return result
}

// PublicKey returns a Public Key
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

// NewPair returns a new pair of (Private Key, Public Key)
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	private := PrivateKey(p)
	return private, PublicKey(private, p, g)
}

// SecretKey returns a Secret Key
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
