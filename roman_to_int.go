package dsa

// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
var h = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

var d = map[string]int{
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

func romanToInt(s string) int {
	l := len(s)
	res := 0
	ignore := false
	for i, e := range s {
		if ignore {
			ignore = false
			continue
		}
		if (i + 1) < l {
			if val, ok := d[string(e)+string(s[i+1])]; ok {
				res += val
				ignore = true
				continue
			}
		}
		res += h[e]
	}
	return res
}
