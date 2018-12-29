package defs

// requests
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

// data model

type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCting string
}

type SimpleSession struct {
	Username string // login name
	TTL      int64  // 过期时间
}
