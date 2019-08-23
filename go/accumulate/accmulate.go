package accumulate

//Accumulate calls the provided converter to each element of the given array
func Accumulate(given []string, converter func(string) string) []string {
	for index := 0; index < len(given); index++ {
		given[index] = converter(given[index])
	}
	return given
}
