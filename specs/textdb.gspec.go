package textdb

// FormatTextDB specifies textdb_xx.def format
type FormatTextDB struct {
	NumBlobs  uint32
	Unknown   uint32
	TextBlobs []TextBlob
}

// TextBlob specifies the text entry
type TextBlob struct {
	TextID     uint32
	TextOffset uint32 // Absolute position to a null-terminated string
}
