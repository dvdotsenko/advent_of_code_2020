grammar Expressions;

// Tokens
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
NUMBER: [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;

// this turns into
// [Parser].TreeStart() method returning antlr.ParseTree
treeStart : value=expr EOF;

expr
   : left=expr op=('*'|'/') right=expr
   | left=expr op=('+'|'-') right=expr
   | NUMBER
   ;
