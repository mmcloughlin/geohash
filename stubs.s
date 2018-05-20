// +build !amd64

#include "textflag.h"

TEXT ·EncodeInt(SB),NOSPLIT,$0
	JMP ·encodeInt(SB)
