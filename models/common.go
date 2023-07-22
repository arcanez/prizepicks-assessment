package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return err
	}

	db.Exec("CREATE TABLE dinosaurs (id integer NOT NULL PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, species TEXT NOT NULL, food_preference TEXT CHECK (food_preference IN ('carnivore', 'herbivore')) NOT NULL)")
	db.Exec("CREATE TABLE cages (id integer NOT NULL PRIMARY KEY AUTOINCREMENT, power_status TEXT CHECK (power_status IN ('ACTIVE', 'DOWN')) NOT NULL DEFAULT 'DOWN', maximum_capacity integer NOT NULL DEFAULT 1)")
	db.Exec("CREATE TABLE dinosaurs_in_cages (cage_id integer NOT NULL, dinosaur_id integer NOT NULL, CONSTRAINT cage_id_fk FOREIGN KEY (cage_id) REFERENCES cages (id), CONSTRAINT dinosaurs_id_fk FOREIGN KEY (dinosaur_id) REFERENCES dinosaurs (id) ON DELETE CASCADE)")

	db.Exec(`INSERT INTO dinosaurs (name, species, food_preference) VALUES ('Al', 'Tyrannosaurus', 'carnivore'), ('Bob', 'Velociraptor', 'carnivore'), ('Stewart', 'Spinosaurus', 'carnivore'), ('Ralph', 'Megalosaurus','carnivore'),
	('Ernie', 'Brachiosaurus', 'herbivore'), ('Harvey', 'Stegosaurus', 'herbivore'), ('Mike', 'Ankylosaurus', 'herbivore'), ('Harold', 'Triceratops', 'herbivore')`)

	db.Exec("INSERT INTO cages (power_status, maximum_capacity) VALUES ('ACTIVE', 1), ('ACTIVE', 5), ('DOWN', 5)")

	DB = db
	return nil
}

func ConnectTestDatabase() error {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return err
	}

	db.Exec("CREATE TABLE dinosaurs (id integer NOT NULL PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, species TEXT NOT NULL, food_preference TEXT CHECK (food_preference IN ('carnivore', 'herbivore')) NOT NULL)")
	db.Exec("CREATE TABLE cages (id integer NOT NULL PRIMARY KEY AUTOINCREMENT, power_status TEXT CHECK (power_status IN ('ACTIVE', 'DOWN')) NOT NULL DEFAULT 'DOWN', maximum_capacity integer NOT NULL DEFAULT 1)")
	db.Exec("CREATE TABLE dinosaurs_in_cages (cage_id integer NOT NULL, dinosaur_id integer NOT NULL, CONSTRAINT cage_id_fk FOREIGN KEY (cage_id) REFERENCES cages (id), CONSTRAINT dinosaurs_id_fk FOREIGN KEY (dinosaur_id) REFERENCES dinosaurs (id) ON DELETE CASCADE)")

	db.Exec(`INSERT INTO dinosaurs (name, species, food_preference) VALUES ('Al', 'Tyrannosaurus', 'carnivore'), ('Bob', 'Velociraptor', 'carnivore'), ('Stewart', 'Spinosaurus', 'carnivore'), ('Ralph', 'Megalosaurus','carnivore'),
	('Ernie', 'Brachiosaurus', 'herbivore'), ('Harvey', 'Stegosaurus', 'herbivore'), ('Mike', 'Ankylosaurus', 'herbivore'), ('Harold', 'Triceratops', 'herbivore')`)

	db.Exec("INSERT INTO cages (power_status, maximum_capacity) VALUES ('ACTIVE', 1), ('ACTIVE', 5), ('DOWN', 5)")

	DB = db
	return nil
}
