package mnu

// FormatMNU specifies .mnu format
type FormatMNU struct {
	Magic       [4]uint8 // should be 'Menu'
	Unknown     uint32
	NumControls uint32
	Controls    []Control
}

// Control specifies menu entry
type Control struct {
	Unknown         uint32
	Type            [4]uint8   // @string
	Position        [2]float32 // @plain [X, Y]
	Scale           [2]float32 // @plain [W, H]
	TextID          uint32     // taken from Textdb_xx.def
	TextColor       uint16
	BackgroundColor uint16
}
