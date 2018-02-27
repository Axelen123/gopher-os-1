package aml

// List of AML opcodes.
const (
	// Regular opcode list
	pOpZero             = uint16(0x00)
	pOpOne              = uint16(0x01)
	pOpAlias            = uint16(0x06)
	pOpName             = uint16(0x08)
	pOpBytePrefix       = uint16(0x0a)
	pOpWordPrefix       = uint16(0x0b)
	pOpDwordPrefix      = uint16(0x0c)
	pOpStringPrefix     = uint16(0x0d)
	pOpQwordPrefix      = uint16(0x0e)
	pOpScope            = uint16(0x10)
	pOpBuffer           = uint16(0x11)
	pOpPackage          = uint16(0x12)
	pOpVarPackage       = uint16(0x13)
	pOpMethod           = uint16(0x14)
	pOpExternal         = uint16(0x15)
	pOpLocal0           = uint16(0x60)
	pOpLocal1           = uint16(0x61)
	pOpLocal2           = uint16(0x62)
	pOpLocal3           = uint16(0x63)
	pOpLocal4           = uint16(0x64)
	pOpLocal5           = uint16(0x65)
	pOpLocal6           = uint16(0x66)
	pOpLocal7           = uint16(0x67)
	pOpArg0             = uint16(0x68)
	pOpArg1             = uint16(0x69)
	pOpArg2             = uint16(0x6a)
	pOpArg3             = uint16(0x6b)
	pOpArg4             = uint16(0x6c)
	pOpArg5             = uint16(0x6d)
	pOpArg6             = uint16(0x6e)
	pOpStore            = uint16(0x70)
	pOpRefOf            = uint16(0x71)
	pOpAdd              = uint16(0x72)
	pOpConcat           = uint16(0x73)
	pOpSubtract         = uint16(0x74)
	pOpIncrement        = uint16(0x75)
	pOpDecrement        = uint16(0x76)
	pOpMultiply         = uint16(0x77)
	pOpDivide           = uint16(0x78)
	pOpShiftLeft        = uint16(0x79)
	pOpShiftRight       = uint16(0x7a)
	pOpAnd              = uint16(0x7b)
	pOpNand             = uint16(0x7c)
	pOpOr               = uint16(0x7d)
	pOpNor              = uint16(0x7e)
	pOpXor              = uint16(0x7f)
	pOpNot              = uint16(0x80)
	pOpFindSetLeftBit   = uint16(0x81)
	pOpFindSetRightBit  = uint16(0x82)
	pOpDerefOf          = uint16(0x83)
	pOpConcatRes        = uint16(0x84)
	pOpMod              = uint16(0x85)
	pOpNotify           = uint16(0x86)
	pOpSizeOf           = uint16(0x87)
	pOpIndex            = uint16(0x88)
	pOpMatch            = uint16(0x89)
	pOpCreateDWordField = uint16(0x8a)
	pOpCreateWordField  = uint16(0x8b)
	pOpCreateByteField  = uint16(0x8c)
	pOpCreateBitField   = uint16(0x8d)
	pOpObjectType       = uint16(0x8e)
	pOpCreateQWordField = uint16(0x8f)
	pOpLand             = uint16(0x90)
	pOpLor              = uint16(0x91)
	pOpLnot             = uint16(0x92)
	pOpLEqual           = uint16(0x93)
	pOpLGreater         = uint16(0x94)
	pOpLLess            = uint16(0x95)
	pOpToBuffer         = uint16(0x96)
	pOpToDecimalString  = uint16(0x97)
	pOpToHexString      = uint16(0x98)
	pOpToInteger        = uint16(0x99)
	pOpToString         = uint16(0x9c)
	pOpCopyObject       = uint16(0x9d)
	pOpMid              = uint16(0x9e)
	pOpContinue         = uint16(0x9f)
	pOpIf               = uint16(0xa0)
	pOpElse             = uint16(0xa1)
	pOpWhile            = uint16(0xa2)
	pOpNoop             = uint16(0xa3)
	pOpReturn           = uint16(0xa4)
	pOpBreak            = uint16(0xa5)
	pOpBreakPoint       = uint16(0xcc)
	pOpOnes             = uint16(0xff)
	// Extended opcodes
	pOpMutex       = uint16(0xff + 0x01)
	pOpEvent       = uint16(0xff + 0x02)
	pOpCondRefOf   = uint16(0xff + 0x12)
	pOpCreateField = uint16(0xff + 0x13)
	pOpLoadTable   = uint16(0xff + 0x1f)
	pOpLoad        = uint16(0xff + 0x20)
	pOpStall       = uint16(0xff + 0x21)
	pOpSleep       = uint16(0xff + 0x22)
	pOpAcquire     = uint16(0xff + 0x23)
	pOpSignal      = uint16(0xff + 0x24)
	pOpWait        = uint16(0xff + 0x25)
	pOpReset       = uint16(0xff + 0x26)
	pOpRelease     = uint16(0xff + 0x27)
	pOpFromBCD     = uint16(0xff + 0x28)
	pOpToBCD       = uint16(0xff + 0x29)
	pOpUnload      = uint16(0xff + 0x2a)
	pOpRevision    = uint16(0xff + 0x30)
	pOpDebug       = uint16(0xff + 0x31)
	pOpFatal       = uint16(0xff + 0x32)
	pOpTimer       = uint16(0xff + 0x33)
	pOpOpRegion    = uint16(0xff + 0x80)
	pOpField       = uint16(0xff + 0x81)
	pOpDevice      = uint16(0xff + 0x82)
	pOpProcessor   = uint16(0xff + 0x83)
	pOpPowerRes    = uint16(0xff + 0x84)
	pOpThermalZone = uint16(0xff + 0x85)
	pOpIndexField  = uint16(0xff + 0x86)
	pOpBankField   = uint16(0xff + 0x87)
	pOpDataRegion  = uint16(0xff + 0x88)
	// Special internal opcodes which are not part of the spec; these are
	// for internal use by the AML parser.
	pOpIntScopeBlock           = uint16(0xff + 0xf7)
	pOpIntByteList             = uint16(0xff + 0xf8)
	pOpIntConnection           = uint16(0xff + 0xf9)
	pOpIntNamedField           = uint16(0xff + 0xfa)
	pOpIntResolvedNamePath     = uint16(0xff + 0xfb)
	pOpIntNamePath             = uint16(0xff + 0xfc)
	pOpIntNamePathOrMethodCall = uint16(0xff + 0xfd)
	pOpIntMethodCall           = uint16(0xff + 0xfe)
	// Sentinel value to indicate freed objects
	pOpIntFreedObject = uint16(0xff + 0xff)
)

