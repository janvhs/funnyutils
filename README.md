# Funny Utils

A reimplementation of the GNU core utilities in the Go programming language.

This implementation uses only the Go standard library and <https://github.com/spf13/pflag> package.

## Goals

This implementation has the following goals:

1. Implementing a set of the core utilities in a modern, readable language
2. Being able to be run on multiple platforms
3. Making it possible to teach the "unix philosophy" to new developers by
   preferring readability over performance
4. Using only the standard library and a limited amount of selected packages

## Non Goals

These goals are not pursued:

1. Implementing all of the GNU core utilities
2. Being a drop-in replacement
3. Being "bug for bug" compatible with the GNU implementation
4. Having the best performance possible

If you're looking for core utilities, which pursue these goals and are written in a
memory safe and modern language, you might want to take a look at the [uutils](https://github.com/uutils/coreutils) project, which is written in [rust](https://www.rust-lang.org/).

## Implementation references

This implementation references the following implementations:

- [GNU core utilities](https://github.com/coreutils/coreutils/tree/3cb862ce5f10db392cc8e6907dd9d888acfa2a30)
- [SerenityOS Utilities](https://github.com/SerenityOS/serenity/tree/8ed3cc5f7b1f84a4499cfcb4e4eae1785fae8c2e/Userland/Utilities)
