package main

import (
	"context"
	"fmt"
	lnk "github.com/parsiya/golnk"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"os/exec"
	"os/user"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Test() {
	fmt.Println("打开应用。。。")
	cmd := exec.Command("D:\\QQ\\Bin\\QQScLauncher.exe")
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	return
}

type LinkFile struct {
	name       string
	targetPath string
}

func (a *App) Search() []string {
	u, err := user.Current()
	fmt.Println("当前用户")
	fmt.Printf("HomeDir %s\n", u.HomeDir)
	fmt.Println("寻找桌面所有快捷方式。。。")
	linkFiles := make([]LinkFile, 0)
	arr := make([]string, 0)
	files, err := ioutil.ReadDir(u.HomeDir + "\\Desktop")
	if err != nil {
		fmt.Print(err)
	}
	for _, file := range files {
		//fmt.Println(file.Name()[len(file.Name())-3 : len(file.Name())])
		if len(file.Name()) > 3 && file.Name()[len(file.Name())-3:len(file.Name())] == "lnk" {
			var linkFile LinkFile
			linkFile.name = file.Name()[:len(file.Name())-4]
			Lnk, err := lnk.File(u.HomeDir + "\\Desktop\\" + file.Name())
			if err != nil {
				panic(err)
			}
			// 中文路径需要解码，英文路径可忽略
			targetPath, _ := simplifiedchinese.GBK.NewDecoder().String(Lnk.LinkInfo.LocalBasePath)
			linkFile.targetPath = targetPath
			linkFiles = append(linkFiles, linkFile)
			arr = append(arr, linkFile.name)
		}
	}
	fmt.Println(linkFiles)
	//file = linkFiles
	//fmt.Println(linkFiles)
	return arr
}

func (a *App) Open(name string) {
	u, err := user.Current()
	fmt.Println("当前用户")
	fmt.Printf("HomeDir %s\n", u.HomeDir)
	fmt.Println("寻找桌面所有快捷方式。。。")
	files, err := ioutil.ReadDir(u.HomeDir + "\\Desktop")
	if err != nil {
		fmt.Print(err)
	}
	for _, file := range files {
		//fmt.Println(file.Name()[len(file.Name())-3 : len(file.Name())])
		if len(file.Name()) > 3 && file.Name()[len(file.Name())-3:len(file.Name())] == "lnk" {
			if name == file.Name()[:len(file.Name())-4] {
				Lnk, err := lnk.File(u.HomeDir + "\\Desktop\\" + file.Name())
				if err != nil {
					panic(err)
				}
				// 中文路径需要解码，英文路径可忽略
				targetPath, _ := simplifiedchinese.GBK.NewDecoder().String(Lnk.LinkInfo.LocalBasePath)
				cmd := exec.Command(targetPath)
				_, err = cmd.CombinedOutput()
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
			}
		}
	}
	fmt.Println("打开应用。。。")
	return
}

func (a *App) Test2() []string {
	//arr := make([]LinkFile, 0)
	//arr = append(arr, LinkFile{name: "Hello", targetPath: "hello"})
	arr := make([]string, 0)
	arr = append(arr, "1111111")
	return arr
}
