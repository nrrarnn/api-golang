package models

import "sync"

var TokenBlacklist = make(map[string]struct{})
var mu sync.Mutex 

func BlacklistToken(token string) {
	mu.Lock()
	defer mu.Unlock()
	TokenBlacklist[token] = struct{}{}
}
