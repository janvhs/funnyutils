// TODO: Support all flags
// TODO: Add tests
// TODO: Add show usage error
// TODO: Handle help error
package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

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
	flag := pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)

	flag.Usage = func() {
		fmt.Printf(
			"usage: %s [--help] [-ch] [-r file]\n       file ...\n", // FIXME: use tab here?
			filepath.Base(os.Args[0]),
		)
	}

	var help bool
	flag.BoolVar(&help, "help", false, "show help")

	// Define flags
	var noCreate bool
	flag.BoolVarP(&noCreate, "no-create", "c", false, "do not create a new file")

	// Required
	var reference string
	flag.StringVarP(&reference, "reference", "r", "", "use times from a reference file")

	var noDereference bool
	flag.BoolVarP(&noDereference, "no-dereference", "h", false, "change the times of an existing symlink")

	// Parse arguments and set values to flag pointers
	if err := flag.Parse(os.Args[1:]); err != nil {
		return err
	}

	if help {
		return pflag.ErrHelp
	}

	// FIXME: Implement flags
	if reference != "" || noDereference {
		panic("flag not implemented")
	}

	paths := flag.Args()

	for _, path := range paths {
		if !noCreate {
			_, err := touch(path)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Creates a file with standard permissions, if it doesn't exist
func touch(path string) (*os.File, error) {
	// GNU implementation uses a modified version of open to check, if the file exists.
	// touch source: https://github.com/coreutils/coreutils/blob/3cb862ce5f10db392cc8e6907dd9d888acfa2a30/src/touch.c#L132
	// fd_reopen source: https://github.com/coreutils/coreutils/blob/3cb862ce5f10db392cc8e6907dd9d888acfa2a30/gl/lib/fd-reopen.c#L32

	// Serenity OS uses stat to check, if a file exists.
	// touch source: https://github.com/SerenityOS/serenity/blob/8ed3cc5f7b1f84a4499cfcb4e4eae1785fae8c2e/Userland/Utilities/touch.cpp#L244
	// LibFileSystem FileSystem::check source https://github.com/SerenityOS/serenity/blob/8ed3cc5f7b1f84a4499cfcb4e4eae1785fae8c2e/Userland/Libraries/LibFileSystem/FileSystem.cpp#L62C27-L62C27
	fileExists, file, err := fileExists(path)
	if err != nil {
		return nil, err
	}

	if !fileExists {
		// Even though, go's os.Create differs from the GNU implementation, it
		// is used, because it is more idiomatic and readable.
		// The golang implementation, doesn't set the "O_NONBLOCK" and "O_NOCTTY"
		// flags and uses a, O_TRUNC, which will not be effective, because we checked
		// for the files existence above.

		// The file get's created with the following permissions: a+rwx,u-x,g-x,o-x (before umask)
		file, err := os.Create(path)
		if err != nil {
			return file, err
		}
	}

	return file, nil
}

// Checks if a file exists
func fileExists(path string) (bool, *os.File, error) {
	// Even though, golang's os.Open differs from the GNU implementation, it
	// is used, because it is more idiomatic and readable.
	// The golang implementation, opens the file in readonly mode, whereas
	// the GNU implementation uses write-only.
	// FIXME: Open for writing to apply stat changes
	f, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return false, f, nil
	} else {
		return true, f, err
	}
}
