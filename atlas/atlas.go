package atlas

func NewWorld(size int, density int, seed int64) *World {
	world := newWorld(size, density, seed)
	return world
}

//"deepWater" : 0,
//"water" : 1,
//"grass" : 2,
//dirt : 3
//"sand" : 4,
//"sandstone" : 5,
//"dryGrass" : 6,
//"coals" : 7,
//"ruins" : 8,
//"snow" : 9,
//ice : 10

func TemplateWorld(size int) *World {
	biomes := []Biome{
		Biome{
			[]Modifier{
				NewCropCircle(6.3, 7,9),
				NewSelectiveBorder(9,10),
				NewSelectiveBorder(8,7),
				NewBorder(10),
			},
		},
		Biome{
			[]Modifier{
				NewVoronoi(40, 7,8,8,9,10),
				NewSelectiveBorder(8,9),
				NewSelectiveBorder(9,10),
				NewBorder(9),
			},
		},
		Biome{
			[]Modifier{
				NewVoronoi(40, 7,7,8,8,8,8,8,9),
				NewSelectiveBorder(8,9),
				NewBorder(9),
			},
		},
		Biome{
			[]Modifier{
				NewVoronoi(40, 7,7,8,8,9),
				NewSelectiveBorder(8,9),
				NewBorder(8),
			},
		},
		Biome{
			[]Modifier{
				NewCropCircle(30, 4,5),
			},
		},
		Biome{
			[]Modifier{
				NewVoronoi(40, 4,5,6),
				NewSelectiveBorder(5,6),
			},
		},
		Biome{
			[]Modifier{
				NewCropCircle(4.7, 6,4),
				NewSelectiveBorder(5,4),
			},
		},
		Biome{
			[]Modifier{
				NewPattern(3.2, 6,5),
				NewBorder(5),
			},
		},
		Biome{
			[]Modifier{
				NewPattern(3.2, 2,8),
				NewBorder(8),
			},
		},
		Biome{
			[]Modifier{
				NewCropCircle(20, 1,2),
				NewSelectiveBorder(8,1),
			},
		},
	}
	
	world := newWorld(size, size, 5)
	world.infect(biomes, 0.7)
	
	return world
}

func NewBiomes(biomes int8) []Biome {
	biome := make([]Biome, biomes)
	
	for idx := range biome {
		biome[idx] = Biome{[]Modifier{NewBase()}}
	}
	
	return biome
}

