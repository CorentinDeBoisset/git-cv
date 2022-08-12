# git-cv

Tool to make commits that follow a convention - such as conventionnal commits

## ❯ Install

To install this tool, you can check out the latest release and install it using a pre-built package.

Alternatively, you can build it from the source (you will require golang v1.18+),
by cloning the repository and run:

    $ make install

## ❯ Usage

Run the executable using:

    $ git cv

If you want information about the usage, use:

    $ git cv --help

All the flags are the same as the `git commit` subcommand

## ❯ Contributing

If you want to open an MR, be sure to run the tests with:

    $ golangci-lint run
    $ go test ./...

If you want to run all these tests automatically before every commit, add the custom git-hooks with:

    $ git config core.hooksPath .githooks
