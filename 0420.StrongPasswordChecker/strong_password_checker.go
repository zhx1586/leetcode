package strongpasswordchecker

func strongPasswordChecker(password string) int {
	ins, del := 0, 0

	upper, lower, digit := false, false, false
	for i := 0; i < len(password); i++ {
		if i > 2 && password[i] == password[i-1] && password[i] == password[i-2] {
			del++
		}
		if password[i] >= 'A' && password[i] <= 'Z' {
			upper = true
		} else if password[i] >= 'a' && password[i] <= 'z' {
			lower = true
		} else if password[i] >= '0' && password[i] <= '9' {
			digit = true
		}
	}
	if !upper {
		ins++
	}
	if !lower {
		ins++
	}
	if !digit {
		ins++
	}
	if len(password) > 20 {
		del = max(del, len(password)-20)
	} else if len(password) < 6 {
		ins = max(ins, 6-len(password))
	}

	return abs(ins - del)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
