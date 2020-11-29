package model

type (
	CreateUserRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	GetTokenRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	GetTokenResponse struct {
		Token string `json:"token"`
	}

	GetUserRequest struct{}

	GetUserResponse struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
)
