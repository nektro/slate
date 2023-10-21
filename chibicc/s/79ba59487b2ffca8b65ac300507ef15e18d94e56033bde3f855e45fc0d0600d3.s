  .globl main
main:
  mov $5, %rax
  push %rax
  mov $34, %rax
  push %rax
  mov $12, %rax
  pop %rdi
  add %rdi, %rax
  pop %rdi
  sub %rdi, %rax
  ret
