## wintray

[EN](./README_EN.md)

go/win32 实现的最简托盘，包含托盘菜单和气泡通知

|                       托盘                       |                       通知                       |
| :----------------------------------------------: | :----------------------------------------------: |
| <img src="https://thumbsnap.com/i/Rpw3yBbs.png"> | <img src="https://thumbsnap.com/i/mxjos6b2.png"> |

### 起源

现有 go/win 托盘年久失修，还缺少气泡通知，便重新实现

### 注意（必看）

**气泡通知：**

由于使用 win32 实现的气泡通知，使用时需要结合.rc 文件，具体原因如下：

| 特性             | 传统通知 (Shell_NotifyIcon)                                                     | 现代通知 (ToastNotification)       |
| ---------------- | ------------------------------------------------------------------------------- | ---------------------------------- |
| 实现方式         | 使用 `Shell_NotifyIcon` Win32 API                                               | 使用 `ToastNotification` WinRT API |
| 窗口、通知等图标 | 注册窗口时定义                                                                  | 通过注册表内定义的 AUMID 获取      |
| APP 名称         | 从 `.rc` 文件的 `FileDescription` 获取，若无则使用文件名加后缀（如 `main.exe`） | 通过注册表内定义的 AUMID 获取      |

**.rc 文件：**

`FileDescription`：定义气泡通知左上角应用名称，为空则显示`文件名.exe`

`IDI_MAIN ICON`：定义`文件名.exe`的图标

其他定义请自行查找

**编译：**

由于有.rc 文件，编译时需要包含

```
go build -o main.exe ./

// 隐藏命令行窗口
go build -ldflags="-H=windowsgui" -o main.exe ./

// 错误方式，这种不会包含.rc文件
go build -o main.exe ./main.go
```

**非阻塞运行：**

```
// 必须这种格式，其他方式会出问题，不信自测
go func() {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		app := wintray.New("wintray", "./p1.ico")
		app.Run()
}()
```

### 使用

见 example/main.go
