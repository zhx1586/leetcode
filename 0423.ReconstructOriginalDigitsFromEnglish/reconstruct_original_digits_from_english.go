package reconstructoriginaldigitsfromenglish

import "strconv"

func originalDigits(s string) string {
	code := []string{
		"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}
	counts := map[byte]int{
		'e': 0, // zero, one, three, five, seven, eight, nine
		'g': 0, // eight
		'f': 0, // four, five
		'i': 0, // five, six, eight, nine
		'h': 0, // three, eight
		'o': 0, // zero, one, two, four
		'n': 0, // one, seven, nine
		's': 0, // six, seven
		'r': 0, // zero, three, four
		'u': 0, // four
		't': 0, // two, three, eight
		'w': 0, // two
		'v': 0, // five, seven
		'x': 0, // six
		'z': 0, // zero
	}
	digits := make([]int, 10)

	for i := 0; i < len(s); i++ {
		counts[s[i]]++
	}

	seq := []struct {
		letter byte
		num    int
	}{
		{'z', 0},
		{'w', 2},
		{'u', 4},
		{'x', 6},
		{'g', 8},
		{'s', 7},
		{'v', 5},
		{'r', 3},
		{'o', 1},
		{'i', 9},
	}
	for _, s := range seq {
		digits[s.num] = counts[s.letter]
		for i := 0; i < len(code[s.num]); i++ {
			counts[code[s.num][i]] = counts[code[s.num][i]] - digits[s.num]
		}
	}

	ret := ""
	for i := 0; i < 10; i++ {
		for j := 0; j < digits[i]; j++ {
			ret = ret + strconv.Itoa(i)
		}
	}

	return ret
}
