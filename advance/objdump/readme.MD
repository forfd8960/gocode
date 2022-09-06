
## objdump

```sh

go build main.go; and go tool objdump ./main | grep -E "main.go:4|main.go:7|main.go:10"

  main.go:4  0x1053f98  488d05e1440000  LEAQ type.*+17184(SB), AX  
  main.go:4  0x1053f9f  bba0860100  MOVL $0x186a0, BX   
  main.go:4  0x1053fa4  4889d9   MOVQ BX, CX    
  main.go:4  0x1053fa7  e8b4b2feff  CALL runtime.makeslice(SB)  
  main.go:4  0x1053fac  4889442428  MOVQ AX, 0x28(SP)   
  main.go:7  0x1053fd2  488d05e73e0000  LEAQ type.*+15712(SB), AX  
  main.go:7  0x1053fd9  bb05000000  MOVL $0x5, BX    
  main.go:7  0x1053fde  6690   NOPW     
  main.go:7  0x1053fe0  e89bf4faff  CALL runtime.makechan(SB)  
  main.go:7  0x1053fe5  4889442420  MOVQ AX, 0x20(SP)   
  main.go:10  0x1054005  440f117c2430  MOVUPS X15, 0x30(SP)   
  main.go:10  0x105400b  440f117c2440  MOVUPS X15, 0x40(SP)   
  main.go:10  0x1054011  440f117c2450  MOVUPS X15, 0x50(SP)   
  main.go:10  0x1054017  488d0582590000  LEAQ type.*+22592(SB), AX  
  main.go:10  0x105401e  bb0a000000  MOVL $0xa, BX    
  main.go:10  0x1054023  488d4c2430  LEAQ 0x30(SP), CX   
  main.go:10  0x1054028  e85377fbff  CALL runtime.makemap(SB)  
  main.go:10  0x105402d  4889442418  MOVQ AX, 0x18(SP)


```