package dirs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetDirectoryPath(t *testing.T) {
	// Test case 1: Test getting the executable directory
	execDir, err := GetDirectoryPath("exec")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else {
		t.Logf("Executable directory: %s", execDir)
	}

	// Test case 2: Test getting the output directory
	outputDir, err := GetDirectoryPath("output")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else {
		t.Logf("Output directory: %s", outputDir)
	}

	// Test case 3: Test getting the user profile directory
	userProfileDir, err := GetDirectoryPath("userProfile")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else {
		t.Logf("User profile directory: %s", userProfileDir)
	}

	// Test case 4: Test getting the user desktop directory
	desktopDir, err := GetDirectoryPath("desktop")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else {
		t.Logf("User desktop directory: %s", desktopDir)
		expectedDesktopDir := filepath.Join(userProfileDir, "Desktop")
		if desktopDir != expectedDesktopDir {
			t.Errorf("Expected desktop directory: %s, but got: %s", expectedDesktopDir, desktopDir)
		}
	}

	// Test case 5: Test getting the preferences directory
	preferencesDir, err := GetDirectoryPath("preferences")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else {
		t.Logf("Preferences directory: %s", preferencesDir)
	}

	// Test case 6: Test getting the config directory
	configDir, err := GetDirectoryPath("config")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else {
		t.Logf("Config directory: %s", configDir)
	}

	// Test case 7: Test getting the Dropbox directory
	dropboxDir, err := GetDirectoryPath("dropbox")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else {
		t.Logf("Dropbox directory: %s", dropboxDir)
	}

	// Test case 8: Test getting the OneDrive directory
	oneDriveDir, err := GetDirectoryPath("oneDrive")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else {
		t.Logf("OneDrive directory: %s", oneDriveDir)
	}

	// Test case 9: Test getting an unsupported directory type
	_, err = GetDirectoryPath("unsupported")
	if err == nil {
		t.Error("Expected an error for unsupported directory type, but got nil")
	} else {
		t.Logf("Unsupported directory type error: %s", err.Error())
	}
}

func TestCreateDirIfNotExists(t *testing.T) {
	// Test case 1: Create a new directory (overwrite = false)
	newDir := filepath.Join(os.TempDir(), "test_dir")
	err := CreateDirIfNotExists(newDir, false)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	defer os.RemoveAll(newDir)

	// Check if the directory exists
	if _, err := os.Stat(newDir); os.IsNotExist(err) {
		t.Error("Expected directory to be created, but it doesn't exist")
	}

	// Test case 2: Try to create an existing directory (overwrite = false)
	err = CreateDirIfNotExists(newDir, false)
	if err != nil {
		t.Errorf("Unexpected error when creating an existing directory: %v", err)
	}

	// Test case 3: Overwrite an existing directory (overwrite = true)
	err = CreateDirIfNotExists(newDir, true)
	if err != nil {
		t.Errorf("Unexpected error when overwriting an existing directory: %v", err)
	}
}

func TestCreateFileIfNotExists(t *testing.T) {
	// Test case 1: Create a new file (overwrite = false)
	newFile := filepath.Join(os.TempDir(), "test_file.txt")
	err := CreateFileIfNotExists(newFile, false)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	defer os.Remove(newFile)

	// Check if the file exists
	if _, err := os.Stat(newFile); os.IsNotExist(err) {
		t.Error("Expected file to be created, but it doesn't exist")
	}

	// Test case 2: Try to create an existing file (overwrite = false)
	err = CreateFileIfNotExists(newFile, false)
	if err != nil {
		t.Errorf("Unexpected error when creating an existing file: %v", err)
	}

	// Test case 3: Overwrite an existing file (overwrite = true)
	err = CreateFileIfNotExists(newFile, true)
	if err != nil {
		t.Errorf("Unexpected error when overwriting an existing file: %v", err)
	}
}
