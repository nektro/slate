  .globl main
main:
  mov $7, %rax
  push %rax
  mov $6, %rax
  pop %rdi
  imul %rdi, %rax
  push %rax
  mov $5, %rax
  pop %rdi
  add %rdi, %rax
  ret
