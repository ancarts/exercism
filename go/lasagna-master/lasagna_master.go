package lasagna

// TODO: define the 'PreparationTime()' function

func PreparationTime(layers []string, minutesPerlayer int) int {
	if minutesPerlayer == 0 {
		minutesPerlayer = 2
	}
	return len(layers) * minutesPerlayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodlesQ int, sauceQ float64) {
	for _, v := range layers {
		switch v {
		case "noodles":
			noodlesQ += 50
		case "sauce":
			sauceQ += 0.2
		}
	}
	return noodlesQ, sauceQ
}

// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(friendsList []string, myList []string) {
	secretIngredient := friendsList[len(friendsList)-1]
	myList = append((myList)[:len(myList)-1], secretIngredient)
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, numberOfpeople int) []float64 {
	scaledQuantities := make([]float64, len(quantities))
	for i := 0; i < len(quantities); i++ {

		scaledQuantities[i] = (quantities[i] / 2) * float64(numberOfpeople)
	}
	return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
