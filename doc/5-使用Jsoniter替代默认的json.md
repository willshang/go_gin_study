# 使用 jsoniter 编译
- Gin 使用 encoding/json 作为默认的 json 包，但是你可以在编译中使用标签将其修改为 jsoniter。
```sh
$ go build -tags=jsoniter .
```
- 相关链接
- https://github.com/json-iterator/go
---
## Build with [jsoniter](https://github.com/json-iterator/go)

Gin uses `encoding/json` as default json package but you can change to [jsoniter](https://github.com/json-iterator/go) by build from other tags.

```sh
$ go build -tags=jsoniter .
```