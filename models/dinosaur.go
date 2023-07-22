package models

type Dinosaur struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Species        string `json:"species"`
	FoodPreference string `json:"food_preference"`
	CageID         *int   `json:"cage_id,omitempty"`
}

func GetDinosaur(id int) (dinosaur Dinosaur, err error) {
	err = DB.QueryRow("SELECT id, name, species, food_preference, dc.cage_id FROM dinosaurs d LEFT JOIN dinosaurs_in_cages dc ON (d.id = dc.dinosaur_id) WHERE id = ?", id).Scan(&dinosaur.ID, &dinosaur.Name, &dinosaur.Species, &dinosaur.FoodPreference, &dinosaur.CageID)
	return dinosaur, err
}

func GetDinosaurs() (dinosaurs []Dinosaur, err error) {
	rows, err := DB.Query("SELECT id, name, species, food_preference, dc.cage_id FROM dinosaurs d LEFT JOIN dinosaurs_in_cages dc ON (d.id = dc.dinosaur_id) ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var d Dinosaur
		err = rows.Scan(&d.ID, &d.Name, &d.Species, &d.FoodPreference, &d.CageID)
		if err != nil {
			return nil, err
		}
		dinosaurs = append(dinosaurs, d)
	}
	return dinosaurs, nil
}

func AddDinosaur(dinosaur Dinosaur) (Dinosaur, error) {
	res, err := DB.Exec("INSERT INTO dinosaurs (name, species, food_preference) VALUES (?, ?, ?)", dinosaur.Name, dinosaur.Species, dinosaur.FoodPreference)
	if err != nil {
		return Dinosaur{}, err
	}
	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return Dinosaur{}, err
	}
	return GetDinosaur(int(id))
}

func UpdateDinosaur(dinosaur Dinosaur) (Dinosaur, error) {
	_, err := DB.Exec("UPDATE dinosaurs SET name = ?, species = ?, food_preference = ? WHERE id = ?", dinosaur.Name, dinosaur.Species, dinosaur.FoodPreference, dinosaur.ID)
	if err != nil {
		return Dinosaur{}, err
	}
	return GetDinosaur(dinosaur.ID)
}

func DeleteDinosaur(id int) error {
	_, err := GetDinosaur(id)
	if err != nil {
		return err
	}
	_, err = DB.Exec("DELETE FROM dinosaurs WHERE id = ?", id)
	return err
}
