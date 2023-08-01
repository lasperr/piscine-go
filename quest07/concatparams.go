package piscine

func ConcatParams(args []string) string {
	result := ""
	if len(args) > 0 {
		for index, str := range args {
			if index == len(args)-1 {
				result += str
			} else {
				result += str + "\n"
			}
		}
	}
	return result
}
