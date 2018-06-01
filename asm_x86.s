// +build amd64,go1.6

#include "textflag.h"

// func cpuid(eaxArg, ecxArg uint32) (eax, ebx, ecx, edx uint32)
TEXT ·cpuid(SB), NOSPLIT, $0-24
	MOVL eaxArg+0(FP), AX
	MOVL ecxArg+4(FP), CX
	CPUID
	MOVL AX, eax+8(FP)
	MOVL BX, ebx+12(FP)
	MOVL CX, ecx+16(FP)
	MOVL DX, edx+20(FP)
	RET

#define LATF	X0
#define LATI	R8
#define LNGF	X1
#define LNGI	R9
#define TEMP	R10
#define GHSH	R11
#define MASK	BX
#define BITS	CX
#define LATB	R12
#define LNGB	R13
#define LATE	X2
#define LNGE	X3

// func EncodeInt(lat, lng float64) uint64
TEXT ·EncodeInt(SB), NOSPLIT, $0
	CMPB ·useAsm(SB), $1
	JNE  fallback

	MOVSD lat+0(FP), LATF
	MOVSD lng+8(FP), LNGF

	MOVQ $0x5555555555555555, MASK

	MULSD $(0.005555555555555556), LATF
	ADDSD $(1.5), LATF

	MULSD $(0.002777777777777778), LNGF
	ADDSD $(1.5), LNGF

	MOVQ LNGF, LNGI
	SHRQ $20, LNGI

	MOVQ  LATF, LATI
	SHRQ  $20, LATI
	PDEPQ MASK, LATI, GHSH

	PDEPQ MASK, LNGI, TEMP

	SHLQ $1, TEMP
	XORQ TEMP, GHSH

	MOVQ GHSH, ret+16(FP)
	RET

fallback:
	JMP ·encodeInt(SB)

// func BoundingBoxIntWithPrecision(hash uint64, bits uint) Box
TEXT ·BoundingBoxIntWithPrecision(SB), NOSPLIT, $0
	CMPB ·useAsm(SB), $1
	JNE  fallback

	MOVQ hash+0(FP), GHSH
	MOVQ bits+8(FP), BITS

	RORQ BITS, GHSH

	MOVQ  $0x5555555555555555, MASK
	PEXTQ MASK, GHSH, LATI

	MOVQ  $0xaaaaaaaaaaaaaaaa, MASK
	PEXTQ MASK, GHSH, LNGI

	MOVQ $0x3ff0000000000000, MASK

	SHLQ $20, LATI
	ORQ  MASK, LATI
	MOVQ LATI, LATF

	SHLQ $20, LNGI
	ORQ  MASK, LNGI
	MOVQ LNGI, LNGF

	SUBSD $(1.5), LATF
	MULSD $(180.0), LATF

	SUBSD $(1.5), LNGF
	MULSD $(360.0), LNGF

	// Compute errors.
	MOVQ BITS, LATB
	SHRQ $1, LATB

	MOVQ  $1023, TEMP
	SUBQ  LATB, TEMP
	SHLQ  $52, TEMP
	MOVQ  TEMP, LATE
	MULSD $(180.0), LATE

	MOVQ BITS, LNGB
	SUBQ LATB, LNGB

	MOVQ  $1023, TEMP
	SUBQ  LNGB, TEMP
	SHLQ  $52, TEMP
	MOVQ  TEMP, LNGE
	MULSD $(360.0), LNGE

	MOVSD LATF, ret+16(FP)
	ADDSD LATE, LATF
	MOVSD LATF, ret+24(FP)
	MOVSD LNGF, ret+32(FP)
	ADDSD LNGE, LNGF
	MOVSD LNGF, ret+40(FP)
	RET

fallback:
	JMP ·boundingBoxIntWithPrecision(SB)
