## 生成 .o 目标文件，并把目标的汇编内容输出


```sh
go tool compile -S ./main.go |grep "main.go:7"

 0x0014 00020 (./main.go:7) XORL AX, AX
 0x0016 00022 (./main.go:7) LEAQ go.string."hello"(SB), BX
 0x001d 00029 (./main.go:7) MOVL $5, CX
 0x0022 00034 (./main.go:7) PCDATA $1, $0
 0x0022 00034 (./main.go:7) CALL runtime.stringtoslicebyte(SB)
```