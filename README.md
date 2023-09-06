# mylisp
My lisp interpreter in Go. For educationnal purposes.

Inspired by scheme syntax.

See https://www.gnu.org/software/guile/manual/guile.html#SEC_Contents

See https://web.mit.edu/scheme/scheme_v9.2/doc/mit-scheme-ref/index.html


Redesigned needed to comply :

* check lists are not implementing an end-of-list empty cell as a marker, but only a nil cdr on the last cell.
* redo eval


# Examples


# Implementation notes


## Variable



## Builtins

 

## Numbers

* Numbers are integer or rationnal. 
* The decimal point is not accepted. 
* Rational are written 8/3. 
* Rational are immutable and always normalized, with a positive denoominator. 
* Zero is 0/1, while 1/0 represents NaN.
## Internal representation of lists

The base construction is the Cell object. 
* () is represented as Cell{}. The go nil **does not** represent the empty list.
* ( a ) is Cell {a,Cell{}}. 
* The pair, ( a . ), represented as Cell {a}.
* ~~ The pair ( a . ()) is the same as ( a ), represented as Cell{a,Cell{}} ~~

