package sublist

type Relation string

func Sublist(A, B []int) Relation {
	inverted := false
	if len(A) == 0 && len(B) == 0 {
		return Relation("equal")
	}
	if len(A) == 0 {
		return Relation("sublist")
	}
	if len(B) == 0 {
		return Relation("superlist")
	}
	if equals(A, B) {
		return Relation("equal")
	}
	// keep track we inverted A with B
	if len(A) > len(B) {
		inverted = true
		A, B = B, A
	}
	starts := []int{}
	for i, b := range B {
		if b == A[0] {
			starts = append(starts, i)
		}
	}

	if len(starts) == 0 {
		return Relation("unequal")
	}

	// for each valid start, iterate to check equals
	for _, s := range starts {
		sublist := true
		for i, a := range A {
			if s+i >= len(B) || B[s+i] != a {
				sublist = false
				break
			}
		}
		if sublist {
			if inverted {
				return Relation("superlist")
			}
			return Relation("sublist")
		}
	}
	return Relation("unequal")
}

func equals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, aa := range a {
		if b[i] != aa {
			return false
		}
	}
	return true
}
