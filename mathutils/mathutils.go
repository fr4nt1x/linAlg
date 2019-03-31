package mathutils

func Prod(values []uint) uint {
	result := values[0]
	for _, val := range values[1:] {
		result *= val
	}
	return result
}
