package models

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCages(t *testing.T) {
	t.Run("GetCage", func(t *testing.T) {
		t.Run("OK", func(t *testing.T) {
			c := Cage{ID: 1, PowerStatus: "ACTIVE", MaximumCapacity: 1}
			cage, err := GetCage(1)
			require.NoError(t, err)
			require.Equal(t, c, cage)
		})

		t.Run("InvalidID", func(t *testing.T) {
			cage, err := GetCage(0)
			require.ErrorIs(t, err, sql.ErrNoRows)
			require.Equal(t, Cage{}, cage)
		})
	})

	t.Run("GetCages", func(t *testing.T) {
		c := []Cage{
			{ID: 1, PowerStatus: "ACTIVE", MaximumCapacity: 1, Dinosaurs: []Dinosaur{}},
			{ID: 2, PowerStatus: "ACTIVE", MaximumCapacity: 5, Dinosaurs: []Dinosaur{}},
			{ID: 3, PowerStatus: "DOWN", MaximumCapacity: 5, Dinosaurs: []Dinosaur{}},
		}
		cages, err := GetCages()
		require.NoError(t, err)
		require.Equal(t, c, cages)
	})
}
