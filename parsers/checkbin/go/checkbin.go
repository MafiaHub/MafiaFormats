/* This file has been generated by go-specgen */
package main

/* PointType */
const (
	Pedestrian = 0x1 
	AI = 0x2 
	Vehicle = 0x4 
	Tram_Station = 0x8 
	Special = 0x10 
)

/* LinkType */
const (
	Pedestrian = 1 
	AI = 2 
	Trains_Forward = 4 
	Trains_Reverse = 0x8400 
	Other = 0x1000 
)


type FormatCheckBIN struct {
	Magic uint32  /* Should be 0x1ABCEDF */
	NumPoints uint32  /* Number of navigation nodes */
	Points []Point 
	Links []Link  /* Size depends on number of links used by all points */
}

type Point struct {
	Type PointType 
	ID uint16 
	AreaSize uint16 
	Unk [10]uint8  `spec:"plain"`  /* Unknown values */
	EnterLinks uint8  /* How many consequent links belong to this node */
	ExitLinks uint8  /* Is the same as EnterLinks */
}

type Link struct {
	TargetPoint uint16  /* Refers back to a point it connects to */
	LinkType LinkType 
	Unk float32 
}

