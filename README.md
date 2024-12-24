## wintray

[EN](./README_EN.md)

go/win32 实现的最简托盘，包含托盘菜单和气泡通知

|                       托盘                       |                       通知                       |
| :----------------------------------------------: | :----------------------------------------------: |
| <img src="https://thumbsnap.com/i/Rpw3yBbs.png"> | <img src="https://thumbsnap.com/i/mxjos6b2.png"> |

### 起源

现有 go/win 托盘年久失修，还缺少气泡通知，便重新实现

### 注意（必看）

**AUMID：**

![](https://thumbsnap.com/i/35oeEvry.png)

图中红色边框里面的是应用图标和名称，这个可以通过AUMID来定义：

```
func init() {
	// 获取图片绝对路径
	iconURL, _ := filepath.Abs("./p1.ico")
	// 在注册表中注册AUMID
	W.RegisterAUMID("wintray", "wintray", iconURL)
	// 将当前进程绑定到上面注册的AUMID上
	W.SetAUMID("wintray")
}

// RegisterAUMID("aumid", "应用名称", "应用图标路径")
// 参数1: 字符串类型，可以随便定义
// 参数2: 字符串类型，应用名称，可以随便定义
// 参数3: 字符串类型, 应用图标路径，必须是绝对路径
```

**.rc 文件：**

AUMID需要写入注册表，如果不想通过AUMID来定义气泡通知的应用名称，则可以使用.rc文件内`FileDescription`的值来定义应用名称，为空显示`文件名.exe`

需要注意的是，.rc文件只能定义气泡通知的应用名称，应用图标则来自`wintray.New("wintray", "./p1.ico")`函数的第二个参数，他也是托盘图标

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
