package types

type ResponseTemplate struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type DbInfo struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
	Sslmode  string
	Timezone string
}

type UserDeatils struct {
	ID            string `json:"user_id"`
	UserName      string `json:"user_name"`
	FullName      string `json:"full_name"`
	Email         string `json:"email"`
	ProfileImgUrl string `json:"profile_img_url"`
}

type UserGameDeatils struct {
	UGPID           string `json:"ugp_id"`
	GameID          string `json:"game_id"`
	GameName        string `json:"game_name"`
	CompanyName     string `json:"company_name"`
	GameDescription string `json:"game_description"`
	GameImgUrl      string `json:"game_img_url"`
}

type UserGameData struct {
	UserID          string            `json:"user_id"`
	UserGameDeatils []UserGameDeatils `json:"user_game_details"`
}
