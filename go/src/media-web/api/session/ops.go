package session

import (
	"media-web/defs"
	"sync"
	"time"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

// 从数据库加载到缓存
func loadSessionsFromDB() {

}

// 生成SESSION id
func GenerateNewSessionId(un string) string {

}

// 是否已经过期
func IsSessionExpired(sid string) (string bool) {

}
