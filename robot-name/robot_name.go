package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var robots = make(map[string]bool, 676000)
var attempt = 0

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	name, err := generateName()

	if err != nil {
		return "", err
	}

	r.name = name

	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

func generateName() (string, error) {
	rand.Seed(time.Now().UnixNano())

	candidate := fmt.Sprintf(
		"%s%s%d%d%d",
		string(letters[rand.Intn(len(letters))]),
		string(letters[rand.Intn(len(letters))]),
		rand.Intn(10),
		rand.Intn(10),
		rand.Intn(10),
	)

	if robots[candidate] {
		if attempt > 676000 {
			return "", fmt.Errorf("name collision")
		}
		attempt++

		return generateName()
	}

	robots[candidate] = true
	attempt = 0

	return candidate, nil
}
