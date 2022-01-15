package sublist

type Relation = string

func Sublist(l1, l2 []int) Relation {
	if len(l1) == len(l2) && equal(l1, l2) {
		return "equal"
	}
	if len(l1) < len(l2) && isSublist(l1, l2) {
		return "sublist"
	}
	if len(l1) > len(l2) && isSublist(l2, l1) {
		return "superlist"
	}
	return "unequal"
}

func isSublist(l1, l2 []int) bool {
	if len(l1) == 0 {
		return true
	}
	for _, i := range l1 {
		for offset, j := range l2 {
			if offset+len(l1) > len(l2) {
				// l1 does not fit anymore in l2
				return false
			}
			if i == j && equal(l1, l2[offset:offset+len(l1)]) {
				return true
			}
		}
	}
	return false
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
