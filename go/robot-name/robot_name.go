package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//Robot represents  a robot
type Robot struct {
	name string
}

const max = 26 * 26 * 10 * 10 * 10

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
var exists = map[string]bool{}

//Name return the name of the robot
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if len(exists) >= max {
		return "", errors.New("Max robots reached")
	}
	name := randName()
	for exists[name] {
		name = randName()
	}
	r.name = name
	exists[r.name] = true

	return r.name, nil
}

//Reset reassigns a name
func (r *Robot) Reset() {
	r.name = ""
}

func randName() string {
	r1 := string(seededRand.Intn(26) + 'A')
	r2 := string(seededRand.Intn(26) + 'A')
	num := seededRand.Intn(1000)
	return fmt.Sprintf("%s%s%03d", r1, r2, num)
}
