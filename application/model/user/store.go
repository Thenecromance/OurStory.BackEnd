package user

import (
	"github.com/Thenecromance/OurStories/utility/log"
	"time"
)

type store struct {
	cache map[string]*time.Timer
}

func (s *store) storeToken(token string) {
	if _, ok := s.cache[token]; ok {
		log.Errorf("token already exists")
		return
	}

	timer := time.AfterFunc(15*time.Minute, func() {
		s.markTokenExpired(token)
	})
	s.cache[token] = timer
}

func (s *store) markTokenExpired(token string) {
	if _, ok := s.cache[token]; ok {
		delete(s.cache, token)
	}
}

func (s *store) hasToken(token string) bool {
	if _, ok := s.cache[token]; ok {
		return true
	}
	return false
}
