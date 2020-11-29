package domain

const ()

type (
	// User holds user data
	User struct {
		ID           int  // id is passed to domain model for simplicity
		Service      bool // if user is a service
		Username     string
		Password     string
		PasswordHash string
		Email        string // TODO: add ent validation
	}
)
