package windows

import (
	"fmt"
	"syscall"
	"unicode/utf16"
	"unsafe"

	"golang.org/x/sys/windows/registry"
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

// 注册AUMID
func RegisterAUMID(aumid, displayName, iconURI string) error {
	keyPath := `Software\Classes\AppUserModelId\` + aumid

	key, _, err := registry.CreateKey(registry.CURRENT_USER, keyPath, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()

	if err := key.SetStringValue("DisplayName", displayName); err != nil {
		return err
	}

	if err := key.SetStringValue("IconUri", iconURI); err != nil {
		return err
	}

	return nil
}

// 删除已注册的AUMID
func UnregisterAUMID(aumid string) error {
	keyPath := `Software\Classes\AppUserModelId\` + aumid

	err := registry.DeleteKey(registry.CURRENT_USER, keyPath)
	if err != nil {
		return err
	}

	return nil
}

// 绑定AUMID
// 最低受支持的客户端: Windows 7 [仅限桌面应用]
// 最低受支持的服务器: Windows Server 2008 R2 [仅限桌面应用]
func SetAUMID(aumid string) error {
	aumidPtr, _ := syscall.UTF16PtrFromString(aumid)
	r1, _, err := SetCurrentProcessExplicitAppUserModelID.Call(uintptr(unsafe.Pointer(aumidPtr)))
	if r1 != 0 {
		return fmt.Errorf("SetCurrentProcessExplicitAppUserModelID failed: %v", err)
	}
	return nil
}
