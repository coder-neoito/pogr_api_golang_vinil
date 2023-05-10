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
	resp.Success = true
	resp.Message = "User Fetched!"

	RespondWithJSON(c.Writer, http.StatusOK, resp)

}

func GetUserGamesByUserID(c *gin.Context) {

	// to fetch user game detials from db
	var games []types.UserGameDeatils
	var data types.UserGameData

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

	query := `select ugp.id as ugp_id, g.id as game_id, g.name as game_name, g.company_name, g.description as game_description, g.image_url as game_img_url from games g
	inner join user_game_profiles ugp on ugp.game_id = g.id
	where ugp.user_id=?`

	gormRes := db.Raw(query, id).Scan(&games)

	if gormRes.RowsAffected == 0 {
		RespondWithNoContent(c.Writer)
		return
	}

	data.UserID = id
	data.UserGameDeatils = games

	var resp types.ResponseTemplate
	resp.Data = data
	resp.Success = true
	resp.Message = "User Games Fetched!"

	RespondWithJSON(c.Writer, http.StatusOK, resp)

}
