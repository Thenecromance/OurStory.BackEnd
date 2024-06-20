package models

// Object is a struct that holds the data for the local file storage
type Object struct {
	// where the file is stored
	path string
	// the name of the file
	name string
	// the file's extension
	extension string
	// the size of the file
	size int64
	// the type of the file
	fileType string
	// the file's content
	content []byte
	// the file's hash
	hash string
}
