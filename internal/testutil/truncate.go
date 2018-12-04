package testutil

// Truncate a string by n
func Truncate(str string, num int) string {
	s := str
	if len(str) > num {
		s = str[0:num]
	}
	return s
}
