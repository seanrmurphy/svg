package svgelem

import (
	"fmt"
	"log"
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

type TPathElement struct {
	Mode       PathElementType
	Parameters []float64
}

type TPath struct {
	Elements []TPathElement
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

func generateMoveToElement(e TPathElement) (s string) {
	if len(e.Parameters) != 2 {
		return
	}
	x := fmt.Sprintf("%v", e.Parameters[0])
	y := fmt.Sprintf("%v", e.Parameters[1])
	s = PathElementTypeCode[e.Mode] + " " + x + " " + y
	return
}

func generateLineToElement(e TPathElement) (s string) {
	if len(e.Parameters) != 2 {
		return
	}
	x := fmt.Sprintf("%v", e.Parameters[0])
	y := fmt.Sprintf("%v", e.Parameters[1])
	s = PathElementTypeCode[e.Mode] + " " + x + " " + y
	return
}

func generateHorizontalLineElement(e TPathElement) (s string) {
	if len(e.Parameters) != 1 {
		return
	}
	p := fmt.Sprintf("%v", e.Parameters[0])
	s = PathElementTypeCode[e.Mode] + " " + p
	return
}

func generateVerticalLineElement(e TPathElement) (s string) {
	if len(e.Parameters) != 1 {
		return
	}
	p := fmt.Sprintf("%v", e.Parameters[0])
	s = PathElementTypeCode[e.Mode] + " " + p
	return
}

func generateCubicCurveElement(e TPathElement) (s string) {
	s = PathElementTypeCode[e.Mode]
	log.Printf("Not implemented\n")
	return
}

func generateShortCubicCurveElement(e TPathElement) (s string) {
	s = PathElementTypeCode[e.Mode]
	log.Printf("Not implemented\n")
	return
}

func generateQuadraticCurveElement(e TPathElement) (s string) {
	s = PathElementTypeCode[e.Mode]
	log.Printf("Not implemented\n")
	return
}

func generateShortQuadraticCurveElement(e TPathElement) (s string) {
	s = PathElementTypeCode[e.Mode]
	log.Printf("Not implemented\n")
	return
}

func generateClosePath(e TPathElement) (s string) {
	s = PathElementTypeCode[e.Mode] + " "
	return
}

func generateArcElement(e TPathElement) (s string) {
	if len(e.Parameters) != 7 {
		return
	}
	rx := fmt.Sprintf("%v", e.Parameters[0])
	ry := fmt.Sprintf("%v", e.Parameters[1])
	xRotation := fmt.Sprintf("%v", e.Parameters[2])
	largeArcFlag := fmt.Sprintf("%v", e.Parameters[3])
	sweepFlag := fmt.Sprintf("%v", e.Parameters[4])
	x := fmt.Sprintf("%v", e.Parameters[5])
	y := fmt.Sprintf("%v", e.Parameters[6])
	s = PathElementTypeCode[e.Mode] + " " + rx + " " + ry + " " + xRotation + " " + largeArcFlag + " " + sweepFlag + " " + x + " " + y
	return
}

func generatePathWithParameters(e TPathElement, params int) (s string) {
	if len(e.Parameters) != params {
		return
	}
	s = PathElementTypeCode[e.Mode] + " "
	for i := 0; i < params; i++ {
		s = s + fmt.Sprintf("%v ", e.Parameters[i])
	}
	//rx := fmt.Sprintf("%v", e.Parameters[0])
	//ry := fmt.Sprintf("%v", e.Parameters[1])
	//xRotation := fmt.Sprintf("%v", e.Parameters[2])
	//largeArcFlag := fmt.Sprintf("%v", e.Parameters[3])
	//sweepFlag := fmt.Sprintf("%v", e.Parameters[4])
	//x := fmt.Sprintf("%v", e.Parameters[5])
	//y := fmt.Sprintf("%v", e.Parameters[6])
	//s = PathElementTypeCode[e.Mode] + " " + rx + " " + ry + " " + xRotation + " " + largeArcFlag + " " + sweepFlag + " " + x + " " + y
	return
}

func GeneratePathString(p TPath) (pathString string) {
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
