package basics

// A return statement without arguments returns the named return values.
// This is known as a "naked" return.
// A function can return any number of results.
// A return statement without arguments returns the named return values.
// This is known as a "naked" return.
// A function can return any number of results.
func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
