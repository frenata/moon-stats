package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"

	"github.com/frenata/moon-stats/lib"
	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
)

func boot() *sql.DB {
	const file string = "games.db"
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}

	init, _ := os.Open("sql/init.sql")
	bs, _ := ioutil.ReadAll(init)
	_, err = db.Exec(string(bs))
	if err != nil {
		log.Fatal(err)
	}

	units, _ := os.Open("sql/units.sql")
	bs, _ = ioutil.ReadAll(units)
	_, err = db.Exec(string(bs))
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func (app *App) upsertPlayer(id uuid.UUID, name string) {
	var existing uuid.UUID
	err := app.db.QueryRow("select id from players where display_name = ?", name).Scan(&existing)
	if err == nil {
		return
	} else if err != nil && err != sql.ErrNoRows {
		log.Fatal("query players ", err)
	}

	stmt, err := app.db.Prepare("insert into players(id, display_name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(id, name)
	if err != nil {
		log.Fatal(err)
	}
}

func (app *App) upsertGame(player_id, game_id uuid.UUID, victory bool, units []lib.Unit) {
	var existing uuid.UUID
	err := app.db.QueryRow("select id from games where id = ?", game_id).Scan(&existing)
	if err == nil {
		return
	} else if err != nil && err != sql.ErrNoRows {
		log.Fatal("query games ", err)
	}

	tx, err := app.db.Begin()
	if err != nil {
		log.Fatal("begin tx insert games ", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("insert into games(id, player0_id, victory) values(?, ?, ?)")
	if err != nil {
		log.Fatal("prep insert games ", err)
	}
	_, err = stmt.Exec(game_id, player_id, victory)
	if err != nil {
		log.Fatal("insert games ", err)
	}

	stmt, err = tx.Prepare("insert into game_units(game_id, player_id, unit_number, unit_type, played, died) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal("prep insert unit ", err)
	}
	for _, unit := range units {
		_, err = stmt.Exec(game_id, player_id, unit.Id, unit.Type, unit.Spawned, unit.Died)
		if err != nil {
			log.Fatal("insert unit ", err)
		}
	}
	tx.Commit()
}
