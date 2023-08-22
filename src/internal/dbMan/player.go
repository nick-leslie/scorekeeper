package dbMan

import (
	"database/sql"
	"github.com/google/uuid"
)

type Player struct {
	Id   string
	Name string
}

func NewPlayer(id string, name string) *Player {
	return &Player{Id: id, Name: name}
}

func AddPlayerToDb(name string) (int64, error) {
	db := Con.DB
	ret, err := db.Exec("INSERT INTO main.player (id,NAME) VALUES (?,?)", uuid.New(), name)
	if err != nil {
		return 0, err
	}
	affected, err := ret.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affected, nil
}
func GetPlayers() []*Player {
	var players []*Player
	rows, err := Con.DB.Query("SELECT id,name FROM main.player")
	if err != nil {
		println(err.Error())
		return nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			println(err.Error())
			return
		}
	}(rows)
	for rows.Next() {
		player := NewPlayer("", "")
		err := rows.Scan(&player.Id, &player.Name)
		if err != nil {
			return nil
		}
		players = append(players, player)
	}
	return players
}
func GetPlayer(name string) {
	Con.DB.QueryRow("SELECT id,name FROM main.player WHERE NAME = ?", name)
}
func GetPlayerFromId(id string) {
	Con.DB.QueryRow("SELECT id,name FROM main.player WHERE id = ?", id)
}
