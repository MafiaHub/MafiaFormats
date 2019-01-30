package checkbin

/*
@enum PointType::
-> Pedestrian = 0x1;
-> AI = 0x2;
-> Vehicle = 0x4;
-> Tram_Station = 0x8;
-> Special = 0x10;
*/

/*
@enum LinkType::
-> Pedestrian = 1;
-> AI = 2;
-> Trains_Forward = 4;
-> Trains_Reverse = 0x8400;
-> Other = 0x1000;
*/

// FormatCheckBIN specifies the check.bin navigation format
type FormatCheckBIN struct {
	Magic     uint32 // Should be 0x1ABCEDF
	NumPoints uint32 // Number of navigation nodes
	Points    []Point
	Links     []Link // Size depends on number of links used by all points
}

// Point specifies an actual nav node with enter/leave links and properties
type Point struct {
	Type       PointType
	ID         uint16
	AreaSize   uint16
	Unk        [10]uint8 /* @plain Unknown values */
	EnterLinks uint8     // How many consequent links belong to this node
	ExitLinks  uint8     // Is the same as EnterLinks
}

// Link specifies the connectio
type Link struct {
	TargetPoint uint16 // Refers back to a point it connects to
	LinkType    LinkType
	Unk         float32
}
