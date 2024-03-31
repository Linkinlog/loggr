package models

//go:generate stringer -type=ItemType
type ItemType int

const (
	_ ItemType = iota
	Plant
	Tool
	Seed
)

var ItemTypes = []ItemType{Plant, Tool, Seed}
