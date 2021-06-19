package strongpasswordchecker

func strongPasswordChecker(password string) int {
	//fmt.Println(password)
	ins, del := 0, 0

	upper, lower, digit := false, false, false
	for i := 0; i < len(password); i++ {
		//fmt.Printf("i:%d, del:%d\n", i, del)
		if i >= 2 && password[i] == password[i-1] && password[i] == password[i-2] {
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
	length := len(password) + ins - del

	if length > 20 {
		del = del + (length - 20)
	} else if length < 6 {
		ins = ins + (6 - length)
	}
	//fmt.Printf("length:%d, ins:%d, del:%d\n", length, ins, del)

	return max(ins, del)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
