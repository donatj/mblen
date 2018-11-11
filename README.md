# mblen

Simple CLI for decomposing / inspecting UTF-8 strings.

## Installation

Binaries are availible for Darwin (macOS), Linux and Windows on the release page:

https://github.com/donatj/mblen/releases

### From Source

```bash
go get -u -v github.com/donatj/mblen
```

## Examples

```bash
$ mblen 🌋
       rune: 🌋
        ord: 127755
byte length: 4
       name: VOLCANO

$ mblen 🇰🇷
       rune: 🇰
        ord: 127472
byte length: 4
       name: REGIONAL INDICATOR SYMBOL LETTER K

       rune: 🇷
        ord: 127479
byte length: 4
       name: REGIONAL INDICATOR SYMBOL LETTER R
```

Also if no arguments are given it will read from STDIN.

```bash
$ echo 日本語 | mblen
       rune: 日
        ord: 26085
byte length: 3
       name: SUN; DAY; DAYTIME

       rune: 本
        ord: 26412
byte length: 3
       name: ROOT, ORIGIN, SOURCE; BASIS

       rune: 語
        ord: 35486
byte length: 3
       name: LANGUAGE, WORDS; SAYING, EXPRESSION
```
