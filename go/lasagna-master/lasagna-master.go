package lasagna

func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		return len(layers) * 2
	}
	return len(layers) * timePerLayer
}

func Quantities(layers []string) (gotNoodles int, gotSauce float64) {
	for _, l := range layers {
		if l == "noodles" {
			gotNoodles += 50
		} else if l == "sauce" {
			gotSauce += 0.2
		}
	}
	return
}

func AddSecretIngredient(friendsList, myList []string) []string {
	myList = append(myList, friendsList[len(friendsList)-1])
	return myList
}

func ScaleRecipe(quantities []float64, portions int) (newQuantities []float64) {
	for _, q := range quantities {
		newQuantities = append(newQuantities, (q * float64(portions) / 2))
	}
	return
}
