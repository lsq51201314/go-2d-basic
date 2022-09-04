package assets

import (
	"fmt"
	"io/fs"
	"os"
)

// 获取文件
func getFiles(folder string) (files []string, err error) {
	var fs []fs.DirEntry
	if fs, err = os.ReadDir(folder); err != nil {
		return
	}
	for _, fi := range fs {
		path := folder + "/" + fi.Name()
		if !fi.IsDir() && path[len(path)-4:] == ".gpk" {
			files = append(files, path)
		}
	}
	return
}

// 获取目录
func getDirs(folder string) (dirs []string, err error) {
	var fs []fs.DirEntry
	if fs, err = os.ReadDir(folder); err != nil {
		return
	}
	for _, fi := range fs {
		path := folder + "/" + fi.Name()
		if fi.IsDir() {
			dirs = append(dirs, path)
		}
	}
	return
}

// 回调
func process(file string, current, count int) {
	fmt.Println(file, current, count)
}
