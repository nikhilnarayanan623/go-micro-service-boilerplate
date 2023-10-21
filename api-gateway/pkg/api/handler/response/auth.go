package response

type SignUp struct {
	UserID string `json:"user_id"`
}

type SignIn struct {
	AccessToken string `json:"access_token"`
}
