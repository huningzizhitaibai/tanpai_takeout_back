package commonResult

type LoginResult struct {
	Username string `json:"username"`
	Type     int    `json:"type"`
	Token    string `json:"token"`
}
