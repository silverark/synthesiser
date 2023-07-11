package notes

import (
	"math"
)

//var (
//	C0  = freq(-57)
//	Db0 = freq(-56)
//	D0  = freq(-55)
//	Eb0 = freq(-54)
//	E0  = freq(-53)
//	F0  = freq(-52)
//	Gb0 = freq(-51)
//	G0  = freq(-50)
//	Ab0 = freq(-49)
//	A0  = freq(-48)
//	Bb0 = freq(-47)
//	B0  = freq(-46)
//	C1  = freq(-45)
//	Db1 = freq(-44)
//	D1  = freq(-43)
//	Eb1 = freq(-42)
//	E1  = freq(-41)
//	F1  = freq(-40)
//	Gb1 = freq(-39)
//	G1  = freq(-38)
//	Ab1 = freq(-37)
//	A1  = freq(-36)
//	Bb1 = freq(-35)
//	B1  = freq(-34)
//	C2  = freq(-33)
//	Db2 = freq(-32)
//	D2  = freq(-31)
//	Eb2 = freq(-30)
//	E2  = freq(-29)
//	F2  = freq(-28)
//	Gb2 = freq(-27)
//	G2  = freq(-26)
//	Ab2 = freq(-25)
//	A2  = freq(-24)
//	Bb2 = freq(-23)
//	B2  = freq(-22)
//	C3  = freq(-21)
//	Db3 = freq(-20)
//	D3  = freq(-19)
//	Eb3 = freq(-18)
//	E3  = freq(-17)
//	F3  = freq(-16)
//	Gb3 = freq(-15)
//	G3  = freq(-14)
//	Ab3 = freq(-13)
//	A3  = freq(-12)
//	Bb3 = freq(-11)
//	B3  = freq(-10)
//	C4  = freq(-9)
//	Db4 = freq(-8)
//	D4  = freq(-7)
//	Eb4 = freq(-6)
//	E4  = freq(-5)
//	F4  = freq(-4)
//	Gb4 = freq(-3)
//	G4  = freq(-2)
//	Ab4 = freq(-1)
//	A4  = freq(0)
//	Bb4 = freq(1)
//	B4  = freq(2)
//	C5  = freq(3)
//	Db5 = freq(4)
//	D5  = freq(5)
//	Eb5 = freq(6)
//	E5  = freq(7)
//	F5  = freq(8)
//	Gb5 = freq(9)
//	G5  = freq(10)
//	Ab5 = freq(11)
//	A5  = freq(12)
//	Bb5 = freq(13)
//	B5  = freq(14)
//	C6  = freq(15)
//	Db6 = freq(16)
//	D6  = freq(17)
//	Eb6 = freq(18)
//	E6  = freq(19)
//	F6  = freq(20)
//	Gb6 = freq(21)
//	G6  = freq(22)
//	Ab6 = freq(23)
//	A6  = freq(24)
//	Bb6 = freq(25)
//	B6  = freq(26)
//	C7  = freq(27)
//	Db7 = freq(28)
//	D7  = freq(29)
//	Eb7 = freq(30)
//	E7  = freq(31)
//	F7  = freq(32)
//	Gb7 = freq(33)
//	G7  = freq(34)
//	Ab7 = freq(35)
//	A7  = freq(36)
//	Bb7 = freq(37)
//	B7  = freq(38)
//	C8  = freq(39)
//	Db8 = freq(40)
//	D8  = freq(41)
//	Eb8 = freq(42)
//	E8  = freq(43)
//	F8  = freq(44)
//	Gb8 = freq(45)
//	G8  = freq(46)
//	Ab8 = freq(47)
//	A8  = freq(48)
//	Bb8 = freq(49)
//	B8  = freq(50)
//	C9  = freq(51)
//)

type Note int

const (
	C0 = iota - 57
	Db0
	D0
	Eb0
	E0
	F0
	Gb0
	G0
	Ab0
	A0
	Bb0
	B0
	C1
	Db1
	D1
	Eb1
	E1
	F1
	Gb1
	G1
	Ab1
	A1
	Bb1
	B1
	C2
	Db2
	D2
	Eb2
	E2
	F2
	Gb2
	G2
	Ab2
	A2
	Bb2
	B2
	C3
	Db3
	D3
	Eb3
	E3
	F3
	Gb3
	G3
	Ab3
	A3
	Bb3
	B3
	C4
	Db4
	D4
	Eb4
	E4
	F4
	Gb4
	G4
	Ab4
	A4
	Bb4
	B4
	C5
	Db5
	D5
	Eb5
	E5
	F5
	Gb5
	G5
	Ab5
	A5
	Bb5
	B5
	C6
	Db6
	D6
	Eb6
	E6
	F6
	Gb6
	G6
	Ab6
	A6
	Bb6
	B6
	C7
	Db7
	D7
	Eb7
	E7
	F7
	Gb7
	G7
	Ab7
	A7
	Bb7
	B7
	C8
	Db8
	D8
	Eb8
	E8
	F8
	Gb8
	G8
	Ab8
	A8
	Bb8
	B8
	C9
)

func (n *Note) Add(step int) Note {
	return Note(int(*n) + step)
}

func Freq(step Note) float64 {
	return 440.0 * (math.Pow(2, (float64(step) / 12.0)))
}
