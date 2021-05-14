// Implementation of GitHub updater.
//
// @author TSS

package github

import (
	"archive/tar"
	"compress/gzip"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/mashmb/1pass/1pass-core/core/domain"
)

type githubUpdater struct{}

func NewGithubUpdater() *githubUpdater {
	return &githubUpdater{}
}

func httpClient(timeout int64) http.Client {
	httpTransport := http.Transport{
		Dial: func(network, address string) (net.Conn, error) {
			tout := time.Duration(timeout) * time.Second

			return net.DialTimeout(network, address, tout)
		},
	}

	httpClient := http.Client{
		Transport: &httpTransport,
	}

	return httpClient
}

func (up *githubUpdater) CheckForUpdate(timeout int64) (*domain.UpdateInfo, error) {
	httpClient := httpClient(timeout)
	resp, err := httpClient.Get(domain.GithubReleases)

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
	changelog := latestJson["body"].(string)
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

	return domain.NewUpdateInfo(archiveUrl, checksumUrl, changelog, version, newer), nil
}

func (up *githubUpdater) CheckTimestamp(dirPath string) {
	timestamp := time.Now().Unix()
	file := filepath.Join(dirPath, domain.LastCheckFile)

	if _, err := os.Stat(dirPath); err != nil {
		os.MkdirAll(dirPath, 0700)
	}

	ioutil.WriteFile(file, []byte(fmt.Sprint(timestamp)), 0644)
}

func (up *githubUpdater) DownloadFile(destination, url string, timeout int64) error {
	httpClient := httpClient(timeout)
	resp, err := httpClient.Get(url)

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

func (up *githubUpdater) ReplaceBinary(src string) error {
	dst, err := os.Executable()

	if err != nil {
		return err
	}

	if err := os.Rename(src, dst); err != nil {
		return err
	}

	return nil
}

func (up *githubUpdater) ShouldCheck(period int, dirPath string) bool {
	should := false
	lastCheckFile := filepath.Join(dirPath, domain.LastCheckFile)

	if _, err := os.Stat(lastCheckFile); err != nil {
		should = true
	}

	file, err := ioutil.ReadFile(lastCheckFile)

	if err != nil {
		should = true
	}

	savedVal, err := strconv.ParseInt(strings.TrimSpace(string(file)), 10, 64)

	if err != nil {
		should = true
	}

	timestamp := time.Unix(savedVal, 0)
	now := time.Now()
	between := math.Abs(now.Sub(timestamp).Hours() / 24)

	if int(between) >= period {
		should = true
	}

	return should
}

func (up *githubUpdater) ValidateChecksum(binary, file string) error {
	fileContent, err := ioutil.ReadFile(file)

	if err != nil {
		return err
	}

	chunks := strings.Fields(strings.TrimSpace(string(fileContent)))
	binaryFile, err := os.Open(binary)

	if err != nil {
		return err
	}

	defer binaryFile.Close()
	hash := md5.New()

	if _, err := io.Copy(hash, binaryFile); err != nil {
		return err
	}

	bytes := hash.Sum(nil)
	calculated := hex.EncodeToString(bytes)

	if chunks[0] != calculated {
		return domain.ErrInvalidMd5
	}

	return nil
}
