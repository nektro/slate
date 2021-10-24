```
$ llc-12 -march=aarch64 -mattr=help
```
```
Available CPUs for this target:

  a64fx         - Select the a64fx processor.
  apple-a10     - Select the apple-a10 processor.
  apple-a11     - Select the apple-a11 processor.
  apple-a12     - Select the apple-a12 processor.
  apple-a13     - Select the apple-a13 processor.
  apple-a14     - Select the apple-a14 processor.
  apple-a7      - Select the apple-a7 processor.
  apple-a8      - Select the apple-a8 processor.
  apple-a9      - Select the apple-a9 processor.
  apple-latest  - Select the apple-latest processor.
  apple-s4      - Select the apple-s4 processor.
  apple-s5      - Select the apple-s5 processor.
  carmel        - Select the carmel processor.
  cortex-a34    - Select the cortex-a34 processor.
  cortex-a35    - Select the cortex-a35 processor.
  cortex-a53    - Select the cortex-a53 processor.
  cortex-a55    - Select the cortex-a55 processor.
  cortex-a57    - Select the cortex-a57 processor.
  cortex-a65    - Select the cortex-a65 processor.
  cortex-a65ae  - Select the cortex-a65ae processor.
  cortex-a72    - Select the cortex-a72 processor.
  cortex-a73    - Select the cortex-a73 processor.
  cortex-a75    - Select the cortex-a75 processor.
  cortex-a76    - Select the cortex-a76 processor.
  cortex-a76ae  - Select the cortex-a76ae processor.
  cortex-a77    - Select the cortex-a77 processor.
  cortex-a78    - Select the cortex-a78 processor.
  cortex-a78c   - Select the cortex-a78c processor.
  cortex-r82    - Select the cortex-r82 processor.
  cortex-x1     - Select the cortex-x1 processor.
  cyclone       - Select the cyclone processor.
  exynos-m3     - Select the exynos-m3 processor.
  exynos-m4     - Select the exynos-m4 processor.
  exynos-m5     - Select the exynos-m5 processor.
  falkor        - Select the falkor processor.
  generic       - Select the generic processor.
  kryo          - Select the kryo processor.
  neoverse-e1   - Select the neoverse-e1 processor.
  neoverse-n1   - Select the neoverse-n1 processor.
  neoverse-n2   - Select the neoverse-n2 processor.
  neoverse-v1   - Select the neoverse-v1 processor.
  saphira       - Select the saphira processor.
  thunderx      - Select the thunderx processor.
  thunderx2t99  - Select the thunderx2t99 processor.
  thunderx3t110 - Select the thunderx3t110 processor.
  thunderxt81   - Select the thunderxt81 processor.
  thunderxt83   - Select the thunderxt83 processor.
  thunderxt88   - Select the thunderxt88 processor.
  tsv110        - Select the tsv110 processor.

Available features for this target:

  CONTEXTIDREL2                      - Enable RW operand CONTEXTIDR_EL2.
  a35                                - Cortex-A35 ARM processors.
  a53                                - Cortex-A53 ARM processors.
  a55                                - Cortex-A55 ARM processors.
  a57                                - Cortex-A57 ARM processors.
  a64fx                              - Fujitsu A64FX processors.
  a65                                - Cortex-A65 ARM processors.
  a72                                - Cortex-A72 ARM processors.
  a73                                - Cortex-A73 ARM processors.
  a75                                - Cortex-A75 ARM processors.
  a76                                - Cortex-A76 ARM processors.
  a77                                - Cortex-A77 ARM processors.
  aes                                - Enable AES support.
  aggressive-fma                     - Enable Aggressive FMA for floating-point..
  alternate-sextload-cvt-f32-pattern - Use alternative pattern for sextload convert to f32.
  altnzcv                            - Enable alternative NZCV format for floating point comparisons.
  am                                 - Enable v8.4-A Activity Monitors extension.
  amvs                               - Enable v8.6-A Activity Monitors Virtualization support.
  apple-a10                          - Apple A10.
  apple-a11                          - Apple A11.
  apple-a12                          - Apple A12.
  apple-a13                          - Apple A13.
  apple-a14                          - Apple A14.
  apple-a7                           - Apple A7 (the CPU formerly known as Cyclone).
  arith-bcc-fusion                   - CPU fuses arithmetic+bcc operations.
  arith-cbz-fusion                   - CPU fuses arithmetic + cbz/cbnz operations.
  balance-fp-ops                     - balance mix of odd and even D-registers for fp multiply(-accumulate) ops.
  bf16                               - Enable BFloat16 Extension.
  brbe                               - Enable Branch Record Buffer Extension.
  bti                                - Enable Branch Target Identification.
  call-saved-x10                     - Make X10 callee saved..
  call-saved-x11                     - Make X11 callee saved..
  call-saved-x12                     - Make X12 callee saved..
  call-saved-x13                     - Make X13 callee saved..
  call-saved-x14                     - Make X14 callee saved..
  call-saved-x15                     - Make X15 callee saved..
  call-saved-x18                     - Make X18 callee saved..
  call-saved-x8                      - Make X8 callee saved..
  call-saved-x9                      - Make X9 callee saved..
  carmel                             - Nvidia Carmel processors.
  ccdp                               - Enable v8.5 Cache Clean to Point of Deep Persistence.
  ccidx                              - Enable v8.3-A Extend of the CCSIDR number of sets.
  ccpp                               - Enable v8.2 data Cache Clean to Point of Persistence.
  cmp-bcc-fusion                     - CPU fuses cmp+bcc operations.
  complxnum                          - Enable v8.3-A Floating-point complex number support.
  cortex-a78                         - Cortex-A78 ARM processors.
  cortex-a78c                        - Cortex-A78C ARM processors.
  cortex-r82                         - Cortex-R82 ARM Processors.
  cortex-x1                          - Cortex-X1 ARM processors.
  crc                                - Enable ARMv8 CRC-32 checksum instructions.
  crypto                             - Enable cryptographic instructions.
  custom-cheap-as-move               - Use custom handling of cheap instructions.
  disable-latency-sched-heuristic    - Disable latency scheduling heuristic.
  dit                                - Enable v8.4-A Data Independent Timing instructions.
  dotprod                            - Enable dot product support.
  ecv                                - Enable enhanced counter virtualization extension.
  ete                                - Enable Embedded Trace Extension.
  exynos-cheap-as-move               - Use Exynos specific handling of cheap instructions.
  exynosm3                           - Samsung Exynos-M3 processors.
  exynosm4                           - Samsung Exynos-M4 processors.
  f32mm                              - Enable Matrix Multiply FP32 Extension.
  f64mm                              - Enable Matrix Multiply FP64 Extension.
  falkor                             - Qualcomm Falkor processors.
  fgt                                - Enable fine grained virtualization traps extension.
  flagm                              - Enable v8.4-A Flag Manipulation Instructions.
  force-32bit-jump-tables            - Force jump table entries to be 32-bits wide except at MinSize.
  fp-armv8                           - Enable ARMv8 FP.
  fp16fml                            - Enable FP16 FML instructions.
  fptoint                            - Enable FRInt[32|64][Z|X] instructions that round a floating-point number to an integer (in FP format) forcing it to fit into a 32- or 64-bit int.
  fullfp16                           - Full FP16.
  fuse-address                       - CPU fuses address generation and memory operations.
  fuse-aes                           - CPU fuses AES crypto operations.
  fuse-arith-logic                   - CPU fuses arithmetic and logic operations.
  fuse-crypto-eor                    - CPU fuses AES/PMULL and EOR operations.
  fuse-csel                          - CPU fuses conditional select operations.
  fuse-literals                      - CPU fuses literal generation operations.
  harden-sls-blr                     - Harden against straight line speculation across BLR instructions.
  harden-sls-retbr                   - Harden against straight line speculation across RET and BR instructions.
  hcx                                - Enable Armv8.7-A HCRX_EL2 system register.
  i8mm                               - Enable Matrix Multiply Int8 Extension.
  jsconv                             - Enable v8.3-A JavaScript FP conversion instructions.
  kryo                               - Qualcomm Kryo processors.
  lor                                - Enables ARM v8.1 Limited Ordering Regions extension.
  ls64                               - Enable Armv8.7-A LD64B/ST64B Accelerator Extension.
  lse                                - Enable ARMv8.1 Large System Extension (LSE) atomic instructions.
  lsl-fast                           - CPU has a fastpath logical shift of up to 3 places.
  mpam                               - Enable v8.4-A Memory system Partitioning and Monitoring extension.
  mte                                - Enable Memory Tagging Extension.
  neon                               - Enable Advanced SIMD instructions.
  neoversee1                         - Neoverse E1 ARM processors.
  neoversen1                         - Neoverse N1 ARM processors.
  neoversen2                         - Neoverse N2 ARM processors.
  neoversev1                         - Neoverse V1 ARM processors.
  no-neg-immediates                  - Convert immediates and instructions to their negated or complemented equivalent when the immediate does not fit in the encoding..
  nv                                 - Enable v8.4-A Nested Virtualization Enchancement.
  outline-atomics                    - Enable out of line atomics to support LSE instructions.
  pan                                - Enables ARM v8.1 Privileged Access-Never extension.
  pan-rwv                            - Enable v8.2 PAN s1e1R and s1e1W Variants.
  pauth                              - Enable v8.3-A Pointer Authentication extension.
  perfmon                            - Enable ARMv8 PMUv3 Performance Monitors extension.
  pmu                                - Enable v8.4-A PMU extension.
  predictable-select-expensive       - Prefer likely predicted branches over selects.
  predres                            - Enable v8.5a execution and data prediction invalidation instructions.
  rand                               - Enable Random Number generation instructions.
  ras                                - Enable ARMv8 Reliability, Availability and Serviceability Extensions.
  rcpc                               - Enable support for RCPC extension.
  rcpc-immo                          - Enable v8.4-A RCPC instructions with Immediate Offsets.
  rdm                                - Enable ARMv8.1 Rounding Double Multiply Add/Subtract instructions.
  reserve-x1                         - Reserve X1, making it unavailable as a GPR.
  reserve-x10                        - Reserve X10, making it unavailable as a GPR.
  reserve-x11                        - Reserve X11, making it unavailable as a GPR.
  reserve-x12                        - Reserve X12, making it unavailable as a GPR.
  reserve-x13                        - Reserve X13, making it unavailable as a GPR.
  reserve-x14                        - Reserve X14, making it unavailable as a GPR.
  reserve-x15                        - Reserve X15, making it unavailable as a GPR.
  reserve-x18                        - Reserve X18, making it unavailable as a GPR.
  reserve-x2                         - Reserve X2, making it unavailable as a GPR.
  reserve-x20                        - Reserve X20, making it unavailable as a GPR.
  reserve-x21                        - Reserve X21, making it unavailable as a GPR.
  reserve-x22                        - Reserve X22, making it unavailable as a GPR.
  reserve-x23                        - Reserve X23, making it unavailable as a GPR.
  reserve-x24                        - Reserve X24, making it unavailable as a GPR.
  reserve-x25                        - Reserve X25, making it unavailable as a GPR.
  reserve-x26                        - Reserve X26, making it unavailable as a GPR.
  reserve-x27                        - Reserve X27, making it unavailable as a GPR.
  reserve-x28                        - Reserve X28, making it unavailable as a GPR.
  reserve-x3                         - Reserve X3, making it unavailable as a GPR.
  reserve-x30                        - Reserve X30, making it unavailable as a GPR.
  reserve-x4                         - Reserve X4, making it unavailable as a GPR.
  reserve-x5                         - Reserve X5, making it unavailable as a GPR.
  reserve-x6                         - Reserve X6, making it unavailable as a GPR.
  reserve-x7                         - Reserve X7, making it unavailable as a GPR.
  reserve-x9                         - Reserve X9, making it unavailable as a GPR.
  saphira                            - Qualcomm Saphira processors.
  sb                                 - Enable v8.5 Speculation Barrier.
  sel2                               - Enable v8.4-A Secure Exception Level 2 extension.
  sha2                               - Enable SHA1 and SHA256 support.
  sha3                               - Enable SHA512 and SHA3 support.
  slow-misaligned-128store           - Misaligned 128 bit stores are slow.
  slow-paired-128                    - Paired 128 bit loads and stores are slow.
  slow-strqro-store                  - STR of Q register with register offset is slow.
  sm4                                - Enable SM3 and SM4 support.
  spe                                - Enable Statistical Profiling extension.
  spe-eef                            - Enable extra register in the Statistical Profiling Extension.
  specrestrict                       - Enable architectural speculation restriction.
  ssbs                               - Enable Speculative Store Bypass Safe bit.
  strict-align                       - Disallow all unaligned memory access.
  sve                                - Enable Scalable Vector Extension (SVE) instructions.
  sve2                               - Enable Scalable Vector Extension 2 (SVE2) instructions.
  sve2-aes                           - Enable AES SVE2 instructions.
  sve2-bitperm                       - Enable bit permutation SVE2 instructions.
  sve2-sha3                          - Enable SHA3 SVE2 instructions.
  sve2-sm4                           - Enable SM4 SVE2 instructions.
  tagged-globals                     - Use an instruction sequence for taking the address of a global that allows a memory tag in the upper address bits.
  thunderx                           - Cavium ThunderX processors.
  thunderx2t99                       - Cavium ThunderX2 processors.
  thunderx3t110                      - Marvell ThunderX3 processors.
  thunderxt81                        - Cavium ThunderX processors.
  thunderxt83                        - Cavium ThunderX processors.
  thunderxt88                        - Cavium ThunderX processors.
  tlb-rmi                            - Enable v8.4-A TLB Range and Maintenance Instructions.
  tme                                - Enable Transactional Memory Extension.
  tpidr-el1                          - Permit use of TPIDR_EL1 for the TLS base.
  tpidr-el2                          - Permit use of TPIDR_EL2 for the TLS base.
  tpidr-el3                          - Permit use of TPIDR_EL3 for the TLS base.
  tracev8.4                          - Enable v8.4-A Trace extension.
  trbe                               - Enable Trace Buffer Extension.
  tsv110                             - HiSilicon TS-V110 processors.
  uaops                              - Enable v8.2 UAO PState.
  use-aa                             - Use alias analysis during codegen.
  use-experimental-zeroing-pseudos   - Hint to the compiler that the MOVPRFX instruction is merged with destructive operations.
  use-postra-scheduler               - Schedule again after register allocation.
  use-reciprocal-square-root         - Use the reciprocal square root approximation.
  v8.1a                              - Support ARM v8.1a instructions.
  v8.2a                              - Support ARM v8.2a instructions.
  v8.3a                              - Support ARM v8.3a instructions.
  v8.4a                              - Support ARM v8.4a instructions.
  v8.5a                              - Support ARM v8.5a instructions.
  v8.6a                              - Support ARM v8.6a instructions.
  v8.7a                              - Support ARM v8.7a instructions.
  v8r                                - Support ARM v8r instructions.
  vh                                 - Enables ARM v8.1 Virtual Host extension.
  wfxt                               - Enable Armv8.7-A WFET and WFIT instruction.
  xs                                 - Enable Armv8.7-A limited-TLB-maintenance instruction.
  zcm                                - Has zero-cycle register moves.
  zcz                                - Has zero-cycle zeroing instructions.
  zcz-fp                             - Has zero-cycle zeroing instructions for FP registers.
  zcz-fp-workaround                  - The zero-cycle floating-point zeroing instruction has a bug.
  zcz-gp                             - Has zero-cycle zeroing instructions for generic registers.

Use +feature to enable a feature, or -feature to disable it.
For example, llc -mcpu=mycpu -mattr=+feature1,-feature2
