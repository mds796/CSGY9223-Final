package follow

type FollowRequest struct {
	FollowerUserID string
	FollowedUserID string
}

type FollowResponse struct {
}

type ViewRequest struct {
	UserID string
}

type ViewResponse struct {
	UserIDs []string
}

type Service interface {
	Follow(request FollowRequest) (FollowResponse, error)
	View(request ViewRequest) (ViewResponse, error)
}
