package main

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"prizepicks-assessment/models"
	"testing"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	models.ConnectTestDatabase()
	os.Exit(m.Run())
}

func TestDinosaurGET(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/dinosaurs/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	jsonassert.New(t).Assertf(w.Body.String(), `{"data":{"id":1,"name":"Al","species":"Tyrannosaurus","food_preference":"carnivore"}}`)
}

func TestDinosaursGET(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/dinosaurs", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	jsonassert.New(t).Assertf(w.Body.String(), `{"data":["<<UNORDERED>>",{"id":1,"name":"Al","species":"Tyrannosaurus","food_preference":"carnivore"},{"id":2,"name":"Bob","species":"Velociraptor","food_preference":"carnivore"},{"id":3,"name":"Stewart","species":"Spinosaurus","food_preference":"carnivore"},{"id":4,"name":"Ralph","species":"Megalosaurus","food_preference":"carnivore"},{"id":5,"name":"Ernie","species":"Brachiosaurus","food_preference":"herbivore"},{"id":6,"name":"Harvey","species":"Stegosaurus","food_preference":"herbivore"},{"id":7,"name":"Mike","species":"Ankylosaurus","food_preference":"herbivore"},{"id":8,"name":"Harold","species":"Triceratops","food_preference":"herbivore"}]}`)
}

func TestDinosaursPOST(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/dinosaurs", bytes.NewReader([]byte(`{"name": "Willie", "species": "Pterodactylus", "food_preference": "carnivore"}`)))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	jsonassert.New(t).Assertf(w.Body.String(), `{"data":{"id":9,"name":"Willie","species":"Pterodactylus","food_preference":"carnivore"}}`)
}

func TestDinosaurPUT(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodPut, "/dinosaurs/9", bytes.NewReader([]byte(`{"name": "Willy"}`)))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	jsonassert.New(t).Assertf(w.Body.String(), `{"data":{"id":9,"name":"Willy","species":"Pterodactylus","food_preference":"carnivore"}}`)
}

func TestDinosaurDELETE(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodDelete, "/dinosaurs/9", bytes.NewReader([]byte(`{"name": "Willy"}`)))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	jsonassert.New(t).Assertf(w.Body.String(), `{"data":"dinosaur id 9 successfully removed."}`)
}

func TestDinosaurToCage(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/cages/1/dinosaur/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	jsonassert.New(t).Assertf(w.Body.String(), `{"cage_id":1,"dinosaur_id":1}`)
}
