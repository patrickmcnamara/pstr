# pstr

![latest version](https://img.shields.io/github/v/tag/patrickmcnamara/pstr?label=latest%20version)
![last commit](https://img.shields.io/github/last-commit/patrickmcnamara/pstr)
![top language](https://img.shields.io/github/languages/top/patrickmcnamara/pstr)
![licence](https://img.shields.io/github/license/patrickmcnamara/pstr?label=licence)

A [paster.xyz](https://paster.xyz) command line helper tool.
It can post pastes and print pastes.

## Installation

Run `go get -u github.com/patrickmcnamara/pstr`.

## Usage

Running `pstr` will post `stdin` to paster.xyz.
For example, `gpg --export --armor F0F49A6E | pstr` will post that PGP key.
The URL of that paste, `https://paster.xyz/GItSTbZZqzo`, is printed out.

`pstr` can also optionally take a paste ID as the first argument.
It will print that paste to `stdout`.
For example, `pstr GItSTbZZqzo` will post the paste from the previous example.

## Licence

This project is licenced under the European Union Public Licence v1.2.
