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
    expr                     { $$ = $1}
                

expr:
    '(' ')'                  { $$ = nil  }
    | '('  atoms  ')'        { $$ = $2   }
    | '(' atom '.' atom ')'  { $$ = $2  }
    |  '\''  expr            { $$ = $2 }  
    

atoms:
    atom                     { $$ = $1  }
    | atom atoms             { $$ = $1  }

atom:
    ATOM                     { $$ = $1  }
    | NUMBER                 { $$ = $1  }
    | STRING                 { $$ = $1  }
    | expr                   { $$ = $1  }


%%