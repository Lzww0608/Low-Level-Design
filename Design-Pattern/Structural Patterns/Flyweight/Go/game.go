package main

type game struct {
	terrorists []*Player
	counterTerrorists []*Player
}

func newGame() *game {
	return &game {
		terrorists: make([]*Player, 1),
		counterTerrorists: make([]*Player, 1),
	}
}

func (g *game) addTerrorist(dressType string) {
	g.terrorists = append(g.terrorists, newPlayer("Terrorist", dressType))
	return 
}

func (g *game) addCounterTerrorist(dressType string) {
	g.counterTerrorists = append(g.counterTerrorists, newPlayer("CounterTerrorist", dressType))
	return 
}





