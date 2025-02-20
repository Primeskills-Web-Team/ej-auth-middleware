package authmiddleware

// AuthResponse merepresentasikan response dari auth/validate
type AuthResponse struct {
	Code      string    `json:"code"`
	IsSuccess bool      `json:"is_success"`
	Message   string    `json:"message"`
	Data      *AuthData `json:"data"`
}

// AuthData merepresentasikan data user jika token valid
type AuthData struct {
	ID            int    `json:"id"`
	FullName      string `json:"full_name"`
	Email         string `json:"email"`
	AccountActive string `json:"account_active"`
}
