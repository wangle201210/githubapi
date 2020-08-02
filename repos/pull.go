package repos

import (
	"encoding/json"
	"fmt"

	"github.com/wangle201210/githubapi"
)

const pullUrl = "pulls"

type pull struct {
	Url      string `json:"url"`
	Id       int64  `json:"id"`
	NodeId   string `json:"node_id"`
	HtmlUrl  string `json:"html_url"`
	DiffUrl  string `json:"diff_url"`
	PatchUrl string `json:"patch_url"`
	IssueUrl string `json:"issue_url"`
	Number   int64  `json:"number"`
	State    string `json:"state"`
	Locked   bool   `json:"locked"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	//User	string `json:"user"` // 以后加上结构体后再补充user
}

type pullList []pull

func (re *Repos) GetPullList() (list pullList) {
	url := re.Url(pullUrl)
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
