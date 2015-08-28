// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line grammar.y:18

// This grammar.y.go file was auto-generated by yacc from grammar.y.

package parse

import __yyfmt__ "fmt"

//line grammar.y:20
import (
	"math/big"
	"strings"
)

type intPos struct {
	int *big.Int
	pos Pos
}

type ratPos struct {
	rat *big.Rat
	pos Pos
}

type imagPos struct {
	imag *BigImag
	pos  Pos
}

// typeListToStrList converts a slice of Type to a slice of StringPos.  Each
// type must be a TypeNamed with an empty PackageName, otherwise errors are
// reported, and ok=false is returned.
func typeListToStrList(yylex yyLexer, typeList []Type) (strList []StringPos, ok bool) {
	ok = true
	for _, t := range typeList {
		var tn *TypeNamed
		if tn, ok = t.(*TypeNamed); !ok {
			lexPosErrorf(yylex, t.Pos(), "%s invalid (expected one or more variable names)", t.String())
			return
		}
		if strings.ContainsRune(tn.Name, '.') {
			ok = false
			lexPosErrorf(yylex, t.Pos(), "%s invalid (expected one or more variable names).", tn.Name)
			return
		}
		strList = append(strList, StringPos{tn.Name, tn.P})
	}
	return
}

//line grammar.y:67
type yySymType struct {
	yys        int
	pos        Pos
	strpos     StringPos
	intpos     intPos
	ratpos     ratPos
	imagpos    imagPos
	namepos    NamePos
	nameposes  []NamePos
	typeexpr   Type
	typeexprs  []Type
	fields     []*Field
	iface      *Interface
	constexpr  ConstExpr
	constexprs []ConstExpr
	complit    *ConstCompositeLit
	kvlit      KVLit
	kvlits     []KVLit
	errordef   ErrorDef
}

const startFileImports = 57346
const startFile = 57347
const startConfigImports = 57348
const startConfig = 57349
const startExprs = 57350
const tOROR = 57351
const tANDAND = 57352
const tLE = 57353
const tGE = 57354
const tNE = 57355
const tEQEQ = 57356
const tLSH = 57357
const tRSH = 57358
const tCONST = 57359
const tENUM = 57360
const tERROR = 57361
const tIMPORT = 57362
const tINTERFACE = 57363
const tMAP = 57364
const tPACKAGE = 57365
const tSET = 57366
const tSTREAM = 57367
const tSTRUCT = 57368
const tTYPE = 57369
const tTYPEOBJECT = 57370
const tUNION = 57371
const tIDENT = 57372
const tSTRLIT = 57373
const tINTLIT = 57374
const tRATLIT = 57375
const tIMAGLIT = 57376
const notPackage = 57377
const notConfig = 57378

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"startFileImports",
	"startFile",
	"startConfigImports",
	"startConfig",
	"startExprs",
	"';'",
	"':'",
	"','",
	"'.'",
	"'('",
	"')'",
	"'['",
	"']'",
	"'{'",
	"'}'",
	"'<'",
	"'>'",
	"'='",
	"'!'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'|'",
	"'&'",
	"'^'",
	"'?'",
	"tOROR",
	"tANDAND",
	"tLE",
	"tGE",
	"tNE",
	"tEQEQ",
	"tLSH",
	"tRSH",
	"tCONST",
	"tENUM",
	"tERROR",
	"tIMPORT",
	"tINTERFACE",
	"tMAP",
	"tPACKAGE",
	"tSET",
	"tSTREAM",
	"tSTRUCT",
	"tTYPE",
	"tTYPEOBJECT",
	"tUNION",
	"tIDENT",
	"tSTRLIT",
	"tINTLIT",
	"tRATLIT",
	"tIMAGLIT",
	"notPackage",
	"notConfig",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 27,
	13, 40,
	17, 40,
	-2, 118,
	-1, 176,
	18, 149,
	-2, 144,
	-1, 291,
	18, 149,
	-2, 144,
}

const yyNprod = 150
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 800

