// TODO: Support all flags
// TODO: Add tests
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

func main() {
	if err := mainE(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func mainE() error {
	// Flag set
	cli := pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)

	// Pointers to flags
	noCreatePtr := cli.BoolP("no-create", "c", false, "do not create any files")

	// Parse arguments and set values to flag pointers
	cli.Parse(os.Args[1:])

	// Dereference pointers to flags
	noCreate := *noCreatePtr

	paths := cli.Args()

	for _, path := range paths {
		if !noCreate {
			if err := touch(path); err != nil {
				return err
			}
		}
	}

	return nil
}

// Creates a file with standard permissions, if it doesn't exist
func touch(path string) error {
	// GNU implementation uses a modified version of open to check, if the file exists.
	// touch source: https://github.com/coreutils/coreutils/blob/3cb862ce5f10db392cc8e6907dd9d888acfa2a30/src/touch.c#L132
	// fd_reopen source: https://github.com/coreutils/coreutils/blob/3cb862ce5f10db392cc8e6907dd9d888acfa2a30/gl/lib/fd-reopen.c#L32

	// Serenity OS uses stat to check, if a file exists.
	// touch source: https://github.com/SerenityOS/serenity/blob/8ed3cc5f7b1f84a4499cfcb4e4eae1785fae8c2e/Userland/Utilities/touch.cpp#L244
	// LibFileSystem FileSystem::check source https://github.com/SerenityOS/serenity/blob/8ed3cc5f7b1f84a4499cfcb4e4eae1785fae8c2e/Userland/Libraries/LibFileSystem/FileSystem.cpp#L62C27-L62C27
	fileExists, err := fileExists(path)
	if err != nil {
		return err
	}

	if !fileExists {
		// Even though, go's os.Create differs from the GNU implementation, it
		// is used, because it is more idiomatic and readable.
		// The golang implementation, doesn't set the "O_NONBLOCK" and "O_NOCTTY"
		// flags and uses a, O_TRUNC, which will not be effective, because we checked
		// for the files existence above.

		// The file get's created with the following permissions: a+rwx,u-x,g-x,o-x (before umask)
		if _, err := os.Create(path); err != nil {
			return err
		}
	}

	return nil
}

// Checks if a file exists
func fileExists(path string) (bool, error) {
	// Even though, golang's os.Open differs from the GNU implementation, it
	// is used, because it is more idiomatic and readable.
	// The golang implementation, opens the file in readonly mode, whereas
	// the GNU implementation uses write-only.
	_, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return true, err
	}
}
