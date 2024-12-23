## wintray

A minimal system tray implementation using Go/Win32, including tray menus and balloon notifications.

### Origin

Existing Go/Win tray programs are outdated and lack balloon notifications, so I decided to reimplement it.

### Important Notes (Must Read)

**Balloon Notifications:**

Since balloon notifications are implemented using Win32, they require the use of a `.rc` file. The reasons are as follows:

| Feature                             | Traditional Notifications (Shell_NotifyIcon)                                                                               | Modern Notifications (ToastNotification)    |
| ----------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------- |
| Implementation                      | Uses `Shell_NotifyIcon` Win32 API                                                                                          | Uses `ToastNotification` WinRT API          |
| Icons (Window, Notifications, etc.) | Defined when registering the window                                                                                        | Retrieved via AUMID defined in the registry |
| App Name                            | Retrieved from `FileDescription` in the `.rc` file, or uses the file name with extension (e.g., `main.exe`) if not defined | Retrieved via AUMID defined in the registry |

**.rc File:**

`FileDescription`: Defines the application name in the top-left corner of balloon notifications. If empty, it will display the `filename.exe`.

`IDI_MAIN ICON`: Defines the icon for `filename.exe`.

For other definitions, please refer to relevant documentation.

**Compilation:**

Since the `.rc` file is required, it must be included during compilation:

```
go build -o main.exe ./

// Hide the command-line window
go build -ldflags="-H=windowsgui" -o main.exe ./

// Incorrect method, this will not include the .rc file
go build -o main.exe ./main.go
```

**Non-blocking Execution:**

```
// This format is mandatory; other methods may cause issues. Feel free to test it yourself.
go func() {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		app := wintray.New("wintray", "./p1.ico")
		app.Run()
}()
```

### Usage

See `example/main.go`.
