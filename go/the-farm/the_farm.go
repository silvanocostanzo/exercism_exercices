package thefarm

import "errors"

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()

	if err != nil {
		switch err.Error() {
		case ErrScaleMalfunction.Error():
			if fodder < 0 {
				return 0, errors.New("negative fodder")
			}
			return fodder * 2 / float64(cows), nil
		default:
			return 0, err
		}
	} else {
		if fodder < 0 {
			return 0, errors.New("negative fodder")
		}
		if cows == 0 {
			return 0, errors.New("division by zero")
		}
		return fodder / float64(cows), nil
	}
}
