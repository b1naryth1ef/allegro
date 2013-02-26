package allegro

// #include <allegro5/allegro.h>
import "C"

const (
	FixToRad = int32(1608)
	RadToFix = int32(2670177)
)

func IntToFix(x int32) int32 {
	return int32(C.al_itofix(C.int(x)))
}

func FixToInt(x int32) int32 {
	return int32(C.al_fixtoi(C.al_fixed(x)))
}

func FixFloor(x int32) int32 {
	return int32(C.al_fixfloor(C.al_fixed(x)))
}

func FixCeil(x int32) int32 {
	return int32(C.al_fixceil(C.al_fixed(x)))
}

func FloatToFix(x float64) int32 {
	return int32(C.al_ftofix(C.double(x)))
}

func FixToFloat(x int32) float64 {
	return float64(C.al_fixtof(C.al_fixed(x)))
}

func FixMul(x, y int32) int32 {
	return int32(C.al_fixmul(C.al_fixed(x), C.al_fixed(y)))
}

func FixDiv(x, y int32) int32 {
	return int32(C.al_fixdiv(C.al_fixed(x), C.al_fixed(y)))
}

func FixAdd(x, y int32) int32 {
	return int32(C.al_fixadd(C.al_fixed(x), C.al_fixed(y)))
}

func FixSub(x, y int32) int32 {
	return int32(C.al_fixsub(C.al_fixed(x), C.al_fixed(y)))
}

/****************************/
/* Fixed point trigonometry */
/****************************/

func FixSin(x int32) int32 {
	return int32(C.al_fixsin(C.al_fixed(x)))
}

func FixCos(x int32) int32 {
	return int32(C.al_fixcos(C.al_fixed(x)))
}

func FixTan(x int32) int32 {
	return int32(C.al_fixtan(C.al_fixed(x)))
}

func FixAsin(x int32) int32 {
	return int32(C.al_fixasin(C.al_fixed(x)))
}

func FixAcos(x int32) int32 {
	return int32(C.al_fixacos(C.al_fixed(x)))
}

func FixAtan(x int32) int32 {
	return int32(C.al_fixatan(C.al_fixed(x)))
}

func FixAtan2(x, y int32) int32 {
	return int32(C.al_fixatan2(C.al_fixed(x), C.al_fixed(y)))
}

func FixSqrt(x int32) int32 {
	return int32(C.al_fixsqrt(C.al_fixed(x)))
}

func FixHypot(x, y int32) int32 {
	return int32(C.al_fixhypot(C.al_fixed(x), C.al_fixed(y)))
}
