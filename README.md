# Snippy

Snippy is a very basic CLI snippet manager that uses [fzf](https://github.com/junegunn/fzf) for filtering and picking a snippet.

![cli](img/cli.gif)

## Quickstart

### Binaries

#### snippy

The CLI depends on [fzf](https://github.com/junegunn/fzf) and expects a file named *data* to be located at *$HOME/.config/snippy*

#### data

The data file contains snippets and is located at *$HOME/.config/snippy/data*.
If the data file doesn't exist it will be automatically created with example data.
Each line represents a snippet and should be defined starting with a name and ending with the snippet value itself between square brackets. Example: `Current directory [pwd]`

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
