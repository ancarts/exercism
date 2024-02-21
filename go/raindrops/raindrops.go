package raindrops

import "fmt"

func Convert(number int) string {
	type soundFinder struct {
		key   int
		sound string
	}

	soundFinderSlice := []soundFinder{
		{key: 3, sound: "Pling"},
		{key: 5, sound: "Plang"},
		{key: 7, sound: "Plong"},
	}
	var finalSound string
	for _, v := range soundFinderSlice {
		if number%v.key == 0 {
			finalSound += v.sound
		}
	}
	if finalSound == "" {
		finalSound = fmt.Sprint(number)
	}
	return finalSound
}
