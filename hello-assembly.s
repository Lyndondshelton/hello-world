.data
string: .asciiz "Hello World"

.text
.globl main
main:

#print the string
la $a0, string
li $v0, 4
syscall

#end the program
li $v0, 10
syscall
