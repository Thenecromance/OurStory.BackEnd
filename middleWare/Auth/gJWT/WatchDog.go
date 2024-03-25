package gJWT

type watchDog struct {
	watchedToken map[string]bool
}

func (w *watchDog) watchToken(token string) {
	w.watchedToken[token] = true

}

func (w *watchDog) isValidToken(token string) bool {
	return w.watchedToken[token]
}

func (w *watchDog) expiredToken(token string) {
	delete(w.watchedToken, token)
}

func newWatchDog() *watchDog {
	return &watchDog{
		watchedToken: make(map[string]bool),
	}
}
