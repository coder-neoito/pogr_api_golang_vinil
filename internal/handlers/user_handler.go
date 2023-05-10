package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinilunni/pogr_api_golang_vinil/internal/common/types"
	"github.com/vinilunni/pogr_api_golang_vinil/models"
)

func GetUserByID(c *gin.Context) {

	// to fetch user detials from db
	var data types.UserDeatils

	id := c.Param("id")

	if len(id) == 0 {
		RespondWithError(c.Writer, http.StatusBadRequest, "Invalid Request Payload")
		return
	}

	db := DBInit()

	//closing connection to db
	sqlDB, dberr := db.DB()
	if dberr != nil {
		log.Fatalln(dberr)
	}
	defer sqlDB.Close()

	gormRes := db.Model(&models.User{}).Where("id = ?", id).Find(&data)

	if gormRes.RowsAffected == 0 {
		RespondWithNoContent(c.Writer)
		return
	}

	var resp types.ResponseTemplate
	resp.Data = data
	resp.Success = false
	resp.Message = "User Fetched!"

	RespondWithJSON(c.Writer, http.StatusOK, resp)

}
