package database

type Database struct{}

type Game struct {
	Name        string
	ReleaseYear int
	Platforms   []string
	Publisher   string
}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) GetAllGames() []Game {
	game1 := Game{Name: "Super Mario Bros.", ReleaseYear: 1985, Platforms: []string{"NES"}, Publisher: "Nintendo"}
	game2 := Game{Name: "The Legend of Zelda", ReleaseYear: 1986, Platforms: []string{"NES", "Switch"}, Publisher: "Nintendo"}
	games := []Game{game1, game2}
	return games
}
