# 标准输出处理包说明

## 问题回答

**问题**: 定位一下哪个包是负责处理标准输出的

**答案**: **`pkg/prog` 包负责处理标准输出 (stdout)**

## 详细说明

### pkg/prog 包的职责

`pkg/prog` 包是 tdl 应用程序中处理标准输出的主要包。它的核心功能包括：

1. **创建进度条写入器**：创建并配置 `progress.Writer` 实例，默认输出到 `os.Stdout`
2. **显示进度条**：在终端显示下载/上传的进度条和统计信息
3. **系统指标显示**：显示 CPU、内存和 goroutine 使用情况
4. **样式格式化**：为所有进度显示应用一致的颜色、宽度和格式

### 技术实现

```go
// pkg/prog/prog.go
func New(formatter progress.UnitsFormatter) progress.Writer {
    pw := progress.NewWriter()  // 内部默认输出到 os.Stdout
    // ... 配置样式和格式
    return pw
}
```

`progress.NewWriter()` 函数来自 `github.com/jedib0t/go-pretty/v6/progress` 库，该函数创建的写入器默认输出到 `os.Stdout`。

### 标准输出流程

```
应用程序操作 (下载/上传/导出)
    ↓
prog.New() 创建 progress.Writer
    ↓
progress.NewWriter() 内部设置输出到 os.Stdout
    ↓
pw.Render() 启动后台协程写入 stdout
    ↓
在终端显示进度条和指标
```

### 使用位置

该包在整个应用程序中被广泛使用：

- **app/dl/dl.go**: 下载进度跟踪
- **app/up/up.go**: 上传进度跟踪
- **app/chat/export.go**: 消息导出进度
- **app/chat/users.go**: 用户枚举进度
- **app/forward/forward.go**: 消息转发进度

### 其他标准输出使用

虽然 `pkg/prog` 是结构化进度输出到 stdout 的主要包，但代码库的其他部分也使用 `fmt.Print*` 函数写入 stdout，用于：

- 简单的状态消息
- JSON 输出（例如 `app/chat/ls.go`、`app/chat/export.go`）
- 表格显示（例如 `app/extension/extension.go`）
- 二维码显示（例如 `app/login/qr.go`）

但是，对于所有**进度跟踪和长时间运行操作的可视化**，`pkg/prog` 是指定的处理程序。

## 参考文档

更多详细信息，请参阅：
- `pkg/prog/README.md` - 完整的英文文档
- `pkg/prog/prog.go` - 包级别文档和函数说明
- `pkg/prog/tracker.go` - 跟踪器和系统指标文档
