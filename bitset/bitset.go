//Package bitset implements uint64 bitvector
package bitset

import (
	"fmt"
	"math/bits"
)

//Bitset is uint64 bitvector primitive
type Bitset struct {
	vec uint64
}

//New bitset
func New() *Bitset {
	return &Bitset{vec: 0}
}

//Ones count
func (b *Bitset) Ones() int {
	return bits.OnesCount64(b.vec)
}

//Set one on position p
func (b *Bitset) Set(p int) {
	b.vec |= (1 << uint64(p))
}

//Clear bit (set to zerp at position p)
func (b *Bitset) Clear(p int) {
	b.vec &= ^(1 << uint64(p))
}

//Has a bit set at position p
func (b *Bitset) Has(p int) bool {
	out := b.vec & (1 << uint64(p))
	fmt.Println(out)
	return out > 0
}

//CountIntersect count of intersection betwen two vectors
func (b *Bitset) CountIntersect(c *Bitset) int {
	return bits.OnesCount64(b.vec & c.vec)
}

//String vector
func (b *Bitset) String() string {
	return fmt.Sprintf("%b", b.vec)
}
