package user

type CreateUserRequest struct {
	Username string
}

type CreateUserResponse struct {
	Uuid string
}

type ViewUserRequest struct {
	Uuid string
}

type ViewUserResponse struct {
	Username string
}

type SearchUserRequest struct {
	Query string
}

type SearchUserResponse struct {
	Uuids []string
}

type Service interface {
	Create(request CreateUserRequest) (CreateUserResponse, error)
	View(request ViewUserRequest) (ViewUserResponse, error)
	Search(request SearchUserRequest) (SearchUserResponse, error)
}
