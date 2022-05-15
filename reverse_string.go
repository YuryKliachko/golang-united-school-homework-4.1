package reverse_string

func ReverseString(input string) (output string) {
	reversed := ""
	for _, c := range input {
		reversed = string(c) + reversed
	}
	return reversed
}
