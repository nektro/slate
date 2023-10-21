  .globl main
main:
  mov $6, %rax
  push %rax
  mov $9, %rax
  pop %rdi
  sub %rdi, %rax
  push %rax
  mov $5, %rax
  pop %rdi
  imul %rdi, %rax
  ret
