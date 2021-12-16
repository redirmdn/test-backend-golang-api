package web

type PostTestRequest struct {
	Name        string `json:"name" binding:"required,max=12"`
	Description string `json:"description" binding:"required"`
	Number      int    `json:"number"`
}
