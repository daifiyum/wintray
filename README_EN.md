## wintray

A minimal system tray implementation using Go/Win32, including tray menus and balloon notifications.

|                       tray                       |                      toast                       |
| :----------------------------------------------: | :----------------------------------------------: |
| <img src="https://thumbsnap.com/i/Rpw3yBbs.png"> | <img src="https://thumbsnap.com/i/mxjos6b2.png"> |

### Origin

Existing Go/Win tray programs are outdated and lack balloon notifications, so I decided to reimplement it.

### Important Notes (Must Read)

**AUMID:**

![](https://thumbsnap.com/i/35oeEvry.png)

The application icon and name inside the red border in the image can be defined using AUMID:

```go
func init() {
	// Get the absolute path of the image
	iconURL, _ := filepath.Abs("./p1.ico")
	// Register AUMID in the registry
	W.RegisterAUMID("wintray", "wintray", iconURL)
	// Bind the current process to the registered AUMID
	W.SetAUMID("wintray")
}

// RegisterAUMID("aumid", "Application Name", "Application Icon Path")
// Parameter 1: String type, can be defined arbitrarily
// Parameter 2: String type, application name, can be defined arbitrarily
// Parameter 3: String type, application icon path, must be an absolute path
```

**.rc File:**

AUMID requires writing to the registry. If you do not want to define the application name for the toast notification using AUMID, you can use the value of `FileDescription` in the .rc file to define the application name. If left empty, it will display `filename.exe`.

It is important to note that the .rc file can only define the application name for the toast notification. The application icon comes from the second parameter of the `wintray.New("wintray", "./p1.ico")` function, which is also the tray icon.

Since the .rc file is involved, it needs to be included during compilation:

```sh
go build -o main.exe ./

# Hide the command line window
go build -ldflags="-H=windowsgui" -o main.exe ./

# Incorrect method, this will not include the .rc file
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
