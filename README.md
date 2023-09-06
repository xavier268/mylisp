# mylisp
My lisp interpreter in Go. For educationnal purposes.

Inspired by scheme syntax.



See https://web.mit.edu/scheme/scheme_v9.2/doc/mit-scheme-ref/index.html  (prefered) or  https://www.gnu.org/software/guile/manual/guile.html#SEC_Contents


Redesigned needed to comply :

* think about error reporting  as a separate return value ? as a list in the interpreter ?
* check lists are not implementing an end-of-list empty cell as a marker, but only a nil cdr on the last cell.
* redesign scopes and names ; ensure names point to the *location* where value is stored. *Locations* should contain also some meta information about what is stored : unassigned/unbound, variables, procedures, macros, ...
* redo eval ... following mit-scheme 


# Examples


# Implementation notes


## Variable



## Builtins

Target special forms are ( Gnu MIT Scheme)

access 	and 	begin case 	cond 	cons-stream declare 	define define-integrable 	define-structure 	define-syntax
delay 	do 	er-macro-transformer fluid-let 	if 	lambda let 	let* 	let*-syntax let-syntax 	letrec 	letrec-syntax
local-declare 	named-lambda 	non-hygienic-macro-transformer or 	quasiquote 	quote rsc-macro-transformer 	sc-macro-transformer 	set!
syntax-rules 	the-environment 
 
## Numbers

* Numbers are integer or rationnal. 
* The decimal point is not accepted. 
* Rational are written 8/3. 
* Rational are immutable and always normalized, with a positive denoominator. 
* Zero is 0/1, while 1/0 represents NaN.

## Internal representation of lists

The base construction is the Pair object. 
* () is represented as Pair{}. The go nil **does not** represent the empty list.
* ( a ) is Pair {a,Pair{}}. 
* The pair, ( a . ), represented as Pair {a}.
* ~~ The pair ( a . ()) is the same as ( a ), represented as Pair{a,Cell{}} ~~

