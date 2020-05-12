package svgtypes

import (
	"fmt"
)

// This file contains some helper types and ways to use them instead
// of using attrs all the time...

type PathElementType int

const (
	MoveToAbsolute PathElementType = iota
	MoveToRelative
	LineToAbsolute
	LineToRelative
	HorizontalLineAbsolute
	HorizontalLineRelative
	VerticalLineAbsolute
	VerticalLineRelative
	CubicCurveAbsolute
	CubicCurveRelative
	ShortCubicCurveAbsolute
	ShortCubicCurveRelative
	QuadraticCurveAbsolute
	QuadraticCurveRelative
	ShortQuadraticCurveAbsolute
	ShortQuadraticCurveRelative
	ArcAbsolute
	ArcRelative
	ClosePath
)

var PathElementTypeCode map[PathElementType]string

type PathElement struct {
	Mode       PathElementType
	Parameters []float64
}

type Path struct {
	Elements []PathElement
}

type TCircle struct {
	Radius  float64
	CentreX float64
	CentreY float64
}

func init() {
	PathElementTypeCode = map[PathElementType]string{
		MoveToAbsolute:              "M",
		MoveToRelative:              "m",
		LineToAbsolute:              "L",
		LineToRelative:              "l",
		HorizontalLineAbsolute:      "H",
		HorizontalLineRelative:      "h",
		VerticalLineAbsolute:        "V",
		VerticalLineRelative:        "v",
		CubicCurveAbsolute:          "C",
		CubicCurveRelative:          "c",
		ShortCubicCurveAbsolute:     "S",
		ShortCubicCurveRelative:     "s",
		QuadraticCurveAbsolute:      "Q",
		QuadraticCurveRelative:      "q",
		ShortQuadraticCurveAbsolute: "T",
		ShortQuadraticCurveRelative: "t",
		ArcAbsolute:                 "A",
		ArcRelative:                 "a",
		ClosePath:                   "z",
	}
}

func generatePathWithParameters(e PathElement, params int) (s string) {
	if len(e.Parameters) != params {
		return
	}
	s = PathElementTypeCode[e.Mode] + " "
	for i := 0; i < params; i++ {
		s = s + fmt.Sprintf("%v ", e.Parameters[i])
	}
	return
}

func GeneratePathString(p Path) (pathString string) {
	for _, e := range p.Elements {
		s := ""
		switch e.Mode {
		case MoveToAbsolute, MoveToRelative, LineToAbsolute, LineToRelative:
			s = generatePathWithParameters(e, 2)
		case HorizontalLineAbsolute, HorizontalLineRelative, VerticalLineAbsolute, VerticalLineRelative:
			s = generatePathWithParameters(e, 1)
		case CubicCurveAbsolute, CubicCurveRelative:
			s = generatePathWithParameters(e, 6)
		case ShortCubicCurveAbsolute, ShortCubicCurveRelative:
			s = generatePathWithParameters(e, 4)
		case QuadraticCurveAbsolute, QuadraticCurveRelative:
			s = generatePathWithParameters(e, 4)
		case ShortQuadraticCurveAbsolute, ShortQuadraticCurveRelative:
			s = generatePathWithParameters(e, 2)
		case ArcAbsolute, ArcRelative:
			s = generatePathWithParameters(e, 7)
		case ClosePath:
			s = generatePathWithParameters(e, 0)
		}
		pathString += s + " "
	}
	return
}