var yyAct = [...]int{

	14, 27, 13, 140, 139, 242, 251, 235, 220, 151,
	186, 193, 175, 137, 147, 38, 190, 252, 253, 133,
	166, 243, 90, 153, 152, 273, 181, 194, 191, 210,
	78, 21, 134, 217, 192, 249, 144, 127, 80, 89,
	40, 43, 213, 11, 240, 270, 8, 102, 98, 178,
	104, 105, 106, 107, 108, 109, 110, 111, 112, 113,
	114, 115, 116, 117, 118, 119, 120, 121, 122, 87,
	252, 253, 194, 123, 194, 89, 126, 79, 128, 243,
	90, 191, 89, 150, 89, 89, 89, 89, 153, 152,
	125, 207, 297, 148, 189, 185, 158, 276, 159, 142,
	41, 296, 93, 293, 292, 87, 157, 216, 100, 46,
	289, 29, 87, 31, 87, 87, 87, 87, 295, 288,
	17, 18, 19, 153, 152, 298, 275, 277, 20, 37,
	210, 271, 89, 267, 191, 188, 202, 201, 86, 32,
	30, 198, 89, 34, 195, 33, 279, 35, 101, 22,
	36, 40, 23, 24, 25, 26, 103, 233, 170, 226,
	180, 92, 87, 174, 85, 84, 60, 61, 62, 169,
	64, 89, 87, 89, 124, 89, 199, 148, 204, 66,
	67, 131, 81, 135, 136, 141, 141, 196, 168, 15,
	89, 203, 96, 130, 97, 98, 209, 83, 82, 75,
	205, 87, 95, 87, 219, 87, 215, 68, 69, 70,
	71, 76, 89, 221, 77, 300, 223, 229, 299, 266,
	87, 264, 247, 89, 245, 236, 237, 238, 244, 224,
	222, 165, 162, 73, 72, 227, 239, 91, 89, 230,
	74, 172, 87, 248, 48, 246, 49, 291, 254, 257,
	255, 262, 260, 87, 176, 263, 259, 89, 89, 231,
	265, 244, 258, 261, 228, 268, 225, 206, 87, 272,
	197, 184, 141, 183, 200, 182, 171, 167, 89, 89,
	281, 237, 285, 89, 280, 218, 149, 87, 87, 211,
	99, 89, 104, 286, 287, 214, 187, 208, 290, 58,
	59, 60, 61, 62, 63, 64, 65, 42, 87, 87,
	10, 211, 156, 87, 66, 67, 12, 44, 45, 155,
	47, 87, 141, 2, 3, 4, 5, 6, 154, 7,
	29, 179, 31, 9, 94, 284, 1, 172, 250, 17,
	18, 19, 232, 146, 28, 278, 16, 20, 37, 241,
	274, 269, 39, 132, 0, 0, 141, 200, 32, 30,
	0, 0, 34, 0, 33, 0, 35, 0, 22, 36,
	40, 23, 24, 25, 26, 0, 0, 141, 283, 0,
	0, 0, 172, 29, 0, 31, 0, 0, 145, 0,
	294, 0, 17, 18, 19, 0, 0, 0, 256, 0,
	20, 37, 31, 0, 0, 0, 0, 0, 0, 0,
	0, 32, 30, 0, 0, 34, 0, 33, 37, 35,
	0, 22, 36, 40, 23, 24, 25, 26, 32, 30,
	31, 0, 34, 0, 33, 282, 35, 0, 88, 36,
	40, 90, 0, 0, 234, 31, 37, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 32, 30, 0, 0,
	34, 37, 33, 0, 35, 0, 88, 36, 40, 90,
	31, 32, 30, 0, 0, 34, 0, 33, 0, 35,
	0, 88, 36, 40, 90, 173, 37, 0, 0, 31,
	0, 0, 0, 0, 0, 0, 32, 30, 0, 212,
	34, 0, 33, 0, 35, 37, 88, 36, 40, 90,
	0, 0, 0, 0, 31, 32, 30, 143, 0, 34,
	0, 33, 0, 35, 0, 88, 36, 40, 90, 31,
	37, 0, 138, 0, 0, 0, 0, 0, 0, 0,
	32, 30, 0, 0, 34, 37, 33, 0, 35, 0,
	88, 36, 40, 90, 31, 32, 30, 0, 0, 34,
	0, 33, 0, 35, 0, 88, 36, 40, 90, 0,
	37, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	32, 30, 0, 0, 34, 177, 33, 0, 35, 0,
	88, 36, 40, 90, 52, 53, 0, 0, 58, 59,
	60, 61, 62, 63, 64, 65, 0, 50, 51, 54,
	55, 56, 57, 66, 67, 164, 0, 0, 52, 53,
	0, 0, 58, 59, 60, 61, 62, 63, 64, 65,
	0, 50, 51, 54, 55, 56, 57, 66, 67, 163,
	0, 0, 52, 53, 0, 0, 58, 59, 60, 61,
	62, 63, 64, 65, 0, 50, 51, 54, 55, 56,
	57, 66, 67, 161, 0, 0, 0, 0, 52, 53,
	0, 0, 58, 59, 60, 61, 62, 63, 64, 65,
	160, 50, 51, 54, 55, 56, 57, 66, 67, 0,
	52, 53, 0, 0, 58, 59, 60, 61, 62, 63,
	64, 65, 0, 50, 51, 54, 55, 56, 57, 66,
	67, 129, 0, 0, 0, 0, 52, 53, 0, 0,
	58, 59, 60, 61, 62, 63, 64, 65, 0, 50,
	51, 54, 55, 56, 57, 66, 67, 52, 53, 0,
	0, 58, 59, 60, 61, 62, 63, 64, 65, 0,
	50, 51, 54, 55, 56, 57, 66, 67, 52, 53,
	0, 0, 58, 59, 60, 61, 62, 63, 64, 65,
	0, 0, 51, 54, 55, 56, 57, 66, 67, 52,
	53, 0, 0, 58, 59, 60, 61, 62, 63, 64,
	65, 0, 0, 0, 54, 55, 56, 57, 66, 67,
}
var yyPact = [...]int{

	319, -1000, 0, 0, -10, -10, 98, -1000, -12, -1000,
	-1000, 88, -1000, 235, 718, -1000, -1000, 98, 98, 98,
	98, 221, 220, 228, -1000, -1000, -1000, 184, 199, 98,
	-1000, 22, 165, 183, 182, 148, 147, 539, 225, 144,
	-1000, -1000, 152, 281, 5, 152, 98, 5, -1000, 98,
	98, 98, 98, 98, 98, 98, 98, 98, 98, 98,
	98, 98, 98, 98, 98, 98, 98, 98, -1000, -1000,
	-1000, -1000, 98, 539, -13, 98, -16, 98, 697, 177,
	539, -21, 539, 539, 514, 499, -1000, -1000, -1000, -1000,
	228, -17, 370, -1000, 277, -1000, -1000, -1000, 70, -1000,
	56, -1000, 671, 56, 718, 739, 760, 276, 276, 276,
	276, 276, 276, 141, 141, -1000, -1000, -1000, 141, -1000,
	141, -1000, -1000, 649, 218, 225, 623, -1000, 599, -1000,
	539, -1000, 268, -1000, -1000, 172, 153, 267, -1000, -1000,
	474, -1000, 267, -1000, -1000, -1000, 243, -1000, 575, -1000,
	35, -1000, -1000, -28, 266, 264, 262, 82, 81, 21,
	-1000, -1000, -1000, -1000, -1000, -1000, 126, -21, -1000, 539,
	123, 539, -1000, 539, 119, 118, 98, 98, -1000, 258,
	-1000, -1000, -1000, -1000, -1000, 77, -1000, -1000, 455, 28,
	-1000, 86, 19, -1000, 200, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 718, 216, -30, -1000, 257, -1000,
	539, -1000, 142, -1000, 255, -1000, 98, -1000, 250, -1000,
	140, 430, -1000, -1000, 213, -24, 26, 210, -25, 718,
	208, -26, -1000, 17, -1000, 239, 387, -1000, -1000, -1000,
	-1000, 247, -1000, 200, -1000, -1000, -1000, -1000, -1000, -1000,
	240, -1000, -1000, 245, 207, 539, 539, 205, 115, -32,
	-3, 113, -36, -29, -1000, -1000, -1000, -1000, -1000, 84,
	108, -1000, -1000, -1000, 129, -1000, 539, 415, -1000, 317,
	239, 387, -1000, 99, -1000, 236, 76, 75, -1000, 539,
	100, 98, 59, 50, 105, -1000, 204, 201, -1000, -1000,
	-1000,
}
var yyPgo = [...]int{

	0, 1, 15, 19, 353, 100, 31, 352, 3, 351,
	13, 4, 7, 8, 350, 349, 5, 0, 189, 346,
	345, 2, 344, 14, 343, 342, 338, 6, 336, 329,
	307, 102, 108, 310, 334, 331, 20, 9, 328, 319,
	312, 297, 10, 296, 295, 16, 285, 11, 12,
}
var yyR1 = [...]int{

	0, 28, 28, 28, 28, 28, 31, 31, 31, 31,
	29, 29, 33, 33, 30, 30, 34, 34, 34, 35,
	35, 37, 37, 32, 32, 32, 32, 38, 38, 38,
	38, 39, 39, 39, 40, 40, 40, 41, 41, 42,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 5, 5, 4, 4, 3, 10, 10, 11,
	8, 8, 43, 43, 15, 15, 16, 16, 13, 13,
	13, 12, 12, 14, 14, 14, 9, 9, 9, 9,
	20, 20, 20, 21, 21, 44, 44, 45, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 18, 18, 18,
	18, 18, 18, 18, 19, 19, 19, 19, 19, 19,
	19, 19, 19, 19, 22, 22, 24, 24, 23, 23,
	46, 46, 47, 25, 25, 25, 26, 26, 27, 27,
	1, 1, 2, 2, 7, 7, 36, 36, 48, 48,
}
var yyR2 = [...]int{

	0, 4, 4, 4, 4, 3, 0, 1, 1, 1,
	0, 3, 0, 4, 0, 3, 3, 5, 2, 1,
	3, 1, 2, 0, 3, 3, 3, 3, 5, 2,
	2, 3, 5, 2, 3, 5, 2, 1, 3, 2,
	1, 1, 4, 3, 5, 4, 5, 5, 3, 5,
	3, 2, 1, 1, 1, 3, 1, 1, 3, 2,
	1, 3, 4, 6, 1, 3, 5, 1, 2, 4,
	4, 1, 3, 1, 6, 6, 0, 3, 4, 6,
	0, 2, 4, 1, 3, 1, 3, 3, 1, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 1, 2, 2,
	2, 2, 4, 4, 1, 1, 1, 1, 1, 1,
	3, 4, 4, 3, 3, 5, 1, 3, 1, 3,
	1, 3, 3, 0, 2, 4, 1, 3, 1, 3,
	1, 3, 1, 3, 0, 1, 0, 1, 0, 1,
}
var yyChk = [...]int{

	-1000, -28, 4, 5, 6, 7, 8, -29, 46, -29,
	-33, 53, -33, -21, -17, -18, -19, 22, 23, 24,
	30, -6, 51, 54, 55, 56, 57, -1, -22, 13,
	42, 15, 41, 47, 45, 49, 52, 31, -2, -7,
	53, -5, -30, 53, -30, -30, 21, -30, 9, 11,
	32, 33, 19, 20, 34, 35, 36, 37, 23, 24,
	25, 26, 27, 28, 29, 30, 38, 39, -18, -18,
	-18, -18, 13, 13, 12, 15, 12, 15, -17, 55,
	16, 17, 15, 15, 17, 17, -5, -6, 51, -1,
	54, 12, 17, -31, -34, 50, 40, 42, 43, 9,
	-32, -31, -17, -32, -17, -17, -17, -17, -17, -17,
	-17, -17, -17, -17, -17, -17, -17, -17, -17, -17,
	-17, -17, -17, -17, -5, -2, -17, 53, -17, 14,
	16, -5, -4, -3, 53, -5, -5, -10, 18, -11,
	-8, -5, -10, 18, 53, 18, -24, -23, -17, 9,
	13, -37, 54, 53, -38, -39, -40, 50, 40, 42,
	9, 14, 14, 16, 16, -5, -36, 9, 16, 16,
	-36, 9, -5, 11, -36, -48, 11, 10, 14, -35,
	-37, 54, 9, 9, 9, 13, -42, -43, 53, 13,
	-45, 53, 13, -47, 53, 18, -3, -5, 18, -11,
	-5, 18, 18, -23, -17, -36, 9, 14, -41, -42,
	53, -5, 44, 14, -44, -45, 21, 14, -46, -47,
	-13, 13, 14, -37, -36, 9, 17, -36, 9, -17,
	-36, 9, -25, 17, 14, -12, -8, -11, 14, -42,
	18, -15, -16, 53, -1, 14, -45, 14, -47, 18,
	-26, -27, 53, 54, -48, 11, 11, -48, -36, 9,
	-13, -48, 11, 10, 14, -11, 14, 18, -16, -9,
	48, 18, -27, 54, -14, 42, 13, 19, -20, 17,
	-12, -8, 20, -5, 18, -21, -48, -48, 20, 11,
	-48, 11, 28, 28, -5, 18, 42, 42, 20, 14,
	14,
}
var yyDef = [...]int{

	0, -2, 10, 10, 12, 12, 144, 14, 0, 14,
	14, 0, 14, 0, 83, 88, 107, 144, 144, 144,
	144, 52, 53, 114, 115, 116, 117, -2, 119, 144,
	41, 0, 0, 0, 0, 0, 0, 0, 140, 0,
	142, 145, 6, 0, 23, 6, 144, 23, 5, 144,
	144, 144, 144, 144, 144, 144, 144, 144, 144, 144,
	144, 144, 144, 144, 144, 144, 144, 144, 108, 109,
	110, 111, 144, 0, 0, 144, 0, 144, 0, 0,
	0, 0, 0, 0, 0, 0, 51, 52, 53, 40,
	0, 0, 144, 1, 0, 7, 8, 9, 0, 11,
	2, 3, 0, 4, 84, 89, 90, 91, 92, 93,
	94, 95, 96, 97, 98, 99, 100, 101, 102, 103,
	104, 105, 106, 0, 0, 141, 0, 120, 0, 123,
	0, 43, 146, 54, 56, 0, 0, 146, 48, 57,
	0, 60, 146, 50, 143, 124, 148, 126, 128, 15,
	0, 18, 21, 0, 0, 0, 0, 0, 0, 0,
	13, 112, 113, 122, 121, 42, 0, 147, 45, 0,
	0, 147, 59, 0, 0, 0, -2, 144, 16, 146,
	19, 22, 24, 25, 26, 0, 29, 30, 0, 0,
	33, 0, 0, 36, 0, 44, 55, 46, 47, 58,
	61, 49, 125, 127, 129, 0, 147, 27, 146, 37,
	0, 39, 0, 31, 146, 85, 144, 34, 146, 130,
	133, 0, 17, 20, 0, 147, 0, 0, 147, 87,
	0, 147, 132, 0, 68, 148, 148, 71, 28, 38,
	62, 146, 64, 142, 67, 32, 86, 35, 131, 134,
	148, 136, 138, 0, 0, 149, 149, 0, 0, 147,
	76, 0, 149, 0, 69, 72, 70, 63, 65, 0,
	0, 135, 137, 139, 80, 73, 0, 0, 66, 144,
	148, 148, 77, 0, 81, 148, 0, 0, 78, 0,
	0, -2, 0, 0, 0, 82, 0, 0, 79, 74,
	75,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 22, 3, 3, 3, 27, 29, 3,
	13, 14, 25, 23, 11, 24, 12, 26, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 10, 9,
	19, 21, 20, 31, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 15, 3, 16, 30, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 17, 28, 18,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 32, 33, 34,
	35, 36, 37, 38, 39, 40, 41, 42, 43, 44,
	45, 46, 47, 48, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lookahead func() int
}

