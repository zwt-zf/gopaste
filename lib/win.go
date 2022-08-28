package lib

import (
	"fmt"
	"syscall"
)

var (
	// golang call windows api : https://github.com/golang/go/wiki/WindowsDLLs
	// windows api
	// https://docs.microsoft.com/en-us/windows/win32/dataxchg/clipboard-reference
	user32 = syscall.MustLoadDLL("user32.dll")

	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-openclipboard
	openClipboard = user32.MustFindProc("OpenClipboard")

	closeClipboard   = user32.MustFindProc("CloseClipboard")
	emptyClipboard   = user32.MustFindProc("EmptyClipboard")
	getClipboardData = user32.MustFindProc("GetClipboardData")
	setClipboardData = user32.MustFindProc("SetClipboardData")

	addClipboardFormatListener = user32.MustFindProc("AddClipboardFormatListener")
)

func RegisterListener(hwnd int) {
	handle := uintptr(user32.Handle)
	fmt.Println(handle)
	r1, r2, _ := addClipboardFormatListener.Call(handle)
	fmt.Println(r1)
	fmt.Println(r2)
}

func GetClipboardData() {
	getClipboardData.Call()
}
