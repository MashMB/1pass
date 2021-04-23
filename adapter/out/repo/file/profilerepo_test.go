// Unit tests for file implementation of profile repository.
//
// @author TSS

package file

import (
	"testing"

	"github.com/mashmb/1pass/core/domain"
)

func setupFileProfileRepo() *fileProfileRepo {
	repo := NewFileProfileRepo()
	vault := domain.NewVault("../../../../assets/onepassword_data")
	repo.LoadProfile(vault)

	return repo
}

func TestGetIterations(t *testing.T) {
	repo := setupFileProfileRepo()
	expected := 50000
	iterations := repo.GetIterations()

	if iterations != expected {
		t.Errorf("GetIterations() = %d; expected %d", iterations, expected)
	}
}

func TestGetMasterKey(t *testing.T) {
	repo := setupFileProfileRepo()
	expected := "b3BkYXRhMDEAAQAAAAAAACN8JuE76yN6hbjqzEvd0RGnu3vufPcfAZ35JoyzdR1WPRvr8DMefe9MJu65DmHSwjObPC0jznXpafJQob6CNzKCNoeVC+GXIvLckvAuYUNSwILQQ1jEIcHdyQ0H2MbJ+0YlWEbvlQ8UVH5bcrMqDmTPPSRkbUG3/dV1NKHdgI0V6N/kKZ737oo+kj3ChJZQTKywvmR6RgB5et5stBaUwutNQbZ0znYtZumIlf3pjdqGK4RyCHSwmwgLUO+VFLTqDjoZ9dUcy4hQzSZiPlba3vK8vGJRlN0Qf2Y6dUj5kYAwdYdOzE/Ji3hbTNVsPOm8sjzPcPGQj8haW5UgzSDZ0mo7+ymsKJwSYjAsgvawh31WY2m5j7VR+50ERDTEyxxQ3LW7WgetAxX9l0LX0O3Jue1oW/p2l44ij9qiN9rkFScx"
	masterKey := repo.GetMasterKey()

	if masterKey != expected {
		t.Errorf("GetMasterKey() = %v; expected %v", masterKey, expected)
	}
}

func TestGetOverviewKey(t *testing.T) {
	repo := setupFileProfileRepo()
	expected := "b3BkYXRhMDFAAAAAAAAAAIy1hZwIGeiLn4mLE1R8lEwIOye95GEyfZcPKlyXkkb0IBTfCXM+aDxjD7hOliuTM/YMIqxK+firVvW3c5cp2QMgvQHpDW2AsAQpBqcgBgRUCSP+THMVg15ZeR9lI77mHBpTQ70D+bchvkSmw3hoEGot7YcnQCATbouhMXIMO52D"
	overviewKey := repo.GetOverviewKey()

	if overviewKey != expected {
		t.Errorf("GetOverviewKey() = %v; expected %v", overviewKey, expected)
	}
}

func TestGetSalt(t *testing.T) {
	repo := setupFileProfileRepo()
	expected := "P0pOMMN6Ow5wIKOOSsaSQg=="
	salt := repo.GetSalt()

	if salt != expected {
		t.Errorf("GetSalt() = %v; expected %v", salt, expected)
	}
}

func TestLoadProfile(t *testing.T) {
	repo := NewFileProfileRepo()
	vault := domain.NewVault("../../../../assets/onepassword_data")
	repo.LoadProfile(vault)

	if repo.profileJson == nil || len(repo.profileJson) == 0 {
		t.Error("LoadProfile() should initialize repository")
	}
}
