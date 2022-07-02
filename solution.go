package square

import "math"

const SidesSquare = 4
const SidesTriangle = 3
const SidesCircle = 0

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum int64) float64 {
	if sidesNum == int64(SidesSquare) {
		return sideLen * sideLen
	} else if sidesNum == int64(SidesTriangle) {
		return (float64(math.Sqrt(3) / 4)) * sideLen * sideLen
	} else if sidesNum == int64(SidesCircle) {
		return (float64(math.Pi)) * sideLen * sideLen
	} else {
		return 0.0
	}
}