func (p *yyParserImpl) Lookahead() int {
	return p.lookahead()
}

func yyNewParser() yyParser {
	p := &yyParserImpl{
		lookahead: func() int { return -1 },
	}
	return p
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yytoken := -1 // yychar translated into internal numbering
	yyrcvr.lookahead = func() int { return yychar }
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yychar = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar, yytoken = yylex1(yylex, &yylval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yychar = -1
		yytoken = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar, yytoken = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yychar = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:143
		{
			lexStoreExprs(yylex, yyDollar[2].constexprs)
		}
	case 6:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line grammar.y:152
		{
			lexGenEOF(yylex)
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:154
		{
			lexGenEOF(yylex)
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:156
		{
			lexGenEOF(yylex)
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:158
		{
			lexGenEOF(yylex)
		}
	case 10:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line grammar.y:163
		{
			lexPosErrorf(yylex, Pos{}, "vdl file must start with package clause")
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:165
		{
			lexVDLFile(yylex).PackageDef = NamePos{Name: yyDollar[2].strpos.String, Pos: yyDollar[2].strpos.Pos}
		}
	case 12:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line grammar.y:170
		{
			lexPosErrorf(yylex, Pos{}, "config file must start with config clause")
		}
	case 13:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line grammar.y:172
		{
			// We allow "config" as an identifier; it is not a keyword.  So we check
			// manually to make sure the syntax is correct.
			if yyDollar[1].strpos.String != "config" {
				lexPosErrorf(yylex, yyDollar[1].strpos.Pos, "config file must start with config clause")
				return 1 // Any non-zero code indicates an error
			}
			file := lexVDLFile(yylex)
			file.PackageDef = NamePos{Name: "config", Pos: yyDollar[1].strpos.Pos}
			file.ConstDefs = []*ConstDef{{Expr: yyDollar[3].constexpr}}
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:200
		{
			imps := &lexVDLFile(yylex).Imports
			*imps = append(*imps, &Import{Path: yyDollar[1].strpos.String, NamePos: NamePos{Pos: yyDollar[1].strpos.Pos}})
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammar.y:205
		{
			imps := &lexVDLFile(yylex).Imports
			*imps = append(*imps, &Import{Path: yyDollar[2].strpos.String, NamePos: NamePos{Name: yyDollar[1].strpos.String, Pos: yyDollar[1].strpos.Pos}})
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammar.y:240
		{
			tds := &lexVDLFile(yylex).TypeDefs
			*tds = append(*tds, &TypeDef{Type: yyDollar[2].typeexpr, NamePos: NamePos{Name: yyDollar[1].strpos.String, Pos: yyDollar[1].strpos.Pos}})
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:262
		{
			yyVAL.typeexpr = &TypeNamed{Name: yyDollar[1].strpos.String, P: yyDollar[1].strpos.Pos}
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:264
		{
			yyVAL.typeexpr = &TypeNamed{Name: "error", P: yyDollar[1].pos}
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line grammar.y:266
		{
			yyVAL.typeexpr = &TypeArray{Len: int(yyDollar[2].intpos.int.Int64()), Elem: yyDollar[4].typeexpr, P: yyDollar[1].pos}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:268
		{
			yyVAL.typeexpr = &TypeList{Elem: yyDollar[3].typeexpr, P: yyDollar[1].pos}
		}
	case 44:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line grammar.y:270
		{
			yyVAL.typeexpr = &TypeEnum{Labels: yyDollar[3].nameposes, P: yyDollar[1].pos}
		}
	case 45:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line grammar.y:272
		{
			yyVAL.typeexpr = &TypeSet{Key: yyDollar[3].typeexpr, P: yyDollar[1].pos}
		}
	case 46:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line grammar.y:274
		{
			yyVAL.typeexpr = &TypeMap{Key: yyDollar[3].typeexpr, Elem: yyDollar[5].typeexpr, P: yyDollar[1].pos}
		}
	case 47:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line grammar.y:276
		{
			yyVAL.typeexpr = &TypeStruct{Fields: yyDollar[3].fields, P: yyDollar[1].pos}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:278
		{
			yyVAL.typeexpr = &TypeStruct{P: yyDollar[1].pos}
		}
	case 49:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line grammar.y:280
		{
			yyVAL.typeexpr = &TypeUnion{Fields: yyDollar[3].fields, P: yyDollar[1].pos}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:282
		{
			yyVAL.typeexpr = &TypeUnion{P: yyDollar[1].pos}
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammar.y:284
		{
			yyVAL.typeexpr = &TypeOptional{Base: yyDollar[2].typeexpr, P: yyDollar[1].pos}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:289
		{
			yyVAL.typeexpr = yyDollar[1].typeexpr
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:291
		{
			yyVAL.typeexpr = &TypeNamed{Name: "typeobject", P: yyDollar[1].pos}
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:295
		{
			yyVAL.nameposes = []NamePos{yyDollar[1].namepos}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:297
		{
			yyVAL.nameposes = append(yyDollar[1].nameposes, yyDollar[3].namepos)
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:301
		{
			yyVAL.namepos = NamePos{Name: yyDollar[1].strpos.String, Pos: yyDollar[1].strpos.Pos}
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:305
		{
			yyVAL.fields = yyDollar[1].fields
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:307
		{
			yyVAL.fields = append(yyDollar[1].fields, yyDollar[3].fields...)
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammar.y:345
		{
			if names, ok := typeListToStrList(yylex, yyDollar[1].typeexprs); ok {
				for _, n := range names {
					yyVAL.fields = append(yyVAL.fields, &Field{Type: yyDollar[2].typeexpr, NamePos: NamePos{Name: n.String, Pos: n.Pos}})
				}
			} else {
				lexPosErrorf(yylex, yyDollar[2].typeexpr.Pos(), "perhaps you forgot a comma before %q?.", yyDollar[2].typeexpr.String())
			}
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:357
		{
			yyVAL.typeexprs = []Type{yyDollar[1].typeexpr}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:359
		{
			yyVAL.typeexprs = append(yyDollar[1].typeexprs, yyDollar[3].typeexpr)
		}
	case 62:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line grammar.y:364
		{
			ifs := &lexVDLFile(yylex).Interfaces
			*ifs = append(*ifs, &Interface{NamePos: NamePos{Name: yyDollar[1].strpos.String, Pos: yyDollar[1].strpos.Pos}})
		}
	case 63:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line grammar.y:369
		{
			yyDollar[4].iface.Name, yyDollar[4].iface.Pos = yyDollar[1].strpos.String, yyDollar[1].strpos.Pos
			ifs := &lexVDLFile(yylex).Interfaces
			*ifs = append(*ifs, yyDollar[4].iface)
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:377
		{
			yyVAL.iface = yyDollar[1].iface
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:379
		{
			yyDollar[1].iface.Embeds = append(yyDollar[1].iface.Embeds, yyDollar[3].iface.Embeds...)
			yyDollar[1].iface.Methods = append(yyDollar[1].iface.Methods, yyDollar[3].iface.Methods...)
			yyVAL.iface = yyDollar[1].iface
		}
	case 66:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line grammar.y:387
		{
			yyVAL.iface = &Interface{Methods: []*Method{{InArgs: yyDollar[2].fields, InStream: yyDollar[3].typeexprs[0], OutStream: yyDollar[3].typeexprs[1], OutArgs: yyDollar[4].fields, Tags: yyDollar[5].constexprs, NamePos: NamePos{Name: yyDollar[1].strpos.String, Pos: yyDollar[1].strpos.Pos}}}}
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:389
		{
			yyVAL.iface = &Interface{Embeds: []*NamePos{{Name: yyDollar[1].strpos.String, Pos: yyDollar[1].strpos.Pos}}}
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammar.y:393
		{
			yyVAL.fields = nil
		}
	case 69:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line grammar.y:395
		{
			yyVAL.fields = yyDollar[2].fields
		}
	case 70:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line grammar.y:399
		{
			for _, t := range yyDollar[2].typeexprs {
				yyVAL.fields = append(yyVAL.fields, &Field{Type: t, NamePos: NamePos{Pos: t.Pos()}})
			}
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:409
		{
			yyVAL.fields = yyDollar[1].fields
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:411
		{
			yyVAL.fields = append(yyDollar[1].fields, yyDollar[3].fields...)
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:422
		{
			yyVAL.fields = nil
		}
	case 74:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line grammar.y:424
		{
			yyVAL.fields = yyDollar[2].fields
		}
	case 75:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line grammar.y:428
		{
			for _, t := range yyDollar[2].typeexprs {
				yyVAL.fields = append(yyVAL.fields, &Field{Type: t, NamePos: NamePos{Pos: t.Pos()}})
			}
		}
	case 76:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line grammar.y:436
		{
			yyVAL.typeexprs = []Type{nil, nil}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:438
		{
			yyVAL.typeexprs = []Type{nil, nil}
		}
	case 78:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line grammar.y:440
		{
			yyVAL.typeexprs = []Type{yyDollar[3].typeexpr, nil}
		}
	case 79:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line grammar.y:442
		{
			yyVAL.typeexprs = []Type{yyDollar[3].typeexpr, yyDollar[5].typeexpr}
		}
	case 80:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line grammar.y:446
		{
			yyVAL.constexprs = nil
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammar.y:448
		{
			yyVAL.constexprs = nil
		}
	case 82:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line grammar.y:450
		{
			yyVAL.constexprs = yyDollar[2].constexprs
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:454
		{
			yyVAL.constexprs = []ConstExpr{yyDollar[1].constexpr}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:456
		{
			yyVAL.constexprs = append(yyDollar[1].constexprs, yyDollar[3].constexpr)
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:465
		{
			cds := &lexVDLFile(yylex).ConstDefs
			*cds = append(*cds, &ConstDef{Expr: yyDollar[3].constexpr, NamePos: NamePos{Name: yyDollar[1].strpos.String, Pos: yyDollar[1].strpos.Pos}})
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:472
		{
			yyVAL.constexpr = yyDollar[1].constexpr
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:474
		{
			yyVAL.constexpr = &ConstBinaryOp{"||", yyDollar[1].constexpr, yyDollar[3].constexpr, yyDollar[2].pos}
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:476
		{
			yyVAL.constexpr = &ConstBinaryOp{"&&", yyDollar[1].constexpr, yyDollar[3].constexpr, yyDollar[2].pos}
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:478
		{
			yyVAL.constexpr = &ConstBinaryOp{"<", yyDollar[1].constexpr, yyDollar[3].constexpr, yyDollar[2].pos}
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:480
		{
			yyVAL.constexpr = &ConstBinaryOp{">", yyDollar[1].constexpr, yyDollar[3].constexpr, yyDollar[2].pos}
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:482
		{
			yyVAL.constexpr = &ConstBinaryOp{"<=", yyDollar[1].constexpr, yyDollar[3].constexpr, yyDollar[2].pos}
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:484
		{
			yyVAL.constexpr = &ConstBinaryOp{">=", yyDollar[1].constexpr, yyDollar[3].constexpr, yyDollar[2].pos}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:486
		{
			yyVAL.constexpr = &ConstBinaryOp{"!=", yyDollar[1].constexpr, yyDollar[3].constexpr, yyDollar[2].pos}
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:488
		{
			yyVAL.constexpr = &ConstBinaryOp{"==", yyDollar[1].constexpr, yyDollar[3].constexpr, yyDollar[2].pos}
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:490
		{
			yyVAL.constexpr = &ConstBinaryOp{"+", yyDollar[1].constexpr, yyDollar[3].constexpr, yyDollar[2].pos}
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:492
		{
			yyVAL.constexpr = &ConstBinaryOp{"-", yyDollar[1].constexpr, yyDollar[3].constexpr, yyDollar[2].pos}
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:494
		{
			yyVAL.constexpr = &ConstBinaryOp{"*", yyDollar[1].constexpr, yyDollar[3].constexpr, yyDollar[2].pos}
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:496
		{
			yyVAL.constexpr = &ConstBinaryOp{"/", yyDollar[1].constexpr, yyDollar[3].constexpr, yyDollar[2].pos}
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:498
		{
			yyVAL.constexpr = &ConstBinaryOp{"%", yyDollar[1].constexpr, yyDollar[3].constexpr, yyDollar[2].pos}
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:500
		{
			yyVAL.constexpr = &ConstBinaryOp{"|", yyDollar[1].constexpr, yyDollar[3].constexpr, yyDollar[2].pos}
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:502
		{
			yyVAL.constexpr = &ConstBinaryOp{"&", yyDollar[1].constexpr, yyDollar[3].constexpr, yyDollar[2].pos}
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:504
		{
			yyVAL.constexpr = &ConstBinaryOp{"^", yyDollar[1].constexpr, yyDollar[3].constexpr, yyDollar[2].pos}
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:506
		{
			yyVAL.constexpr = &ConstBinaryOp{"<<", yyDollar[1].constexpr, yyDollar[3].constexpr, yyDollar[2].pos}
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:508
		{
			yyVAL.constexpr = &ConstBinaryOp{">>", yyDollar[1].constexpr, yyDollar[3].constexpr, yyDollar[2].pos}
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:512
		{
			yyVAL.constexpr = yyDollar[1].constexpr
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammar.y:514
		{
			yyVAL.constexpr = &ConstUnaryOp{"!", yyDollar[2].constexpr, yyDollar[1].pos}
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammar.y:516
		{
			yyVAL.constexpr = &ConstUnaryOp{"+", yyDollar[2].constexpr, yyDollar[1].pos}
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammar.y:518
		{
			yyVAL.constexpr = &ConstUnaryOp{"-", yyDollar[2].constexpr, yyDollar[1].pos}
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammar.y:520
		{
			yyVAL.constexpr = &ConstUnaryOp{"^", yyDollar[2].constexpr, yyDollar[1].pos}
		}
	case 112:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line grammar.y:522
		{
			yyVAL.constexpr = &ConstTypeConv{yyDollar[1].typeexpr, yyDollar[3].constexpr, yyDollar[1].typeexpr.Pos()}
		}
	case 113:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line grammar.y:524
		{
			yyVAL.constexpr = &ConstTypeObject{yyDollar[3].typeexpr, yyDollar[1].pos}
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:529
		{
			yyVAL.constexpr = &ConstLit{yyDollar[1].strpos.String, yyDollar[1].strpos.Pos}
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:531
		{
			yyVAL.constexpr = &ConstLit{yyDollar[1].intpos.int, yyDollar[1].intpos.pos}
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:533
		{
			yyVAL.constexpr = &ConstLit{yyDollar[1].ratpos.rat, yyDollar[1].ratpos.pos}
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:535
		{
			yyVAL.constexpr = &ConstLit{yyDollar[1].imagpos.imag, yyDollar[1].imagpos.pos}
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:537
		{
			yyVAL.constexpr = &ConstNamed{yyDollar[1].strpos.String, yyDollar[1].strpos.Pos}
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:539
		{
			yyVAL.constexpr = yyDollar[1].complit
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:541
		{
			lexPosErrorf(yylex, yyDollar[2].pos, "cannot apply selector operator to unnamed constant")
		}
	case 121:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line grammar.y:543
		{
			lexPosErrorf(yylex, yyDollar[2].pos, "cannot apply index operator to unnamed constant")
		}
	case 122:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line grammar.y:545
		{
			yyVAL.constexpr = &ConstIndexed{&ConstNamed{yyDollar[1].strpos.String, yyDollar[1].strpos.Pos}, yyDollar[3].constexpr, yyDollar[1].strpos.Pos}
		}
	case 123:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:547
		{
			yyVAL.constexpr = yyDollar[2].constexpr
		}
	case 124:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:551
		{
			yyVAL.complit = &ConstCompositeLit{yyDollar[1].typeexpr, nil, yyDollar[2].pos}
		}
	case 125:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line grammar.y:553
		{
			yyVAL.complit = &ConstCompositeLit{yyDollar[1].typeexpr, yyDollar[3].kvlits, yyDollar[2].pos}
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:557
		{
			yyVAL.kvlits = []KVLit{yyDollar[1].kvlit}
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:559
		{
			yyVAL.kvlits = append(yyDollar[1].kvlits, yyDollar[3].kvlit)
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:563
		{
			yyVAL.kvlit = KVLit{Value: yyDollar[1].constexpr}
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:565
		{
			yyVAL.kvlit = KVLit{Key: yyDollar[1].constexpr, Value: yyDollar[3].constexpr}
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:574
		{
			// Create *ErrorDef starting with a copy of error_details, filling in the
			// name and params
			ed := yyDollar[3].errordef
			ed.NamePos = NamePos{Name: yyDollar[1].strpos.String, Pos: yyDollar[1].strpos.Pos}
			ed.Params = yyDollar[2].fields
			eds := &lexVDLFile(yylex).ErrorDefs
			*eds = append(*eds, &ed)
		}
	case 133:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line grammar.y:586
		{
			yyVAL.errordef = ErrorDef{}
		}
	case 134:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammar.y:588
		{
			yyVAL.errordef = ErrorDef{}
		}
	case 135:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line grammar.y:590
		{
			yyVAL.errordef = yyDollar[2].errordef
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:594
		{
			yyVAL.errordef = yyDollar[1].errordef
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:596
		{
			// Merge each ErrorDef in-order to build the final ErrorDef.
			yyVAL.errordef = yyDollar[1].errordef
			switch {
			case len(yyDollar[3].errordef.Actions) > 0:
				yyVAL.errordef.Actions = append(yyVAL.errordef.Actions, yyDollar[3].errordef.Actions...)
			case len(yyDollar[3].errordef.Formats) > 0:
				yyVAL.errordef.Formats = append(yyVAL.errordef.Formats, yyDollar[3].errordef.Formats...)
			}
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:609
		{
			yyVAL.errordef = ErrorDef{Actions: []StringPos{yyDollar[1].strpos}}
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:611
		{
			yyVAL.errordef = ErrorDef{Formats: []LangFmt{{Lang: yyDollar[1].strpos, Fmt: yyDollar[3].strpos}}}
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:623
		{
			yyVAL.strpos = yyDollar[1].strpos
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:625
		{
			yyVAL.strpos = StringPos{"\"" + yyDollar[1].strpos.String + "\"." + yyDollar[3].strpos.String, yyDollar[1].strpos.Pos}
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:630
		{
			yyVAL.strpos = yyDollar[1].strpos
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:632
		{
			yyVAL.strpos = StringPos{yyDollar[1].strpos.String + "." + yyDollar[3].strpos.String, yyDollar[1].strpos.Pos}
		}
	case 144:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line grammar.y:636
		{
			yyVAL.typeexpr = nil
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:638
		{
			yyVAL.typeexpr = yyDollar[1].typeexpr
		}
	}
	goto yystack /* stack new state and value */
}
