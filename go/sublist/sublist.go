package sublist

//Relation gives relation between 2 lists
type Relation string

const equal Relation = "equal"
const sublist Relation = "sublist"
const superlist Relation = "superlist"
const unequal Relation = "unequal"

//Sublist returns the relation between the provided lists
func Sublist(l1, l2 []int) Relation {
	size := len(l1) - len(l2)
	if size == 0 { // sizes are equal
		for i, v := range l1 {
			if v != l2[i] {
				return unequal
			}
		}
		return equal

	} else if size > 0 { // size of l1 > size of l2
		if len(l2) == 0 {
			return superlist
		}
		for i := 0; i <= size; i++ {
			if l1[i] == l2[0] {
				superlistFound := true
				for j := 0; j < len(l2); j++ {
					if l1[i+j] != l2[j] {
						superlistFound = false
						break
					}
				}
				if superlistFound {
					return superlist
				}
			}
		}
	} else if size < 0 { // size of l1 < size of l2
		if len(l1) == 0 {
			return sublist
		}
		for i := 0; i <= size*-1; i++ {
			if l2[i] == l1[0] {
				sublistFound := true
				for j := 0; j < len(l1); j++ {
					if l2[i+j] != l1[j] {
						sublistFound = false
						break
					}
				}
				if sublistFound {
					return sublist
				}
			}
		}

	}
	return unequal
}
