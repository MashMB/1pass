// Implementation of GitHub updater.
//
// @author TSS

package github

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/mashmb/1pass/1pass-core/core/domain"
)

type githubUpdater struct {
}

func NewGithubUpdater() *githubUpdater {
	return &githubUpdater{}
}

func (up *githubUpdater) CheckForUpdate() (*domain.UpdateInfo, error) {
	resp, err := http.Get(domain.GithubReleases)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var bodyJson []interface{}

	if err := json.Unmarshal(body, &bodyJson); err != nil {
		return nil, err
	}

	if bodyJson == nil || bodyJson[0] == nil {
		return nil, nil
	}

	latestJson := bodyJson[0].(map[string]interface{})
	version := latestJson["tag_name"].(string)
	newer := false
	var archiveUrl string
	var checksumUrl string
	assets := latestJson["assets"].([]interface{})

	if assets == nil {
		return nil, nil
	}

	for _, assetJson := range assets {
		asset := assetJson.(map[string]interface{})

		if strings.Contains(asset["name"].(string), ".tar.gz") {
			archiveUrl = asset["browser_download_url"].(string)
		} else if strings.Contains(asset["name"].(string), ".md5") {
			checksumUrl = asset["browser_download_url"].(string)
		}
	}

	currVer, _ := strconv.ParseInt(strings.ReplaceAll(domain.Version, ".", ""), 10, 64)
	remoteVer, _ := strconv.ParseInt(strings.ReplaceAll(version, ".", ""), 10, 64)

	if remoteVer > currVer {
		newer = true
	}

	return domain.NewUpdateInfo(archiveUrl, checksumUrl, version, newer), nil
}
