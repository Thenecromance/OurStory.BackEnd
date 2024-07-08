package mq

/*
all topics' name should be defined here
*/

const (
	DevRoute = "dev.route"
)

const (
	FileOp = "file.operation" // when other component need to watch file operation, just subscribe this topic
)
