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
-> TrainsAndSalinas_Forward = 4;
-> TrainsAndSalinas_Reverse = 0x8400;
-> Other = 0x1000;
*/

// Header specifies the file's magic constant and number of navigation points
type Header struct {
	Magic     uint32
	NumPoints uint32
}

// Point specifies an actual nav node with enter/leave links and properties
type Point struct {
	Type       uint16
	ID         uint16
	AreaSize   uint16
	Unk        [10]uint8 /* @plain Unknown values */
	EnterLinks uint8
	ExitLinks  uint8 // Is the same as EnterLinks
}

// Link specifies the connection between points
type Link struct {
	TargetPoint uint16
	LinkType    uint16
	Unk         float32
}
