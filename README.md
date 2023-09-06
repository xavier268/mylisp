# mylisp
My lisp interpreter in Go. For educationnal purposes.

Inspired by https://www.gnu.org/software/emacs/manual/html_node/elisp/index.html#SEC_Contents



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
* The pair ( a . ()) is the same as ( a ), represented as Cell{a,Cell{}} 

