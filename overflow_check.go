package dsa

import (
	"fmt"
	"math"
)

func CheckOverflow() {
	i := int64(9223372036854775807)
	j := int64(2345433343242343)
	k := i * j
	fmt.Println(k)
	fmt.Println(int(k), math.MaxInt, math.MaxInt64, math.MaxInt32)
}
