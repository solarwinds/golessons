# golessons
A series of interactive lessons teaching Go

## Pre-Work

### Install Git
We will be working with Git and GitHub in this workshop. If you haven't installed Git please do so:

#### Windows
* Download official EXE: [https://git-scm.com/download/win](https://git-scm.com/download/win)

#### macOS
(macOS comes with Git pre-installed, but it's a good practice to control the installation so you can upgrade later if you wish)

* `$> brew install git`

#### Ubuntu
* `$> sudo apt update`
* `$> sudo apt-get install git`



### Install Go
You need the Go compiler and CLI tools so you can build/test/analyze Go programs. Follow the instructions below to get everything installed for your platform.

#### Windows
* Install from MSI or source https://golang.org/dl/

#### macOS
* `$> brew install go`

#### Ubuntu
* `$> sudo apt update`
* `$> sudo apt-get install golang-1.9-go`

### Setup $GOPATH
One of the great strengths of is the `go` CLI tool, which you’ll use to manage source, build binaries, etc. The `go` tool expects a $GOPATH env var to be set to a directory path where the Go workspace will live.

That directory contains `src`, `bin`, and `pkg` directories. For now just know that all Go source goes into `src`, organized by domain. The `go get` command takes care of this

If the env var isn’t set, the tool will install will automatically set/use `$HOME/go` on Linux or %USERPROFILE%\go. If you want it somewhere else, create the directory and then add `src`,`bin`, and `pkg` subdirectories manually.

Don't forget to add `$GOPATH` to your environment permanently. For POSIX sysems, this is [what I use](https://github.com/trevrosen/dotfiles/blob/master/zshrc_osx#L14-L19)

### Check your installation and start the Tour
Confirm that your install is working by installing the interactive Go tour. 

```
$> go get golang.org/x/tour/gotour
$> gotour
```

Proceed to complete the first lesson. It gives a good overview of how Go works and will let you feel comfortable starting to read source code.

### Choose an Editor
You can edit Go code however you wish, but at a minimum you'll want something that runs formatting and import checking on save and that allows you to jump into function definitions whether in your code or code imported from elsewhere. For all editors, you'll want to have [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) installed: `go get golang.org/x/tools/cmd/goimports`

#### GoLand (recommended)
The only full-featured IDE available for Go is [GoLand from JetBrains](https://www.jetbrains.com/go/) Thankfully, it's very good. If you're familiar with IntelliJ, PyCharm, CLion, or any of the other IntelliJ products, you'll be right at home with this editor.

#### VSCode
Microsoft's free VSCode editor is quite popular in the Go community. It bears a superficial resemblence to VisualStudio and is **not** an IDE, but the Go packages are very good. After installing VSCode, make sure to install the `vscode-go` package. ([VSCode download](https://code.visualstudio.com), [vscode-go GitHub page](https://github.com/Microsoft/vscode-go))

#### Vim (macvim, gvim, neovim)
If you're a Vim user, you can install `vim-go`. Make sure to setup goimports to work with it as an on-save functionality.