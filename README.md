# Directory Utility Package

The Directory Utility Package (`dirs`) provides a set of functions and commands for working with directories in Go. It allows you to retrieve the user's desktop directory, prompt for a root directory, list non-hidden directories, choose a directory from a list, and open a directory using the appropriate system command.

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
}
```

## Contributing

Contributions to the `dirs` package are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on the [GitHub repository](https://github.com/mariosplit/dirs).

## License

The `dirs` package is open-source software licensed under the [MIT License](https://opensource.org/licenses/MIT).
