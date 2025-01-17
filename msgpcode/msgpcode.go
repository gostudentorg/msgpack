package msgpcode

var (
	PosFixedNumHigh byte = 0x7f
	NegFixedNumLow  byte = 0xe0

	Nil byte = 0xc0

	False byte = 0xc2
	True  byte = 0xc3

	Float  byte = 0xca
	Double byte = 0xcb

	Uint8  byte = 0xcc
	Uint16 byte = 0xcd
	Uint32 byte = 0xce
	Uint64 byte = 0xcf

	Int8  byte = 0xd0
	Int16 byte = 0xd1
	Int32 byte = 0xd2
	Int64 byte = 0xd3

	FixedStrLow  byte = 0xa0
	FixedStrHigh byte = 0xbf
	FixedStrMask byte = 0x1f
	Str8         byte = 0xd9
	Str16        byte = 0xda
	Str32        byte = 0xdb

	Bin8  byte = 0xc4
	Bin16 byte = 0xc5
	Bin32 byte = 0xc6

	FixedArrayLow  byte = 0x90
	FixedArrayHigh byte = 0x9f
	FixedArrayMask byte = 0xf
	Array16        byte = 0xdc
	Array32        byte = 0xdd

	FixedMapLow  byte = 0x80
	FixedMapHigh byte = 0x8f
	FixedMapMask byte = 0xf
	Map16        byte = 0xde
	Map32        byte = 0xdf

	FixExt1  byte = 0xd4
	FixExt2  byte = 0xd5
	FixExt4  byte = 0xd6
	FixExt8  byte = 0xd7
	FixExt16 byte = 0xd8
	Ext8     byte = 0xc7
	Ext16    byte = 0xc8
	Ext32    byte = 0xc9
)

func IsFloat(c byte) bool {
	return c == Float || c == Double
}

func IsInt(c byte) bool {
	return c == Int8 || c == Int16 || c == Int32 || c == Int64
}

func IsUInt(c byte) bool {
	return c == Uint8 || c == Uint16 || c == Uint32 || c == Uint64
}

func IsBool(c byte) bool {
	return c == True || c == False
}

func IsArray(c byte) bool {
	return IsFixedArray(c) || c == Array16 || c == Array32
}

func IsMap(c byte) bool {
	return IsFixedMap(c) || c == Map16 || c == Map32
}

func IsFixedNum(c byte) bool {
	return c <= PosFixedNumHigh || c >= NegFixedNumLow
}

func IsFixedMap(c byte) bool {
	return c >= FixedMapLow && c <= FixedMapHigh
}

func IsFixedArray(c byte) bool {
	return c >= FixedArrayLow && c <= FixedArrayHigh
}

func IsFixedString(c byte) bool {
	return c >= FixedStrLow && c <= FixedStrHigh
}

func IsString(c byte) bool {
	return IsFixedString(c) || c == Str8 || c == Str16 || c == Str32
}

func IsBin(c byte) bool {
	return c == Bin8 || c == Bin16 || c == Bin32
}

func IsFixedExt(c byte) bool {
	return c >= FixExt1 && c <= FixExt16
}

func IsExt(c byte) bool {
	return IsFixedExt(c) || c == Ext8 || c == Ext16 || c == Ext32
}
