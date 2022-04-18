package Response

type GeneralResponse struct {
	Error   bool        `json:"error" binding:"required"`
	Message string      `json:"message" binding:"required"`
	Data    interface{} `json:"data" binding:"required"`
}
