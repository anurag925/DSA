package dsa

import (
	"math/rand"
)

type RandomizedSet struct {
	s map[int]bool
}

func Constructo2() RandomizedSet {
	return RandomizedSet{make(map[int]bool)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.s[val]; !ok {
		this.s[val] = true
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.s[val]; ok {
		delete(this.s, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	l := len(this.s)
	r := rand.Int31n(int32(l))
	var c int32 = 0
	for key, _ := range this.s {
		if c == r {
			return key
		}
		c++
	}
	return -1
}
