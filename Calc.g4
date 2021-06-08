grammar Calc;

// Tokens
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
LEFT: '(';
RIGHT: ')';
NUMBER: [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;

// Rules
start : expression EOF;

expression
   : left=expression op=('*'|'/') right=expression # MulDiv
   | left=expression op=('+'|'-') right=expression # AddSub
   | NUMBER                             # Number
   | LEFT expression RIGHT              # Parenthesis;
