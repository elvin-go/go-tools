## 这个是go开发帮助工具包
## 使用go clean -modcache 来清理依赖缓存
## 使用go get github.com/elvin-go/go-tools@dev,拉到分支最新代码

### 设置无代理
go env -w GOPROXY=direct
### 设置不做检查
go env -w GOSUMDB=off
### 清理缓存
go clean -modcache
go clean -cache
### 强制更新
go get -x -u github.com/elvin-go/go-tools@dev