package sublist

type Relation string

func Sublist(a, b []int) Relation {
	if len(a) == len(b) && contains(a, b) {
		return "equal"
	} else if len(a) > len(b) && contains(a, b) {
		return "superlist"
	} else if contains(b, a) {
		return "sublist"
	}

	return "unequal"
}

func contains(a, b []int) bool {
	for i := 0; i <= len(a)-len(b); i++ {
		equal := true

		for j := 0; j < len(b); j++ {
			if b[j] != a[i+j] {
				equal = false
				break
			}
		}

		if equal {
			return true
		}
	}

	return false
}
