package user

type CreateUserRequest struct {
	Username string
}

type CreateUserResponse struct {
}

type ViewUserRequest struct {
	Username string
}

type ViewUserResponse struct {
	Uuid string
}

type SearchUserRequest struct {
	Query string
}

type SearchUserResponse struct {
	Usernames []string
}

type Service interface {
	Create(request CreateUserRequest) (CreateUserResponse, error)
	View(request ViewUserRequest) (ViewUserResponse, error)
	Search(request SearchUserRequest) (SearchUserResponse, error)
}
