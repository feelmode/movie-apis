package movie

type ReqResp struct {
	ID          uint8  `json:"id"`
	Title       string `json:"title" valid:"required"`
	Description string `json:"description" valid:"required"`
	Rating      uint8  `json:"rating"`
	Image       string `json:"image"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
