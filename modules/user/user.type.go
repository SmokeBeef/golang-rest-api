
package user

type UserLoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}