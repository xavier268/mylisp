// expression grammar for mylisp

%{

package parser

import (

)

func init() {
    myErrorVerbose = true
    myDebug = 0
}


%}

%union{
    // define SymType structure, used to communicate with lexer
    value Term
}

%type <value> top expr atom atoms 

%token <value> '(' ')' '.' '\''
%token <value> NUMBER IDENT STRING ERROR SKIP

%%

top:
    atom                     { $$ = $1
                              mylex.(*myLex).LastResult = $$
                             }
                

expr:  
    '(' ')'                  { $$ = Cell{}  }
    | '('  atoms  ')'        { $$ = $2   }
    | '(' atom '.' atom ')'  { $$ = Cell{ $2, $4}  }
    | '(' atom '.'  ')'      { $$ = Cell{ $2, nil}  }
    | '('  '.' atom ')'      { $$ = Cell{nil, $3}  }
    | '('  '.'  ')'          { $$ = Cell{}  }
   

    

atoms:
    atom                     { $$ = Cell { $1, nil}  }
    | atom atoms             { $$ = Cell { $1, $2 }  }

atom:
    IDENT                    { $$ = $1  }
    | NUMBER                 { $$ = $1  }
    | STRING                 { $$ = $1  }
    | expr                   { $$ = $1  }
    |  '\''  atom            { $$ = Cell{  Atom { "tick"}, $2} }  

%%