// Unit tests for GitHub updater.
//
// @author TSS

package github

import (
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/mashmb/1pass/1pass-core/core/domain"
)

func setupGithubUpdater() *githubUpdater {
	return NewGithubUpdater()
}

func isOnline() bool {
	_, err := http.Get("http://google.com")

	if err != nil {
		return false
	}

	return true
}

func TestCheckForUpdate(t *testing.T) {
	if isOnline() {
		updater := setupGithubUpdater()
		expected := false
		timeout := int64(5)
		info, err := updater.CheckForUpdate(timeout)

		if err != nil {
			t.Error("CheckForUpdate() should pass")
		}

		if info == nil {
			t.Error("CheckForUpdate() should provide any data")
		}

		if info.Newer != expected {
			t.Errorf("CheckForUpdate() = %v; expected = %v", info.Newer, expected)
		}
	} else {
		t.Log("CheckForUpdate() no internet connection")
	}
}

func TestCheckTimestamp(t *testing.T) {
	updater := setupGithubUpdater()
	dirPath := "../../../assets"
	file := filepath.Join(dirPath, domain.LastCheckFile)
	updater.CheckTimestamp(dirPath)

	if _, err := os.Stat(file); err != nil {
		t.Error("CheckTimestamp() should create or update the file")
	}
}

func TestDownloadFile(t *testing.T) {
	if isOnline() {
		updater := setupGithubUpdater()
		timeout := int64(5)
		info, err := updater.CheckForUpdate(timeout)

		if err != nil || info == nil {
			t.Error("CheckForUpdate() should pass (there is always something to download)")
		}

		err = os.MkdirAll("../../../assets/tmp", 0700)

		if err != nil {
			t.Error("MkdirAll() should pass")
		}

		err = updater.DownloadFile("../../../assets/tmp/1pass.tar.gz", info.ArchiveUrl, timeout)

		if err != nil {
			t.Error("DownloadFile() should end up with success")
		}

		err = os.RemoveAll("../../../assets/tmp")

		if err != nil {
			t.Error("RemoveAll() should pass")
		}
	} else {
		t.Log("DownloadFile() no internet connection")
	}
}

func TestExtractArchive(t *testing.T) {
	if isOnline() {
		updater := setupGithubUpdater()
		timeout := int64(5)
		info, err := updater.CheckForUpdate(timeout)

		if err != nil || info == nil {
			t.Error("CheckForUpdate() should pass (there is always something to download)")
		}

		err = os.MkdirAll("../../../assets/tmp", 0700)

		if err != nil {
			t.Error("MkdirAll() should pass")
		}

		err = updater.DownloadFile("../../../assets/tmp/1pass.tar.gz", info.ArchiveUrl, timeout)

		if err != nil {
			t.Error("DownloadFile() should end up with success")
		}

		err = updater.ExtractArchive("../../../assets/tmp/1pass.tar.gz", "../../../assets/tmp")

		if err != nil {
			t.Error("ExtractArchive() should pass")
		}

		err = os.RemoveAll("../../../assets/tmp")

		if err != nil {
			t.Error("RemoveAll() should pass")
		}
	} else {
		t.Log("ExtractArchive() no internet connection")
	}
}

func TestShouldCheck(t *testing.T) {
	updater := setupGithubUpdater()
	expected := false
	should := updater.ShouldCheck(10000, "../../../assets")

	if should != expected {
		t.Errorf("ShouldCheck() = %v; expected = %v", should, expected)
	}

	expected = true
	should = updater.ShouldCheck(0, "../../../assets")

	if should != expected {
		t.Errorf("ShouldCheck() = %v; expected = %v", should, expected)
	}
}

func TestValidateChecksum(t *testing.T) {
	updater := setupGithubUpdater()
	err := updater.ValidateChecksum("../../../assets/gif/1pass-categories.gif", "../../../assets/md5/checksum.md5")

	if err != nil {
		t.Error("ValidateChecksum() should pass")
	}

	err = updater.ValidateChecksum("../../../assets/empty/.gitkeep", "../../../assets/md5/checksum.md5")

	if err == nil {
		t.Error("ValidateChecksum() should fail")
	}
}
