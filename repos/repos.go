package repos

const (
	baseUrl = "https://api.github.com/"
)

type Repos struct {
	PkgName string
}

func (re *Repos) Url(name string) (url string) {
	return baseUrl + "repos/" + re.PkgName + "/" + name
}
