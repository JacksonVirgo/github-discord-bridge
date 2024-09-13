package utils

func UnorderedEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	exists := make(map[string]bool)
	for _, value := range a {
		exists[value] = true
	}
	for _, value := range b {
		if !exists[value] {
			return false
		}
	}
	return true
}

func Difference(a, b []string) []string {
	diff := []string{}
	existsInB := make(map[string]bool)
	for _, tag := range b {
		existsInB[tag] = true
	}
	for _, tag := range a {
		if !existsInB[tag] {
			diff = append(diff, tag)
		}
	}

	return diff
}
