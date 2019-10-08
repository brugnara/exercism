package robotname

import (
	"math/rand"
	"strings"
	"time"
)

// Robot struct
type Robot struct {
	name string
}

// Name returns a name
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.genName()
	}
	return r.name, nil
}

// Reset resets the name
func (r *Robot) Reset() {
	r.genName()
}

func (r *Robot) genName() {
	r.name = ""
	ar1 := strings.Split("QWERTYUIOPLKJHGFDSAZXCVBNM", "")
	ar2 := strings.Split("1234567890", "")
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < 2; i++ {
		r.name += ar1[r1.Intn(len(ar1))]
	}
	for i := 0; i < 3; i++ {
		r.name += ar2[r1.Intn(len(ar2))]
	}
}
