

# Directory Utility Package

The Directory Utility Package (`dirs`) provides a set of functions and commands for working with directories in Go. It allows you to retrieve various directory paths, prompt for a root directory, list non-hidden directories, choose a directory from a list, open a directory using the appropriate system command, and create directories and files if they don't exist, with the option to overwrite existing ones.

## Installation

To use the `dirs` package in your Go project, you can install it using the following command:

```
go get github.com/mariosplit/dirs
```

## Usage

Import the `dirs` package in your Go code:

```go
import "github.com/mariosplit/dirs"
```

### Functions

The `dirs` package provides the following functions:

- `GetUserDesktopDir() (string, error)`: Retrieves the path to the user's desktop directory.
- `PromptForRootDirectory(defaultDir string) (string, error)`: Prompts the user to enter a root directory or use the default directory.
- `IsHidden(directory string, fileInfo os.FileInfo) (bool, error)`: Checks if a file or directory is hidden based on the operating system.
- `ListDirectories(rootDir string) ([]string, error)`: Lists all the non-hidden directories within the specified root directory.
- `ChooseDirectory() (string, error)`: Prompts the user to select a directory from a list of directories.
- `OpenDirectory(path string) error`: Opens the specified directory using the appropriate command based on the operating system.
- `GetDirectoryPath(dirType string) (string, error)`: Retrieves the path of a specific directory based on the provided directory type.
- `CreateDirIfNotExists(dir string, overwrite bool) error`: Creates a directory if it doesn't exist. If `overwrite` is true, it overwrites an existing directory.
- `CreateFileIfNotExists(file string, overwrite bool) error`: Creates a file if it doesn't exist. If `overwrite` is true, it overwrites an existing file.

The `GetDirectoryPath` function supports the following directory types:
- `"exec"`: Returns the directory path of the executable file.
- `"output"`: Returns the path of the "output" directory relative to the current directory.
- `"userProfile"`: Returns the path of the user's home directory.
- `"desktop"`: Returns the path of the user's desktop directory.
- `"preferences"`: Returns the path of the "preferences" directory within the user's home directory.
- `"config"`: Returns the path of the "config" directory within the user's home directory.
- `"dropbox"`: Returns the path of the "Dropbox" directory within the user's home directory.
- `"oneDrive"`: Returns the path of the "OneDrive" directory within the user's home directory.

### Cobra Commands

The `dirs` package also provides Cobra commands for handling command-line arguments and prompting the user for input:

- `rootCmd`: The root command that initializes the directory selection process.
- `selectDirectoryCmd`: A subcommand that prompts the user to select a directory from a list.

## Examples

Here are a few examples of how to use the `dirs` package:

```go
package main

import (
    "fmt"
    "log"

    "github.com/mariosplit/dirs"
)

func main() {
    // Retrieve the user's desktop directory
    desktopDir, err := dirs.GetUserDesktopDir()
    if err != nil {
        log.Fatalf("Failed to get user's desktop directory: %v", err)
    }
    fmt.Println("Desktop directory:", desktopDir)

    // Prompt for a root directory
    rootDir, err := dirs.PromptForRootDirectory(desktopDir)
    if err != nil {
        log.Fatalf("Failed to prompt for root directory: %v", err)
    }
    fmt.Println("Selected root directory:", rootDir)

    // List non-hidden directories within the root directory
    directories, err := dirs.ListDirectories(rootDir)
    if err != nil {
        log.Fatalf("Failed to list directories: %v", err)
    }
    fmt.Println("Directories:", directories)

    // Choose a directory from the list
    selectedDir, err := dirs.ChooseDirectory()
    if err != nil {
        log.Fatalf("Failed to choose directory: %v", err)
    }
    fmt.Println("Selected directory:", selectedDir)

    // Open the selected directory
    err = dirs.OpenDirectory(selectedDir)
    if err != nil {
        log.Fatalf("Failed to open directory: %v", err)
    }

    // Retrieve various directory paths
    execDir, err := dirs.GetDirectoryPath("exec")
    if err != nil {
        log.Fatalf("Failed to get executable directory: %v", err)
    }
    fmt.Println("Executable directory:", execDir)

    outputDir, err := dirs.GetDirectoryPath("output")
    if err != nil {
        log.Fatalf("Failed to get output directory: %v", err)
    }
    fmt.Println("Output directory:", outputDir)

    userProfileDir, err := dirs.GetDirectoryPath("userProfile")
    if err != nil {
        log.Fatalf("Failed to get user profile directory: %v", err)
    }
    fmt.Println("User profile directory:", userProfileDir)

    // Create a directory if it doesn't exist (overwrite = false)
    newDir := "new_directory"
    err = dirs.CreateDirIfNotExists(newDir, false)
    if err != nil {
        log.Fatalf("Failed to create directory: %v", err)
    }
    fmt.Println("Directory created:", newDir)

    // Overwrite an existing directory (overwrite = true)
    existingDir := "existing_directory"
    err = dirs.CreateDirIfNotExists(existingDir, true)
    if err != nil {
        log.Fatalf("Failed to overwrite directory: %v", err)
    }
    fmt.Println("Directory overwritten:", existingDir)

    // Create a file if it doesn't exist (overwrite = false)
    newFile := "new_file.txt"
    err = dirs.CreateFileIfNotExists(newFile, false)
    if err != nil {
        log.Fatalf("Failed to create file: %v", err)
    }
    fmt.Println("File created:", newFile)

    // Overwrite an existing file (overwrite = true)
    existingFile := "existing_file.txt"
    err = dirs.CreateFileIfNotExists(existingFile, true)
    if err != nil {
        log.Fatalf("Failed to overwrite file: %v", err)
    }
    fmt.Println("File overwritten:", existingFile)
}
```

## Tests

The `dirs` package includes a set of tests to ensure the correctness of its functions. You can run the tests using the following command:

```go
go test -v
```

The tests cover various scenarios and edge cases to validate the behavior of the package, including the new functionality for overwriting existing directories and files.

## Contributing

Contributions to the `dirs` package are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on the [GitHub repository](https://github.com/mariosplit/dirs).

## License

The `dirs` package is open-source software licensed under the [MIT License](https://opensource.org/licenses/MIT).


Now the README.md file includes the existing examples in the specified location, demonstrating the usage of various functions provided by the `dirs` package.

Make sure to replace the existing README.md file with this updated version to reflect the latest modifications and include the existing examples.