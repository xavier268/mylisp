
[![Go Reference](https://pkg.go.dev/badge/github.com/xavier268/mylisp.svg)](https://pkg.go.dev/github.com/xavier268/mylisp) 
[![Go Report Card](https://goreportcard.com/badge/github.com/xavier268/mylisp)](https://goreportcard.com/report/github.com/xavier268/mylisp)

# mylisp
My lisp interpreter in Go. For educationnal purposes.

Inspired by scheme syntax.



See https://web.mit.edu/scheme/scheme_v9.2/doc/mit-scheme-ref/index.html  (prefered) 
or  https://www.gnu.org/software/guile/manual/guile.html#SEC_Contents
or https://www.scheme.com/tspl4/

# Examples

Examples lisp/scheme files are in the examples folder.

You can run an example directly from the command line with :

```
go run . -l "examples/couter.scm"
```

## Builtins and help

You can list the available keywords built-in with * (keywords )*.

You can display help on one or more keywoards ( kw1, kw2 ...) by evaluationg ( help kw1 kw2 )

## Non standard keywords :

New keywords :

* ( debug value ) sets the debug/trace level
* ( join-path p1 p2 ... ) constructs a path from path components
* ( load fullPathName ) loads and evaluate a file
* ( file-sep ) retuns the system file separator
* ( version ) gets the current version of the interpreter

Extended behaviour vs standard :
* (display x1 x2 ...) : display can display zero or more than one Terms.


 
## Numbers

* Numbers are integer or rationnal. 
* The decimal point is not accepted. 
* Rational are written 8/3. ( no space ! )
* Rational are immutable and always normalized, with a positive denominator. 
* Zero is 0/1, while 1/0 represents NaN.

## Strings

* string litterlas are quoted with douple quotes (")
* there is no escaping, string litteral may contain newline directly


## Errors 

Error is a specific type that can be returned from evaluation.
