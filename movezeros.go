main

func MoveZeros(arr []int) []int {
	var arrnew []int
	var arrzero []int

	for _, x := range arr {
		if x == 0 {
			arrzero = append(arrzero, x)
		} else {
			arrnew = append(arrnew, x)
		}
	}
	arrnew = append(arrnew, arrzero...)
	return arrnew
}
