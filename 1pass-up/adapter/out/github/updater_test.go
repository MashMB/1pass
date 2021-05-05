// Unit tests for GitHub updater.
//
// @author TSS

package github

import (
	"os"
	"testing"
)

func setupGithubUpdater() *githubUpdater {
	return NewGithubUpdater()
}

func TestCheckForUpdate(t *testing.T) {
	updater := setupGithubUpdater()
	expected := false
	info, err := updater.CheckForUpdate()

	if err != nil {
		t.Error("CheckForUpdate() should pass")
	}

	if info == nil {
		t.Error("CheckForUpdate() should provide any data")
	}

	if info.Newer != expected {
		t.Errorf("CheckForUpdate() = %v; expected = %v", info.Newer, expected)
	}
}

func TestDownloadFile(t *testing.T) {
	updater := setupGithubUpdater()
	info, err := updater.CheckForUpdate()

	if err != nil || info == nil {
		t.Error("CheckForUpdate() should pass (there is always something to download)")
	}

	err = os.MkdirAll("../../../../assets/tmp", 0700)

	if err != nil {
		t.Error("MkdirAll() should pass")
	}

	err = updater.DownloadFile("../../../../assets/tmp/1pass.tar.gz", info.ArchiveUrl)

	if err != nil {
		t.Error("DownloadFile() should end up with success")
	}

	err = os.RemoveAll("../../../../assets/tmp")

	if err != nil {
		t.Error("RemoveAll() should pass")
	}
}

func TestExtractArchive(t *testing.T) {
	updater := setupGithubUpdater()
	info, err := updater.CheckForUpdate()

	if err != nil || info == nil {
		t.Error("CheckForUpdate() should pass (there is always something to download)")
	}

	err = os.MkdirAll("../../../../assets/tmp", 0700)

	if err != nil {
		t.Error("MkdirAll() should pass")
	}

	err = updater.DownloadFile("../../../../assets/tmp/1pass.tar.gz", info.ArchiveUrl)

	if err != nil {
		t.Error("DownloadFile() should end up with success")
	}

	err = updater.ExtractArchive("../../../../assets/tmp/1pass.tar.gz", "../../../../assets/tmp")

	if err != nil {
		t.Error("ExtractArchive() should pass")
	}

	err = os.RemoveAll("../../../../assets/tmp")

	if err != nil {
		t.Error("RemoveAll() should pass")
	}
}

func TestValidateChecksum(t *testing.T) {
	updater := setupGithubUpdater()
	err := updater.ValidateChecksum("../../../../assets/1pass.yml", "../../../../assets/md5/checksum.md5")

	if err != nil {
		t.Error("ValidateChecksum() should pass")
	}

	err = updater.ValidateChecksum("../../../../assets/empty/.gitkeep", "../../../../assets/md5/checksum.md5")

	if err == nil {
		t.Error("ValidateChecksum() should fail")
	}
}
