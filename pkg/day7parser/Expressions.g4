grammar Expressions;

// Parser for lines like:
// bright white bags contain 1 shiny gold bag.
// vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
// faded blue bags contain no other bags.

// this turns into
// [Parser].TreeStart() method returning antlr.ParseTree
treeStart : expr EOF;

expr: label? K_BAGS K_CONTAIN component+ ;
component
    : (K_NO K_OTHER | count=NUMERIC_LITERAL label?) (K_BAGS|K_BAG) (DOT|COMMA)
    ;
label: ATOM+ ;

K_NO: N O ;
K_OTHER: O T H E R ;
K_BAG: B A G ;
K_BAGS: B A G S ;
K_CONTAIN: C O N T A I N ;

NUMERIC_LITERAL: DIGIT+ ( '.' DIGIT* )?;

DOT: '.';
COMMA: ',';

ATOM
 : [a-zA-Z][a-zA-Z0-9]*
 ;

WHITESPACE: [ \r\n\t]+ -> skip;

fragment DIGIT : [0-9];

fragment A : [aA];
fragment B : [bB];
fragment C : [cC];
fragment D : [dD];
fragment E : [eE];
fragment F : [fF];
fragment G : [gG];
fragment H : [hH];
fragment I : [iI];
fragment J : [jJ];
fragment K : [kK];
fragment L : [lL];
fragment M : [mM];
fragment N : [nN];
fragment O : [oO];
fragment P : [pP];
fragment Q : [qQ];
fragment R : [rR];
fragment S : [sS];
fragment T : [tT];
fragment U : [uU];
fragment V : [vV];
fragment W : [wW];
fragment X : [xX];
fragment Y : [yY];
fragment Z : [zZ];
