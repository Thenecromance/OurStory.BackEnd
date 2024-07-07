package fileWatchDog

type FileCallback struct {
	OnChanged func()
	OnDeleted func()
	OnCreated func()
	OnRenamed func()
}

type Subscriber interface {
	// WatchedFilePath target which file is being watched
	WatchedFilePath() string
	// WatchDogCallback need to provide a list callbacks when each event happened
	WatchDogCallback() FileCallback
}
