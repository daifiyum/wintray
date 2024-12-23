package windows

import (
	"syscall"
	"unicode/utf16"
	"unsafe"
)

// 加载本地 .ico
func LoadIconFromFile(iconPath string) (syscall.Handle, error) {
	iconPathPtr, _ := syscall.UTF16PtrFromString(iconPath)

	ret, _, err := LoadImage.Call(
		0,
		uintptr(unsafe.Pointer(iconPathPtr)),
		uintptr(IMAGE_ICON),
		0,
		0,
		uintptr(LR_LOADFROMFILE|LR_DEFAULTSIZE),
	)

	if ret == 0 {
		return 0, err
	}

	return syscall.Handle(ret), nil
}

// LOWORD
func LOWORD(l uint64) uint32 {
	return uint32(l & 0xFFFFFFFF)
}

// HIWORD
func HIWORD(l uint64) uint32 {
	return uint32((l >> 32) & 0xFFFFFFFF)
}

// 字符串处理
func SetUTF16String(dst interface{}, src string) {
	utf16Slice := utf16.Encode([]rune(src))
	switch d := dst.(type) {
	case *[64]uint16:
		copy(d[:], utf16Slice)
	case *[256]uint16:
		copy(d[:], utf16Slice)
	default:
		panic("unsupported array type")
	}
}

// 托盘提示语
func TipFromStr(s string) [128]uint16 {
	utf16Tip, _ := syscall.UTF16FromString(s)
	var szTip [128]uint16
	copy(szTip[:], utf16Tip)
	return szTip
}
