package repos

import (
	"encoding/json"
	"fmt"

	"github.com/wangle201210/githubapi"
)

const tagUrl = "tags"

type commit struct {
	Sha string `json:"sha"`
	Url string `json:"url"`
}

type tag struct {
	Name       string `json:"name"`
	ZipballUrl string `json:"zipball_url"`
	TarballUrl string `json:"tarball_url"`
	Commit     commit `json:"commit"`
}

type tagsList []tag

func (re *Repos) GetTagsList() (list tagsList) {
	url := re.Url(tagUrl)
	body, err := githubapi.HttpGet(url)
	if err != nil {
		fmt.Printf("Get info from github error: %s", err)
		return
	}
	if err = json.Unmarshal(body, &list); err != nil {
		fmt.Printf("Unmarshal body error: %s", err)
		return
	}
	return
}

func (re *Repos) LastTag() tag {
	list := re.GetTagsList()
	if len(list) < 1 {
		fmt.Print("There is no tag")
	}
	return list[0]
}
