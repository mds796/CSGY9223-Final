package user

type CreateUserRequest struct {
	Alias string // The user name (alias)
}

type CreateUserResponse struct {
	UserId string // the HEX-encoded string representation of the UUID.
}

type Service interface {
	Create(CreateUserRequest request) (CreateUserResponse, error)
}
