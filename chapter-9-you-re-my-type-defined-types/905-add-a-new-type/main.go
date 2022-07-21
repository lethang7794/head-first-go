package main

type Liters float64
type Milliliters float64
type Gallons float64

func LitersToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func MillilitersToGallons(ml Milliliters) Gallons {
	return Gallons(ml * 0.000264)
}

func GallonsToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func GallonsToMilliliters(g Gallons) Milliliters {
	return Milliliters(g * 3785.41)
}

func LitersToMilliliters(l Liters) Milliliters {
	return Milliliters(l * 1000)
}

func MillilitersToLiters(l Milliliters) Liters {
	return Liters(l * 0.001)
}
