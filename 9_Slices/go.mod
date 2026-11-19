
// Go 标准库里有一个官方包 slices(Go 1.21+)。
// macOS 的文件系统默认不区分大小写，所以 Slices 和 slices 被视为同一路径，Go 工具链无法判断你指的是哪个，就报了 ambiguous import（导入歧义）警告。

// module Slices  



module Slices_test
go 1.27.1
