package geo

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

func (c *Coordinates) SetLatitude(latitude float64) {
	c.Latitude = latitude
}

func (c *Coordinates) SetLongitude(longitude float64) {
	c.Longitude = longitude
}
