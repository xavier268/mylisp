// expression grammar for mylisp

%{

package inter

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
    listvalue []Term
}

%type <listvalue> top
%type <value>  expr atom atoms 

%token <value> '(' ')' '.' '\''
%token <value> NUMBER IDENT STRING BOOL ERROR SKIP

%%

top:                         { $$ = []Term{} 
                             mylex.(*myLex).LastResult =[]Term{} 
                             }
    |    atom top            { $$ = append( $$ , $1 ) 
                             $$ = append( $$ , $2 ...) 
                             mylex.(*myLex).LastResult = $$   
                             }
                

expr:  
    '(' ')'                  { $$ = Pair{}  }
    | '('  atoms  ')'        { $$ = $2   }
    | '(' atom '.' atom ')'  { $$ = Pair{ $2, $4}  }
    | '(' atom '.'  ')'      { $$ = Pair{ $2, nil}  }
    | '('  '.' atom ')'      { $$ = Pair{nil, $3}  }
    | '('  '.'  ')'          { $$ = Pair{}  }
   

    

atoms:
    atom                     { $$ = Pair { $1, nil}  }
    | atom atoms             { $$ = Pair { $1, $2 }  }

atom:
    IDENT                    { $$ = $1  }
    | NUMBER                 { $$ = $1  }
    | STRING                 { $$ = $1  }
    | BOOL                   { $$ = $1  }
    | expr                   { $$ = $1  }
    |  '\''  atom            { $$ = Pair{  Symbol { "quote"},  $2 }  }

%%