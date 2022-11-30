package main

type Ordinal int64

const (
	North Ordinal = 0
	East Ordinal = 1
	South Ordinal = 3
	West Ordinal = 4
)

func (ordinal Ordinal) String() string {
	switch ordinal {
	case North:
		return "north"
	case East:
		return "east"
	case South:
		return "south"
	case West:
		return "west"
	}
	return ""
}
