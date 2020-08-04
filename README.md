# GitHubAPI

[GitHub API v3](http://developer.github.com/v3/)

```
import "github.com/wangle201210/githubapi/repos"
...
var pkg = repos.Pkg{"beego", "bee"}
// 获取仓库信息
repos, err := pkg.GetRepos()
// 获取这个仓库的所有tags
list, err := pkg.GetTagsList()
// 获取这个仓库的所有pull
list, err := pkg.GetPullList()
// 获取最新tag
tag, err := pkg.LastTag()
...
```
