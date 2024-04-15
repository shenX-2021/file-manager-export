package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"

	"file-manager-export/args"
	"file-manager-export/exit"
	"file-manager-export/model"
)

type Result struct {
	len     int
	success int
	failure int
}

func (r *Result) String() string {
	return fmt.Sprintf(`{
		"len": %d,
		"success": %d,
		"failure": %d
	}`, r.len, r.success, r.failure)
}

func main() {
	a := args.New()
	r := Result{}
	// 设置CPU核心数
	runtime.GOMAXPROCS(a.CpuCount)

	// 初始化数据库
	model.InitDb(a)

	// 判断目录是否存在
	dir, err := os.Stat(a.OutputDir)
	if err != nil {
		fmt.Println("检测目录是否存在失败：", err)
		os.Exit(1)
	}
	if !dir.IsDir() {
		exit.Error(a.OutputDir + "不是一个目录")
	}

	// 查找文件列表
	var fileList []model.FileModel
	result := model.DB.Find(&fileList)
	if result.Error != nil {
		exit.Error(result.Error.Error())
	}

	// 处理导出
	ch := make(chan bool)
	for _, fileData := range fileList {
		go exportHandler(ch, fileData, a)
	}

	// 对处理结果进行汇总统计
	for {
		if <-ch {
			r.success += 1
		} else {
			r.failure += 1
		}
		r.len++

		if r.len == len(fileList) {
			break
		}
	}

	fmt.Printf("导出文件完成，导出文件总数为 %d，导出成功的数量为 %d，导出失败的数量为%d\n", r.len, r.success, r.failure)
}

// 导出文件，协程函数
func exportHandler(ch chan bool, fileData model.FileModel, args args.Args) {
	err := export(fileData, args)
	if err != nil {
		fmt.Printf("%s文件导出时发生错误： %v\n", fileData.FileName, err)
		ch <- false
		return
	}

	ch <- true
}

// 导出文件具体处理
func export(fileData model.FileModel, args args.Args) error {
	file, err := os.Open(fileData.FilePath)
	if err != nil {
		exit.Error("要导出的文件" + fileData.FileName + "不存在")
	}
	defer file.Close()

	dstFilePath := path.Join(args.OutputDir, fileData.FileName)
	// 检测存放目录是否存在同名文件
	_, err = os.Stat(dstFilePath)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			exit.Error(err.Error())
		}
	} else {
		return errors.New("文件已存在")
	}

	outputFile, err := os.OpenFile(dstFilePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		exit.Error(err.Error())
	}
	defer outputFile.Close()

	const NBUF = 512

	var buf [NBUF]byte
	for {
		switch nr, err := file.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0: // EOF
			return nil
		case nr > 0:
			outputFile.Write(buf[:])
		}
	}

}
