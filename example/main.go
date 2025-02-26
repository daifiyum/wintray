package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/daifiyum/wintray"
	W "github.com/daifiyum/wintray/windows"
)

func init() {
	// 设置AUMID，气泡通知左上角的图标和名称基于此
	// 获取图片绝对路径
	iconURL, _ := filepath.Abs("./p1.ico")
	// 在注册表中注册AUMID
	W.RegisterAUMID("wintray", "wintray", iconURL)
	// 将当前进程绑定到上面注册的AUMID上
	W.SetAUMID("wintray")
	// W.UnregisterAUMID("wintray") // 可以删除注册的AUMID
}

func main() {
	app := wintray.New("wintray", "./p1.ico")

	// 菜单
	// 目前菜单项支持：普通菜单项、复选菜单项、子菜单、分隔符
	// 菜单项图标暂不支持，我用不到，后续可能会支持，也可以自己扩展，只需要修改windows/menu.go文件即可

	// 子菜单
	subMenu := W.NewMenu()
	subMenu.AddItem(4, "子菜单项", func() {
		fmt.Println("子菜单项")
	})

	// 主菜单
	menu := W.NewMenu()
	// 添加菜单项，参数分别是：ID、文本、回调函数，ID必须是唯一的，后面操作菜单项都是通过ID来操作
	menu.AddItem(1, "菜单项", func() {
		fmt.Println("菜单项")
	})
	menu.AddCheckMenu(2, "复选菜单项", true, func() {
		if menu.ToggleCheck(2) {
			fmt.Println("选中")
		} else {
			fmt.Println("未选中")
		}
	})
	menu.AddSubMenu("子菜单", subMenu)
	menu.AddSeparator()
	menu.AddItem(3, "退出", func() {
		fmt.Println("退出")
		app.Quit()
	})

	app.SetMenu(menu)

	app.SetOnLeftClick(func() {
		fmt.Println("左键点击")
		app.ShowTrayNotification("通知", "这是一条通知")
	})

	// 这个函数会在托盘运行前调用，可以在这里做一些初始化操作
	// 例如：基于这个托盘实现的REST API服务，就可以在这里面启动API服务
	// 当然，只是为了规范，你也可以不调用这个函数
	app.SetOnInitialize(func() {
		fmt.Println("准备就绪")
	})

	go func() {
		time.Sleep(5 * time.Second)
		// 5秒后重新设置提示和托盘图标
		app.SetToolTip("新提示")
		app.SetIcon("./p2.ico")
	}()

	app.Run()
}
