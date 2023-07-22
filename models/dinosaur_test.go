package models

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	ConnectTestDatabase()
	os.Exit(m.Run())
}

func TestDinosaurSingleGET(t *testing.T) {
	d := Dinosaur{ID: 1, Name: "Al", Species: "Tyrannosaurus", FoodPreference: "carnivore"}
	dinosaur, err := GetDinosaur(1)
	require.NoError(t, err)
	require.EqualValues(t, d, dinosaur)
}

func TestDinosaurMultiGET(t *testing.T) {
	dinosaurs, err := GetDinosaurs()
	require.NoError(t, err)
	require.NotNil(t, dinosaurs)
	require.Len(t, dinosaurs, 8)
}