// pOpIsLocalArg returns true if this opcode represents any of the supported local
// function args 0 to 7.
func pOpIsLocalArg(op uint16) bool {
	return op >= pOpLocal0 && op <= pOpLocal7
}

// pOpIsMethodArg returns true if this opcode represents any of the supported
// input function args 0 to 6.
func pOpIsMethodArg(op uint16) bool {
	return op >= pOpArg0 && op <= pOpArg6
}

// pOpIsArg returns true if this opcode is either a local or a method arg.
func pOpIsArg(op uint16) bool {
	return pOpIsLocalArg(op) || pOpIsMethodArg(op)
}

// pOpIsType2 returns true if this is a Type2Opcode.
//
// Grammar:
// Type2Opcode := DefAcquire | DefAdd | DefAnd | DefBuffer | DefConcat |
//  DefConcatRes | DefCondRefOf | DefCopyObject | DefDecrement |
//  DefDerefOf | DefDivide | DefFindSetLeftBit | DefFindSetRightBit |
//  DefFromBCD | DefIncrement | DefIndex | DefLAnd | DefLEqual |
//  DefLGreater | DefLGreaterEqual | DefLLess | DefLLessEqual | DefMid |
//  DefLNot | DefLNotEqual | DefLoadTable | DefLOr | DefMatch | DefMod |
//  DefMultiply | DefNAnd | DefNOr | DefNot | DefObjectType | DefOr |
//  DefPackage | DefVarPackage | DefRefOf | DefShiftLeft | DefShiftRight |
//  DefSizeOf | DefStore | DefSubtract | DefTimer | DefToBCD | DefToBuffer |
//  DefToDecimalString | DefToHexString | DefToInteger | DefToString |
//  DefWait | DefXOr
func pOpIsType2(op uint16) bool {
	switch op {
	case pOpAcquire, pOpAdd, pOpAnd, pOpBuffer, pOpConcat,
		pOpConcatRes, pOpCondRefOf, pOpCopyObject, pOpDecrement,
		pOpDerefOf, pOpDivide, pOpFindSetLeftBit, pOpFindSetRightBit,
		pOpFromBCD, pOpIncrement, pOpIndex, pOpLand, pOpLEqual,
		pOpLGreater, pOpLLess, pOpMid,
		pOpLnot, pOpLoadTable, pOpLor, pOpMatch, pOpMod,
		pOpMultiply, pOpNand, pOpNor, pOpNot, pOpObjectType, pOpOr,
		pOpPackage, pOpVarPackage, pOpRefOf, pOpShiftLeft, pOpShiftRight,
		pOpSizeOf, pOpStore, pOpSubtract, pOpTimer, pOpToBCD, pOpToBuffer,
		pOpToDecimalString, pOpToHexString, pOpToInteger, pOpToString,
		pOpWait, pOpXor:
		return true
	default:
		return false
	}
}

// pOpIsDataObject returns true if this opcode is part of a DataObject definition
//
// Grammar:
// DataObject := ComputationalData | DefPackage | DefVarPackage
// ComputationalData := ByteConst | WordConst | DWordConst | QWordConst | String | ConstObj | RevisionOp | DefBuffer
// ConstObj := ZeroOp | OneOp | OnesOp
func pOpIsDataObject(op uint16) bool {
	switch op {
	case pOpBytePrefix, pOpWordPrefix, pOpDwordPrefix, pOpQwordPrefix, pOpStringPrefix,
		pOpZero, pOpOne, pOpOnes, pOpRevision, pOpBuffer, pOpPackage, pOpVarPackage:
		return true
	default:
		return false
	}
}
