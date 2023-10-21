  .globl main
main:
  mov $4, %rax
  push %rax
  mov $20, %rax
  push %rax
  mov $5, %rax
  pop %rdi
  add %rdi, %rax
  pop %rdi
  sub %rdi, %rax
  ret
