package dirs

import (
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
