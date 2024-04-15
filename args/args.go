package args

import (
	"bufio"
	"encoding/json"
	"file-manager-export/exit"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Args struct {
	DbPath    string `json:"dbPath"`
	OutputDir string `json:"outputDir"`
	CpuCount  int    `json:"cpuCount"`
}

type argConstruct func(a *Args)

var defaultArgs Args = Args{}

var fileFlag = flag.String("f", "", "通过json配置文件获取命令行参数")
var dbPathFlag = flag.String("d", "", "sqlite数据库文件存放地址，优先级低于-f")
var outputDirFlag = flag.String("o", "", "文件导出的目录位置，优先级低于-f")
var cpuCountFlag = flag.Int("c", 0, "启用的CPU核心数")

func New(opts ...argConstruct) Args {
	args := Args{}

	for _, arg := range opts {
		arg(&args)
	}

	// 参数缺少DbPath
	if args.DbPath == "" {
		// 从命令行参数获取值
		if defaultArgs.DbPath != "" {
			args.DbPath = defaultArgs.DbPath
		} else {
			// 通过命令行让用户输入
			args.DbPath = scanString("请输入sqlite数据库文件的路径：")
		}
	}

	// 参数缺少OutputDir
	if args.OutputDir == "" {
		// 从命令行参数获取值
		if defaultArgs.OutputDir != "" {
			args.OutputDir = defaultArgs.OutputDir
		} else {
			// 通过命令行让用户输入
			args.OutputDir = scanString("请输入文件导出的目录：")
		}
	}

	// 参数缺少cpuCount
	if args.CpuCount == 0 {
		// 从命令行参数获取值
		if defaultArgs.CpuCount > 0 {
			args.CpuCount = defaultArgs.CpuCount
		} else {
			// 通过命令行让用户输入
			fmt.Println("请输入启用CPU核心数：")
			fmt.Scanln(&args.CpuCount)
			if args.CpuCount < 1 {
				exit.Error("请输入正确的CPU核心数")
			}
		}
	}

	return args
}

func WithDbPath(dbPath string) argConstruct {
	return func(a *Args) {
		a.DbPath = dbPath
	}
}

func WithOutputDir(outputDir string) argConstruct {
	return func(a *Args) {
		a.OutputDir = outputDir
	}
}

func WithCpuCount(cpuCount int) argConstruct {
	return func(a *Args) {
		a.CpuCount = cpuCount
	}
}

func scanString(question string) string {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println(question)

	input, err := inputReader.ReadString('\n')
	if err != nil {
		exit.Error(err.Error())
	}
	return strings.Replace(strings.Replace(input, "\r", "", -1), "\n", "", -1)
}

func init() {
	// flag.PrintDefaults()
	flag.Parse()
	if *fileFlag != "" {
		buf, err := os.ReadFile(*fileFlag)
		if err != nil {
			exit.Error(*fileFlag + "文件不存在或无法打开")
		}

		err = json.Unmarshal(buf, &defaultArgs)

		if err != nil {
			exit.Error(*fileFlag + "文件的数据不是json格式")
		}
	}

	if defaultArgs.DbPath == "" && *dbPathFlag != "" {
		defaultArgs.DbPath = *dbPathFlag
	}

	if defaultArgs.OutputDir == "" && *outputDirFlag != "" {
		defaultArgs.OutputDir = *outputDirFlag
	}

	if defaultArgs.CpuCount == 0 && *cpuCountFlag > 0 {
		defaultArgs.CpuCount = *cpuCountFlag
	}
}
