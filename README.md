# go-utils

`go-utils` 是一个实用的 Go 语言工具包，提供了一些常用的功能，帮助开发者更高效地编写代码。

## 安装

使用以下命令安装 `go-utils`：

```sh
go get github.com/MinRadish/go-utils
```

## 导入

在你的 Go 代码中导入 `go-utils` 包：

```go
import "github.com/MinRadish/go-utils"
```

## 使用示例

以下是一些使用 `go-utils` 包的示例代码：

### 示例 1: 字符串工具

```go
package main

import (
    "fmt"
    "github.com/MinRadish/go-utils/strutil"
)

func main() {
    str := "Hello, World!"
    reversed := strutil.Reverse(str)
    fmt.Println(reversed) // 输出: !dlroW ,olleH
}
```

### 示例 2: 数学工具

```go
package main

import (
    "fmt"
    "github.com/MinRadish/go-utils/mathutil"
)

func main() {
    sum := mathutil.Add(1, 2, 3, 4, 5)
    fmt.Println(sum) // 输出: 15
}
```

## 文档

更多详细信息和使用文档，请访问 [官方文档](https://github.com/MinRadish/go-utils)。

## 贡献

欢迎贡献代码！请阅读 [贡献指南](https://github.com/MinRadish/go-utils/blob/main/CONTRIBUTING.md) 了解更多信息。

## 许可证

`go-utils` 使用 MIT 许可证。详情请参阅 [LICENSE](https://github.com/MinRadish/go-utils/blob/main/LICENSE) 文件。