package model

// ShortLink 定义短链接的数据模型（可选）
type ShortLink struct {
	ShortCode string `json:"short_code"`
	LongURL   string `json:"long_url"`
}
