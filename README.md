# Snippy

Snippy is a very basic CLI snippet manager that uses [fzf](https://github.com/junegunn/fzf) for filtering and picking a snippet.

![cli](img/cli.gif)

## Quickstart

### Binaries

#### snippy

The CLI depends on [fzf](https://github.com/junegunn/fzf) and expects a file named [data](#data) to be located at `$HOME/.config/snippy`

### Installation

Make sure [fzf](https://github.com/junegunn/fzf) is installed and added to *$PATH*.

The snippy binary will be placed at the GOBIN path when using *make install*.

```bash
make install
```

### Shell examples

ZSH

```bash
# fcs - command snippets
fcs() {
  print -z $(snippy)
}
```

## Data

The data file contains snippets and is located at the following path. If the file doesn't exist it will be automatically created with example data.

```text
$HOME/.config/snippy/data
```

Example content

```text
Current directory [pwd]
List [ls]
Move or rename files [mv]
```

Each line represents a snippet and should be defined starting with a name and ending with the snippet value itself between square brackets.
