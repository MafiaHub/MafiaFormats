package textdb

// FormatTextDB specifies textdb_xx.def format
type FormatTextDB struct {
	NumBlobs    uint32
	Unknown     uint32
	TextBlobs   []TextBlob `json:"-"`
	DataStrings uint8      `json:"-" spec:"ptr"` // Buffer containing all text entry strings
}

// TextBlob specifies the text entry
type TextBlob struct {
	TextID     uint32 `json:"id"`
	TextOffset uint32 `json:"-"` // Absolute position to a null-terminated string
}
