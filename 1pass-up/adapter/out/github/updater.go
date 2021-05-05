// Implementation of GitHub updater.
//
// @author TSS

package github

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
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

func (up *githubUpdater) DownloadFile(destination, url string) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	file, err := os.Create(destination)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, resp.Body)

	if err != nil {
		return err
	}

	return nil
}

func (up *githubUpdater) ExtractArchive(src, dst string) error {
	archive, err := os.Open(src)

	if err != nil {
		return err
	}

	gz, err := gzip.NewReader(archive)

	if err != nil {
		return err
	}

	defer gz.Close()
	tr := tar.NewReader(gz)

	for {
		header, err := tr.Next()

		switch {
		case err == io.EOF:
			return nil

		case err != nil:
			return err
		}

		target := filepath.Join(dst, header.Name)

		if header.Typeflag == tar.TypeReg {
			file, err := os.Create(target)

			if err != nil {
				return err
			}

			defer file.Close()

			if _, err := io.Copy(file, tr); err != nil {
				return err
			}
		}
	}
}
