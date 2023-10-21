  .globl main
main:
  mov $2, %rax
  push %rax
  mov $5, %rax
  push %rax
  mov $3, %rax
  pop %rdi
  add %rdi, %rax
  pop %rdi
  cqo
  idiv %rdi
  ret
