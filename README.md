# mblen

Simple tool for decomposing / inspecting UTF-8 strings.

Examples:

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