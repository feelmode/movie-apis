package movie

type ReqResp struct {
	ID          uint8  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Rating      uint8  `json:"rating"`
	Image       string `json:"image"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
