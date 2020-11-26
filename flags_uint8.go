package gobitflags

import (
	"errors"
)

func HasFlagByte(flags byte, flag uint8) (bool, error) {
	return HasFlagUint8(flags, flag)
}

func SetFlagByte(flags uint8, flag uint8, set bool) (uint8, error) {
	return SetFlagUint8(flags, flag, set)
}

func HasFlagUint8(flags uint8, flag uint8) (bool, error) {
	if flag > FlagMaxInt8 {
		return false, errors.New(ErrorMsgOutOfRange)
	}

	conv := uint8(1 << flag)
	return flags&conv == conv, nil
}

func SetFlagUint8(flags uint8, flag uint8, set bool) (uint8, error) {
	if flag > FlagMaxInt8 {
		return flags, errors.New(ErrorMsgOutOfRange)
	}

	hasFlag, err := HasFlagUint8(flags, flag)
	if err != nil {
		return flags, err
	}

	if hasFlag == set {
		return flags, nil
	}

	conv := uint8(1 << flag)
	ret := flags ^ conv
	return ret, nil
}
