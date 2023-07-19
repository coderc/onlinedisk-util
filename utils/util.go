package utils

func CheckStrsIsEmpty(strs ...string) bool {
	for _, v := range strs {
		if v == "" {
			return false
		}
	}
	return true
}
