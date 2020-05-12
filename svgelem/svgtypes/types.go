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

type Circle struct {
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

// Note that the specification says that there is a limit of 255 characters on
// this - no checking for this limit is performed here. The specification also
// notes that elements of the same time can be concatenated; logic to handle this
// case is not included here.
//
// It is probably sensible to return an error code, particularly for the case that
// the string ends up being too long.
func (p *Path) generatePathElementWithParameters(e PathElement, params int) (s string) {
	if len(e.Parameters) != params {
		return
	}
	s = PathElementTypeCode[e.Mode] + " "
	for i := 0; i < params; i++ {
		s = s + fmt.Sprintf("%v ", e.Parameters[i])
	}
	return
}

func (p *Path) AddElement(e PathElement) {
	p.Elements = append(p.Elements, e)
}

func (p *Path) ToString() (pathString string) {
	for _, e := range p.Elements {
		s := ""
		switch e.Mode {
		// this follows the logic of the specification - it could be optimized
		// for lines of code, but decided to leave it this way for readability
		// (a smart compiler prob does this optimization in any case...)
		case MoveToAbsolute, MoveToRelative, LineToAbsolute, LineToRelative:
			s = p.generatePathElementWithParameters(e, 2)
		case HorizontalLineAbsolute, HorizontalLineRelative, VerticalLineAbsolute, VerticalLineRelative:
			s = p.generatePathElementWithParameters(e, 1)
		case CubicCurveAbsolute, CubicCurveRelative:
			s = p.generatePathElementWithParameters(e, 6)
		case ShortCubicCurveAbsolute, ShortCubicCurveRelative:
			s = p.generatePathElementWithParameters(e, 4)
		case QuadraticCurveAbsolute, QuadraticCurveRelative:
			s = p.generatePathElementWithParameters(e, 4)
		case ShortQuadraticCurveAbsolute, ShortQuadraticCurveRelative:
			s = p.generatePathElementWithParameters(e, 2)
		case ArcAbsolute, ArcRelative:
			s = p.generatePathElementWithParameters(e, 7)
		case ClosePath:
			s = p.generatePathElementWithParameters(e, 0)
		}
		pathString += s + " "
	}
	return
}
