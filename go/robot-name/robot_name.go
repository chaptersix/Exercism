package robotname

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

//Robot represents  a robot
type Robot struct {
	name string
}

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const max = 26 * 26 * 10 * 10 * 10

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

var exists = map[string]bool{}

//Name return the name of the robot
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		name := ""
		for {
			name = randName()
			if !exists[name] {
				break
			}
			if len(exists) >= max {
				return "", errors.New("Max robots reached")
			}
		}
		r.name = name
		exists[r.name] = true
	}
	return r.name, nil
}

//Reset reassigns a name
func (r *Robot) Reset() {
	name := ""
	for {
		name = randName()
		if !exists[name] {
			break
		}
		if len(exists) >= max {
			return
		}
	}
	r.name = name
	exists[r.name] = true
}

func randName() string {
	b := make([]byte, 2)
	for i := 0; i < 2; i++ {
		b[i] = charset[seededRand.Intn(26)]
	}
	reStr := string(b)
	for i := 2; i < 5; i++ {
		reStr += strconv.Itoa(seededRand.Intn(10))
	}
	return reStr
}
