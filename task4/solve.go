package main
import "unicode"

func RemoveEven(array []int) []int {
	list := make([]int, 0)
	for _, element := range array {
		if element % 2 == 1 {
			list = append(list, element)
		}
	}
	return odd
}

func PowerGenerator(a int) func() int {
	ans := 1
	return func() int {
		ans = ans * a
		return ans
	}
}

func DifferentWordsCount(f string) int {
	set := make(map[string]bool)
	word := ""
	ans := 0
	for _, ch := range (f + " ") {
		if unicode.IsLetter(ch) {
			word = word + string(unicode.ToLower(ch))
		} else if word != "" {
			if !set[word] {
				ans = ans + 1
			}
			set[word] = true
			word = ""
		}
	}
	return ans
}


