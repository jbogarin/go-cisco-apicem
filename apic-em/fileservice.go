package apicem

// Filestring is ...
type Filestring struct {
	NameSpace     string `json:"nameSpace,omitempty"`    // A group of file IDs contained in a common nameSpace
	ID            string `json:"id,omitempty"`           // file indentification number
	Encrypted     bool   `json:"encrypted,omitempty"`    // isEncrypted of the file
	DownloadPath  string `json:"downloadPath,omitempty"` // Absolute path of the file
	FileFormat    string `json:"fileFormat,omitempty"`   // MIME Type of the File. e.g. text/plain, application/xml, audio/mpeg
	FileSize      string `json:"fileSize,omitempty"`     // Size of the file in bytes
	Md5Checksum   string `json:"md5Checksum,omitempty"`  // md5Checksum of the file
	Sha1Checksum  string `json:"sha1Checksum,omitempty"` // sha1Checksum of the file
	Name          string `json:"name,omitempty"`         // Name of the file
	AttributeInfo string `json:"attributeInfo,omitempty"`
}

// FileObjectListResult is ...
type FileObjectListResult struct {
	Version  string       `json:"version,omitempty"`
	Response []Filestring `json:"response,omitempty"`
}

// FileObjectResult is ...
type FileObjectResult struct {
	Version  string     `json:"version,omitempty"`
	Response Filestring `json:"response,omitempty"`
}

// NameSpaceListResult is ...
type NameSpaceListResult struct {
	Version  string   `json:"version,omitempty"`
	Response []string `json:"response,omitempty"`
}
