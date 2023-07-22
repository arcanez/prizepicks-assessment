package models

import "fmt"

type Cage struct {
	ID              int        `json:"id"`
	PowerStatus     string     `json:"power_status"`
	MaximumCapacity int        `json:"maximum_capacity"`
	Dinosaurs       []Dinosaur `json:"dinosaurs"`
}

func GetCage(id int) (cage Cage, err error) {
	if err = DB.QueryRow("SELECT id, power_status, maximum_capacity FROM cages WHERE id = ?", id).Scan(&cage.ID, &cage.PowerStatus, &cage.MaximumCapacity); err != nil {
		return Cage{}, err
	}
	rows, err := DB.Query("SELECT id, name, species, food_preference FROM dinosaurs d JOIN dinosaurs_in_cages dc ON (d.id = dc.dinosaur_id) WHERE dc.cage_id = ?", id)
	if err != nil {
		return Cage{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var d Dinosaur
		err = rows.Scan(&d.ID, &d.Name, &d.Species, &d.FoodPreference)
		if err != nil {
			return Cage{}, err
		}
		cage.Dinosaurs = append(cage.Dinosaurs, d)
	}
	return cage, err
}

func GetCages() (cages []Cage, err error) {
	rows, err := DB.Query("SELECT id, power_status, maximum_capacity FROM cages ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var cage Cage
		err = rows.Scan(&cage.ID, &cage.PowerStatus, &cage.MaximumCapacity)
		if err != nil {
			return nil, err
		}
		cage.Dinosaurs = []Dinosaur{}
		cages = append(cages, cage)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}

	for i := range cages {
		rows, err := DB.Query("SELECT id, name, species, food_preference FROM dinosaurs d JOIN dinosaurs_in_cages dc ON (d.id = dc.dinosaur_id) WHERE dc.cage_id = ?", cages[i].ID)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var d Dinosaur
			err = rows.Scan(&d.ID, &d.Name, &d.Species, &d.FoodPreference)
			if err != nil {
				return nil, err
			}
			cages[i].Dinosaurs = append(cages[i].Dinosaurs, d)
		}
	}
	return cages, nil
}

func AddCage(cage Cage) (Cage, error) {
	res, err := DB.Exec("INSERT INTO cages (power_status, maximum_capacity) VALUES (?, ?)", cage.PowerStatus, cage.MaximumCapacity)
	if err != nil {
		return Cage{}, err
	}
	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return Cage{}, err
	}
	return GetCage(int(id))
}

func UpdateCage(id int, cage Cage) (Cage, error) {
	if len(cage.Dinosaurs) > 0 && cage.PowerStatus == "DOWN" {
		return Cage{}, fmt.Errorf("unable to update cage power status to DOWN while it contains dinosaurs")
	}
	_, err := DB.Exec("UPDATE cages SET power_status = ?, maximum_capacity = ? WHERE id = ?", cage.PowerStatus, cage.MaximumCapacity, id)
	if err != nil {
		return Cage{}, err
	}
	return GetCage(id)
}

func DeleteCage(id int) (int64, error) {
	res, err := DB.Exec("DELETE FROM cages WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func AddDinosaurToCage(cageID int, dinosaurID int) error {
	cage, err := GetCage(cageID)
	if err != nil {
		return err
	}
	dinosaur, err := GetDinosaur(dinosaurID)
	if err != nil {
		return err
	}

	if dinosaur.CageID != nil {
		return fmt.Errorf("unable to place dinosaur id %d in cage %d, it is already in cage %d", dinosaurID, cageID, *dinosaur.CageID)
	}

	if cage.PowerStatus == "DOWN" {
		return fmt.Errorf("unable to place dinosaur id %d in cage %d, power status is %s", dinosaurID, cageID, cage.PowerStatus)
	}

	if len(cage.Dinosaurs) == cage.MaximumCapacity {
		return fmt.Errorf("unable to place dinosaur id %d in cage %d, maximum capacity of %d already met", dinosaurID, cageID, cage.MaximumCapacity)
	}

	for _, existingDinosaur := range cage.Dinosaurs {
		if existingDinosaur.FoodPreference != dinosaur.FoodPreference {
			return fmt.Errorf("unable to place dinosaur id %d in cage %d, %s (%s) is not compatible with %s (%s)", dinosaurID, cageID, dinosaur.Species, dinosaur.FoodPreference, existingDinosaur.Species, existingDinosaur.FoodPreference)
		}
		if dinosaur.FoodPreference == "carnivore" && dinosaur.Species != existingDinosaur.Species {
			return fmt.Errorf("unable to place dinosaur id %d in cage %d, %s (%s) is not compatible with %s (%s)", dinosaurID, cageID, dinosaur.Species, dinosaur.FoodPreference, existingDinosaur.Species, existingDinosaur.FoodPreference)
		}
	}
	_, err = DB.Exec("INSERT INTO dinosaurs_in_cages (cage_id, dinosaur_id) VALUES (?, ?)", cageID, dinosaurID)
	return err
}

func UpdateDinosaurCage(cageID int, dinosaurID int) error {
	cage, err := GetCage(cageID)
	if err != nil {
		return err
	}
	dinosaur, err := GetDinosaur(dinosaurID)
	if err != nil {
		return err
	}

	if cage.PowerStatus == "DOWN" {
		return fmt.Errorf("unable to place dinosaur id %d in cage %d, power status is %s", dinosaurID, cageID, cage.PowerStatus)
	}

	if len(cage.Dinosaurs) == cage.MaximumCapacity {
		return fmt.Errorf("unable to place dinosaur id %d in cage %d, maximum capacity of %d already met", dinosaurID, cageID, cage.MaximumCapacity)
	}

	for _, existingDinosaur := range cage.Dinosaurs {
		if existingDinosaur.FoodPreference != dinosaur.FoodPreference {
			return fmt.Errorf("unable to place dinosaur id %d in cage %d, %s (%s) is not compatible with %s (%s)", dinosaurID, cageID, dinosaur.Species, dinosaur.FoodPreference, existingDinosaur.Species, existingDinosaur.FoodPreference)
		}
		if dinosaur.FoodPreference == "carnivore" && dinosaur.Species != existingDinosaur.Species {
			return fmt.Errorf("unable to place dinosaur id %d in cage %d, %s (%s) is not compatible with %s (%s)", dinosaurID, cageID, dinosaur.Species, dinosaur.FoodPreference, existingDinosaur.Species, existingDinosaur.FoodPreference)
		}
	}
	_, err = DB.Exec("INSERT INTO dinosaurs_in_cages (cage_id, dinosaur_id) VALUES (?, ?)", cageID, dinosaurID)
	if err != nil {
		return err
	}
	return DeleteDinosaurFromCage(*dinosaur.CageID, dinosaurID)
}

func DeleteDinosaurFromCage(cageID int, dinosaurID int) error {
	res, err := DB.Exec("DELETE FROM dinosaurs_in_cages WHERE cage_id = ? AND dinosaur_id = ?", cageID, dinosaurID)
	if err != nil {
		return err
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return fmt.Errorf("invalid cage id %d or dinosaur id %d", cageID, dinosaurID)
	}
	return nil
}
