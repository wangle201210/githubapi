# GitHubAPI

[GitHub API v3](http://developer.github.com/v3/)

```
import "github.com/wangle201210/githubapi/repos"
...
var pkg = repos.Pkg{"beego", "bee"}

// 获取仓库信息
re, err := pkg.GetRepos()

// 获取tags
list, err := pkg.GetTagsList()

// 获取最新tag
tag, err := pkg.LastTag()

// 获取pull
list, err := pkg.GetPullsList()

// 获取comment
comment, err := repos.GetCommentsList()

// 获取全部issues
issues, err := repos.GetIssuesList()

...
```
