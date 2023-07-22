package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"prizepicks-assessment/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()
	r := setupRouter()
	r.Run("127.0.0.1:8080")
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/cages", getCages)
	r.GET("/cages/:id", getCage)
	r.GET("/dinosaurs", getDinosaurs)
	r.GET("/dinosaurs/:id", getDinosaur)
	r.POST("/cages", addCage)
	r.POST("/cages/:id/dinosaur/:dinosaur_id", addDinosaurToCage)
	r.POST("/dinosaurs", addDinosaur)
	r.PUT("/cages/:id", updateCage)
	r.PUT("/cages/:id/dinosaur/:dinosaur_id/cage/:cage_id", updateDinosaurCage)
	r.PUT("/dinosaurs/:id", updateDinosaur)
	r.DELETE("/cages/:id", deleteCage)
	r.DELETE("/cages/:id/dinosaur/:dinosaur_id", deleteDinosaurFromCage)
	r.DELETE("/dinosaurs/:id", deleteDinosaur)
	return r
}

func getCages(c *gin.Context) {
	cages, err := models.GetCages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cages})
}

func getCage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	cage, err := models.GetCage(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("cage id %d was not found.", id)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cage})
}

func getDinosaurs(c *gin.Context) {
	dinosaurs, err := models.GetDinosaurs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": dinosaurs})
}

func getDinosaur(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	dinosaur, err := models.GetDinosaur(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("dinosaur id %d was not found.", id)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": dinosaur})
}

func addCage(c *gin.Context) {
	var cage models.Cage
	if err := c.ShouldBindJSON(&cage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cage, err := models.AddCage(cage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": cage})
}

func addDinosaur(c *gin.Context) {
	var dinosaur models.Dinosaur
	if err := c.ShouldBindJSON(&dinosaur); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dinosaur, err := models.AddDinosaur(dinosaur)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": dinosaur})
}

func updateCage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	cage, err := models.GetCage(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("cage id %d was not found.", id)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&cage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cage, err = models.UpdateCage(id, cage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cage})
}

func updateDinosaur(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	dinosaur, err := models.GetDinosaur(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("dinosaur id %d was not found.", id)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&dinosaur); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dinosaur, err = models.UpdateDinosaur(dinosaur)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": dinosaur})
}

func deleteCage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rowsAffected, err := models.DeleteCage(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("cage id %d was not found.", id)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("cage id %d successfully removed.", id)})
}

func deleteDinosaur(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err = models.DeleteDinosaur(id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("dinosaur id %d was not found.", id)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("dinosaur id %d successfully removed.", id)})
}

func addDinosaurToCage(c *gin.Context) {
	cageID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	dinosaurID, err := strconv.Atoi(c.Param("dinosaur_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = models.AddDinosaurToCage(cageID, dinosaurID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cage_id": cageID, "dinosaur_id": dinosaurID})
}

func updateDinosaurCage(c *gin.Context) {
	cageID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	dinosaurID, err := strconv.Atoi(c.Param("dinosaur_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = models.UpdateDinosaurCage(cageID, dinosaurID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cage_id": cageID, "dinosaur_id": dinosaurID})
}

func deleteDinosaurFromCage(c *gin.Context) {
	cageID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	dinosaurID, err := strconv.Atoi(c.Param("dinosaur_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = models.DeleteDinosaurFromCage(cageID, dinosaurID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cage_id": cageID, "dinosaur_id": dinosaurID})
}
