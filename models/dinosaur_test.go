package models

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDinosaurs(t *testing.T) {
	t.Run("GetDinosaur", func(t *testing.T) {
		t.Run("OK", func(t *testing.T) {
			d := Dinosaur{ID: 1, Name: "Al", Species: "Tyrannosaurus", FoodPreference: "carnivore"}
			dinosaur, err := GetDinosaur(1)
			require.NoError(t, err)
			require.Equal(t, d, dinosaur)
		})

		t.Run("InvalidID", func(t *testing.T) {
			dinosaur, err := GetDinosaur(0)
			require.ErrorIs(t, err, sql.ErrNoRows)
			require.Equal(t, Dinosaur{}, dinosaur)
		})
	})

	t.Run("GetDinosaurs", func(t *testing.T) {
		d := []Dinosaur{
			{ID: 1, Name: "Al", Species: "Tyrannosaurus", FoodPreference: "carnivore"},
			{ID: 2, Name: "Bob", Species: "Velociraptor", FoodPreference: "carnivore"},
			{ID: 3, Name: "Stewart", Species: "Spinosaurus", FoodPreference: "carnivore"},
			{ID: 4, Name: "Ralph", Species: "Megalosaurus", FoodPreference: "carnivore"},

			{ID: 5, Name: "Ernie", Species: "Brachiosaurus", FoodPreference: "herbivore"},
			{ID: 6, Name: "Harvey", Species: "Stegosaurus", FoodPreference: "herbivore"},
			{ID: 7, Name: "Mike", Species: "Ankylosaurus", FoodPreference: "herbivore"},
			{ID: 8, Name: "Harold", Species: "Triceratops", FoodPreference: "herbivore"},
		}
		dinosaurs, err := GetDinosaurs()
		require.NoError(t, err)
		require.Equal(t, d, dinosaurs)
	})

	t.Run("AddDinosaur", func(t *testing.T) {
		d := Dinosaur{Name: "Willie", Species: "Pterodactylus", FoodPreference: "carnivore"}
		dinosaur, err := AddDinosaur(d)
		d.ID = dinosaur.ID
		require.NoError(t, err)
		require.Equal(t, d, dinosaur)
	})

	t.Run("UpdateDinosaur", func(t *testing.T) {
		d, err := GetDinosaur(9)
		require.NoError(t, err)
		d.Name = "Willie"
		dinosaur, err := UpdateDinosaur(d)
		require.NoError(t, err)
		require.Equal(t, d, dinosaur)
	})

	t.Run("DeleteDinosaur", func(t *testing.T) {
		t.Run("OK", func(t *testing.T) {
			err := DeleteDinosaur(9)
			require.NoError(t, err)
		})

		t.Run("InvalidID", func(t *testing.T) {
			err := DeleteDinosaur(0)
			require.ErrorIs(t, err, sql.ErrNoRows)
		})
	})
}
