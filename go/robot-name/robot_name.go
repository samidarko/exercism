package robotname

import (
	"fmt"
)

// Define the Robot type here.

type Robot struct {
	name string
}

func nameGenerator() func() (string, error) {
	var names []string
	for c1 := 'A'; c1 <= 'Z'; c1++ {
		for c2 := 'A'; c2 <= 'Z'; c2++ {
			for n := 0; n <= 999; n++ {
				name := string([]rune{c1, c2}) + fmt.Sprintf("%03d", n)
				names = append(names, name)
			}
		}
	}
	index := 0
	return func() (string, error) {
		if index < len(names) {
			name := names[index]
			index++
			return name, nil
		}
		return "", fmt.Errorf("no more names")
	}
}

var generateName = nameGenerator()

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		name, err := generateName()
		if err != nil {
			return "", err
		}
		r.name = name
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
