# Parsers & Lexers

What is a parser?

When we parse a language (grammer like) we do it in two phases. First we break up
series of characters into tokens.
For SQL like language these tokens may be "whitespace", "number", "select", etc.
This process is called lexing (or tokenizing or scanning).

The tokens get fed to parser. Parser constructs abstract syntax tree (AST) from series
of tokens and the AST is what our application use.

## References

Handwritten Parsers & Lexers in Go **[Article](https://blog.gopheracademy.com/advent-2014/parsers-lexers/)**
