package allergies

import "fmt"

type allergy struct {
	what   string
	points uint
}

// do not use a map[string]uint and a range, because the tests will fail
// due the impossibility to have an ordered range..
var allergies = []allergy{
	allergy{
		what:   "eggs",
		points: 1,
	},
	allergy{
		what:   "peanuts",
		points: 2,
	},
	allergy{
		what:   "shellfish",
		points: 4,
	},
	allergy{
		what:   "strawberries",
		points: 8,
	},
	allergy{
		what:   "tomatoes",
		points: 16,
	},
	allergy{
		what:   "chocolate",
		points: 32,
	},
	allergy{
		what:   "pollen",
		points: 64,
	},
	allergy{
		what:   "cats",
		points: 128,
	},
}

// Allergies comment
func Allergies(score uint) (ret []string) {
	for _, allergy := range allergies {
		// fmt.Println(allergy, sc)
		bitWiseResult := score & allergy.points
		fmt.Printf("%v&%v = %v\n", score, allergy.points, bitWiseResult)
		if bitWiseResult != 0 {
			fmt.Printf("Allergic to %s\n", allergy.what)
			ret = append(ret, allergy.what)
		} else {
			fmt.Printf("Not allergic to %s\n", allergy.what)
		}
	}
	fmt.Println(ret)
	return ret
}

// AllergicTo comment
func AllergicTo(score uint, what string) bool {
	for _, allergy := range allergies {
		if allergy.what == what {
			return score&allergy.points != 0
		}
	}
	return false
}
