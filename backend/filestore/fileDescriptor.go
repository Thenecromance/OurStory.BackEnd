package filestore

type discriptor struct {
	FileId      int    `json:"fileId"      db:"fileId"`
	Owner       int    `json:"owner"       db:"owner"`
	Size        int    `json:"size"        db:"size"`        // record the size of the file
	FileName    string `json:"filename"    db:"filename"`    // record the name of the file
	FilePath    string `json:"filepath"    db:"filepath"`    // record the path of the file
	FingerPrint string `json:"fingerprint" db:"fingerprint"` // record the fingerprint of the file
}
