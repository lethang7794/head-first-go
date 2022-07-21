package main

func main() {
	// We're not allows to have multiple functions named ToGallons,
	// so we had to write long, cumbersome functions names that incorporated the type we were converting.
	// LitersToGallons(Liters(2))
	// MillilitersToGallons(Milliliters(500))

	// But we can have multiple methods named ToGallons,
	// as long as they're defined on separate types.
	// Liters(2).ToGallons()
	// Milliliters(500).ToGallons()
}
