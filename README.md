# file-manager-export
本项目实现了[file-manager](https://github.com/shenX-2021/file-manager)的文件导出功能，也是对GO的代码实践。

## 使用说明
### 步骤
1. 根据[file-manager](https://github.com/shenX-2021/file-manager)的文档安装（需使用本地启动的方式）
2. 克隆代码: `git clone git@github.com:shenX-2021/file-manager-export.git`
3. 安装依赖: `cd file-manager-export && go generate`
4. 启动: `go run main.go`
5. 根据提示填写信息即可

### 配置文件
提供了`config.example.json`作为示例的配置文件。如需使用，可更改完配置文件后，执行：
```bash
go run main.go -f config.example.json
```

## 编译
由于有交叉编译的需求，因此提供了编译Windows平台的命令：

```bash
make build-windows
```

关于交叉编译的细节，[请看此文](https://blog.csdn.net/weixin_41855143/article/details/137794986)。