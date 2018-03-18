#include "textflag.h"

// func EncodeInt(lat, lng float64) uint64
TEXT ·EncodeInt(SB),NOSPLIT,$0
#define LATF	X0
#define LATI	R8
#define LNGF	X1
#define LNGI	R9
#define TEMP	R10
#define GHSH	R11
#define MASK	BX

	MOVSD lat+0(FP), LATF
	MOVSD lng+8(FP), LNGF

	MOVQ $0x5555555555555555, MASK

	DIVSD $(180.0), LATF
	ADDSD $(1.5), LATF

	MOVQ LATF, LATI
	SHRQ $20, LATI
	PDEPQ MASK, LATI, GHSH

	DIVSD $(360.0), LNGF
	ADDSD $(1.5), LNGF

	MOVQ LNGF, LNGI
	SHRQ $20, LNGI
	PDEPQ MASK, LNGI, TEMP

	SHLQ $1, TEMP
	XORQ TEMP, GHSH

	MOVQ GHSH, ret+16(FP)
	RET
