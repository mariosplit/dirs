package dirs

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var rootDir string
var selectedDirectory string // Add this at the package level

var rootCmd = &cobra.Command{
	Use:   "dirs [root directory]",
	Short: "Get root directory list of directories with files to be indexed",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the user's desktop directory
		desktopDir, err := GetUserDesktopDir()
		if err != nil {
			fmt.Printf("Failed to get user's desktop directory: %v\n", err)
			os.Exit(1)
		}

		// Check if a root directory is provided as an argument
		if len(args) == 1 {
			rootDir = args[0]
		} else {
			// Set the default rootDir to the user's desktop directory
			rootDir = desktopDir

			// Prompt for the root directory, potentially updating rootDir based on user input.
			rootDir, err = PromptForRootDirectory(rootDir)
			if err != nil {
				fmt.Printf("Prompt failed: %v\n", err)
				os.Exit(1)
			}
		}

		fmt.Printf("You selected: %s\n", rootDir)

		// Call the selectDirectoryCmd.Run function directly
		selectDirectoryCmd.Run(cmd, args)
	},
}

var selectDirectoryCmd = &cobra.Command{
	Use:   "selectDir",
	Short: "Select a directory from a list",
	Run: func(cmd *cobra.Command, args []string) {
		directories, err := ListDirectories(rootDir)
		if err != nil {
			fmt.Printf("Failed to list directories: %v\n", err)
			return
		}

		if len(directories) == 0 {
			fmt.Println("No directories found.")
			return
		}

		// Define the survey prompt
		prompt := &survey.Select{
			Message: "Choose a directory:",
			Options: directories,
		}

		// Show the prompt to the user and save the selected option
		err = survey.AskOne(prompt, &selectedDirectory)
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("You selected: %s\n", selectedDirectory)
	},
}

func init() {
	rootCmd.AddCommand(selectDirectoryCmd)
}

func GetUserDesktopDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	desktopDir := filepath.Join(homeDir, "Desktop")
	return desktopDir, nil
}

func PromptForRootDirectory(defaultDir string) (string, error) {
	prompt := &survey.Input{
		Message: "Press enter for default or enter new root directory:",
		Default: defaultDir,
		Help:    "Enter the path to the directory you want to use as the root, or press Enter to use the default.",
	}

	var result string
	err := survey.AskOne(prompt, &result)
	if err != nil {
		return "", err
	}

	// Clean the path
	result = strings.TrimSpace(result)
	if result == "" {
		// If the user simply pressed Enter, use the default directory
		return defaultDir, nil
	}

	// Check if the directory exists
	if _, err := os.Stat(result); os.IsNotExist(err) {
		return "", fmt.Errorf("directory does not exist: %s", result)
	}

	return result, nil
}

func IsHidden(directory string, fileInfo os.FileInfo) (bool, error) {
	if runtime.GOOS == "windows" {
		// For Windows
		if fileInfo.Mode()&os.ModeDir != 0 {
			// Directory
			pointer, err := syscall.UTF16PtrFromString(directory)
			if err != nil {
				return false, err
			}
			attributes, err := syscall.GetFileAttributes(pointer)
			if err != nil {
				return false, err
			}
			return attributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0, nil
		} else {
			// File
			pointer, err := syscall.UTF16PtrFromString(filepath.Join(directory, fileInfo.Name()))
			if err != nil {
				return false, err
			}
			attributes, err := syscall.GetFileAttributes(pointer)
			if err != nil {
				return false, err
			}
			return attributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0, nil
		}
	} else {
		// For macOS, Linux, and other Unix-like systems
		return strings.HasPrefix(fileInfo.Name(), "."), nil
	}
}

func ListDirectories(rootDir string) ([]string, error) {
	// Check if the directory exists
	if _, err := os.Stat(rootDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("directory does not exist: %s", rootDir)
	}

	var directories []string
	items, err := os.ReadDir(rootDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %s, error: %v", rootDir, err)
	}
	for _, item := range items {
		if item.IsDir() {
			fullPath := filepath.Join(rootDir, item.Name())
			info, err := item.Info()
			if err != nil {
				continue // Optionally handle the error, e.g., log it or skip this item
			}
			hidden, err := IsHidden(fullPath, info)
			if err != nil || hidden {
				continue
			}
			directories = append(directories, fullPath)
		}
	}
	return directories, nil
}

func ChooseDirectory() (string, error) {
	// Execute rootCmd which includes handling the directory selection
	if err := rootCmd.Execute(); err != nil {
		return "", err
	}

	// Return the globally stored selectedDirectory after command execution
	return selectedDirectory, nil
}

func OpenDirectory(path string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "explorer"
		args = []string{path}
	case "darwin":
		cmd = "open"
		args = []string{path}
	case "linux":
		cmd = "xdg-open"
		args = []string{path}
	default:
		return fmt.Errorf("unsupported platform")
	}

	fmt.Printf("Executing command: %s %v\n", cmd, args)
	err := exec.Command(cmd, args...).Start()
	if err != nil {
		fmt.Printf("Failed to open directory: %v\n", err)
		return err
	}
	fmt.Printf("Directory opened successfully\n")
	return nil
}

func CreateDirIfNotExists(dir string, overwrite bool) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating directory: %w", err)
		}
	} else if overwrite {
		err = os.RemoveAll(dir)
		if err != nil {
			return fmt.Errorf("error removing existing directory: %w", err)
		}
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating directory after removal: %w", err)
		}
	}
	return nil
}

func CreateFileIfNotExists(file string, overwrite bool) error {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		_, err = os.Create(file)
		if err != nil {
			return fmt.Errorf("error creating file: %w", err)
		}
	} else if overwrite {
		err = os.Remove(file)
		if err != nil {
			return fmt.Errorf("error removing existing file: %w", err)
		}
		_, err = os.Create(file)
		if err != nil {
			return fmt.Errorf("error creating file after removal: %w", err)
		}
	}
	return nil
}

func GetDirectoryPath(dirType string) (string, error) {
	var dir string

	switch dirType {
	case "exec":
		execPath, err := os.Executable()
		if err != nil {
			return "", fmt.Errorf("error getting executable directory: %w", err)
		}
		dir = filepath.Dir(execPath)

	case "output":
		dir = filepath.Join(".", "output")

	case "userProfile":
		userProfile, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("error getting user profile directory: %w", err)
		}
		dir = userProfile

	case "desktop":
		userProfile, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("error getting user profile directory: %w", err)
		}
		dir = filepath.Join(userProfile, "Desktop")

	case "preferences":
		userProfile, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("error getting user profile directory: %w", err)
		}
		dir = filepath.Join(userProfile, "preferences")

	case "config":
		userProfile, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("error getting user profile directory: %w", err)
		}
		dir = filepath.Join(userProfile, "config")

	case "dropbox":
		userProfile, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("error getting user profile directory: %w", err)
		}
		dir = filepath.Join(userProfile, "Dropbox")

	case "oneDrive":
		userProfile, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("error getting user profile directory: %w", err)
		}
		dir = filepath.Join(userProfile, "OneDrive")

	default:
		return "", fmt.Errorf("unsupported directory type: %s", dirType)
	}

	return dir, nil
}
