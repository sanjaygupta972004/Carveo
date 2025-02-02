package models

// @Description Represents the details of a User
type UserSwagger struct {
	UserID           string `json:"userID" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"`
	Email            string `json:"email" example:"example@gmail.com"`
	FullName         string `json:"fullName" example:"John Doe"`
	Password         string `json:"password" example:"Sdvdfgdf@234"`
	ProfileImage     string `json:"profileImage" example:"https://example.com/image.jpg"`
	IsAdmin          bool   `json:"isAdmin" example:"false"`
	IsEmailVerified  bool   `json:"isEmailVerified" example:"false"`
	IsMobileVerified bool   `json:"isMobileVerified" example:"false"`
	CreatedAt        string `json:"createdAt" example:"2021-07-01T00:00:00Z"`
	UpdatedAt        string `json:"updatedAt" example:"2021-07-01T00:00:00Z"`
}

type ErrorResponseUserSwagger struct {
	Message string `json:"message" example:"Error message "`
	Status  string `json:"status" example:"error or failed to validation request data or server error"`
	Code    int    `json:"code" example:"200"`
	Error   string `json:"error" example:"Error details"`
}

type SuccessResponseUserSwagger struct {
	Message string `json:"message" example:"Success message "`
	Status  string `json:"status" example:"success"`
	Code    int    `json:"code" example:"200"`
	Data    struct {
		UserID           string `json:"userID" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"`
		Email            string `json:"email" example:"example@gmail.com"`
		FullName         string `json:"fullName" example:"John Doe"`
		ProfileImage     string `json:"profileImage" example:"https://example.com/image.jpg"`
		IsAdmin          bool   `json:"isAdmin" example:"false"`
		IsEmailVerified  bool   `json:"isEmailVerified" example:"false"`
		IsMobileVerified bool   `json:"isMobileVerified" example:"false"`
		CreatedAt        string `json:"createdAt" example:"2021-07-01T00:00:00Z"`
		UpdatedAt        string `json:"updatedAt" example:"2021-07-01T00:00:00Z"`
	} `json:"data"`
}
