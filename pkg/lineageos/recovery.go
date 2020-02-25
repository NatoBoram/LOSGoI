package lineageos

// Recovery is a device's recovery.
type Recovery struct {
	Filename string `json:"filename"`
	Filepath string `json:"filepath"`
	Sha1     string `json:"sha1"`    
	Sha256   string `json:"sha256"`  
	Size     int64  `json:"size"`    
}