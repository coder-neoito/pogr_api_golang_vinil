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
