package dsa

import (
	"fmt"
	"slices"
)

var o, z byte = '1', '0'

func addBinary(a string, b string) string {
	c := z
	al, bl := len(a), len(b)
	if al < bl {
		tmp := a
		a = b
		b = tmp
	}
	al, bl = len(a), len(b)
	ac, bc := al-1, bl-1

	res := []byte("")
	for bc >= 0 {
		fmt.Println(a[ac], b[ac], c)
		var add, carry byte = a[ac], c
		if carry != z {
			carry, add = adds(add, carry)
			c = carry
		}
		carry, add = adds(add, b[bc])
		if carry != z {
			c = carry
		}
		res = append(res, add)
		ac--
		bc--
	}
	var add byte
	for ac >= 0 {
		fmt.Println(a[ac], c)
		add, c = adds(a[ac], c)
		res = append(res, add)
		ac--
	}
	if c == o {
		res = append(res, c)
	}
	fmt.Println(res)
	slices.Reverse(res)
	return string(res)
}

func adds(a, b byte) (byte, byte) {
	if a == o && b == o {
		return o, z
	} else if a == o && b == z {
		return z, o
	} else if a == z && b == o {
		return z, o
	}
	return z, z
}
