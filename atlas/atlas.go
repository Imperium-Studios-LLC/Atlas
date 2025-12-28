package atlas

const DENSITY_LOW = 10
const DENSITY_MED = 100
const DENSITY_HIGH = 500

func GenerateWorld(size int) *World {
	world := newWorld(DENSITY_MED, size, 100)
	return world
}



