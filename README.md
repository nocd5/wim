# Wim

Kick your Vim for Windows

## Installation

`go get github.com/nocd5/wim`

## Requirement

`vim.exe` & `gvim.exe` in your path.

## `-g` option

### Before

```sh
$ vim.exe -g
E25: GUI cannot be used: Not enabled at compile time
exit status 2
```

### After

```sh
$ wim.exe -g
```

Kick your GVim.

## `|` w/o `-`

### Before

```sh
$ echo Hello World | vim.exe
Vim: Warning: Input is not from a terminal
```

And lose control ...

### After

```sh
$ echo Hello World | wim.exe
```

Append `-` option and kick your Vim.
