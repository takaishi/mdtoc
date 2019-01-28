# mdtoc

[![Build Status](https://travis-ci.org/takaishi/mdtoc.svg?branch=master)](https://travis-ci.org/takaishi/mdtoc)

Generatr and insert table of content with markdown.

## Usage

```
$ mdtoc --help
NAME:
   mdtoc - A new cli application

USAGE:
   mdtoc [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --file value, -f value   Specify to generate TOC.
   --in-file, -i            Insert TOC to md file specified --file option.
   --level value, -l value  (default: 2)
   --help, -h               show help
   --version, -v            print the version
```

## Install

```
$ brew install takaishi/fomulas/patt
```

## Example

```md
$ cat example.md
# This is example

<!-- toc -->

## foo

aaa

## bar

bbb
```

Output markdown with TOC to stcout:

```
$ mdtoc -f ./example.md
# This is example

<!-- toc -->
<!-- toc:start -->

  * [foo](#foo)
  * [bar](#bar)

<!-- toc:end -->

## foo

aaa

## bar

bbb
```

Overwrite file with TOC :

```
$ mdtoc -f ./example.md -i
$
$ cat example.md
# This is example

<!-- toc -->
<!-- toc:start -->

  * [foo](#foo)
  * [bar](#bar)

<!-- toc:end -->

## foo

aaa

## bar

bbb
```

## Copyright

* Copyright(c) 2019- r_takaishi
