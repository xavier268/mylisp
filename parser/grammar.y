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
    // define SymType structure
    value Term
}

%type <value> top expr atom atoms 

%token <value> '(' ')' '.' '\'' 
%token <value> NUMBER ATOM STRING

%%

top:
    expr                     { $$ = $1
                              mylex.(*myLex).LastResult = $$
                             }
                

expr:
    '(' ')'                  { $$ = Cell {}  }
    | '('  atoms  ')'        { $$ = $2   }
    | '(' atom '.' atom ')'  { $$ = Cell{ $2, $4}  }
    |  '\''  expr            { $$ = Cell{  String { "quote"}, $1} }  
    

atoms:
    atom                     { $$ = Cell { $1, nil}  }
    | atom atoms             { $$ = Cell { $1, $2 }  }

atom:
    ATOM                     { $$ = $1  }
    | NUMBER                 { $$ = $1  }
    | STRING                 { $$ = $1  }
    | expr                   { $$ = $1  }


%%