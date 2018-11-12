package user

type CreateUserRequest struct {
	Username string
}

type CreateUserResponse struct {
	Uuid string
}

type ViewUserRequest struct {
	Username string
	UserID   string
}

type ViewUserResponse struct {
	Username string
	Uuid     string
}

type SearchUserRequest struct {
	Query string
}

type SearchUserResponse struct {
	Usernames []string
	UserIDs   []string
}

type Service interface {
	Create(request CreateUserRequest) (CreateUserResponse, error)
	View(request ViewUserRequest) (ViewUserResponse, error)
	Search(request SearchUserRequest) (SearchUserResponse, error)
}
