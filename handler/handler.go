package handler

import (
	"net/http"

	"GoWorkSpace/goCLinic/domain"
	"GoWorkSpace/goCLinic/utils"

	"github.com/gin-gonic/gin"
)

var clinics = make(map[string]domain.Clinic)

func CreateClinicHandler(c *gin.Context) {
	var clinic domain.Clinic
	if err := c.ShouldBindJSON(&clinic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clinic.ID = utils.GenerateID()
	clinics[clinic.ID] = clinic

	c.JSON(http.StatusCreated, clinic)
}
