### gc 写屏障

###示例代码
```go
package main

type Test struct {
	Val *int
}

var a,b Test

func main() {
	var c int
	simple(&c)
}

func simple(c *int)  {
	b.Val = c
	a.Val = b.Val
	b.Val = nil
}

```

    go build -gcflags '-N -l' write_barrier.go
    go tool objdump -s 'main\.simple' write_barrier
    
    0x10515e0               65488b0c2530000000      MOVQ GS:0x30, CX                        
      0x10515e9               483b6110                CMPQ 0x10(CX), SP                       
      0x10515ed               7668                    JBE 0x1051657                           
      0x10515ef               4883ec08                SUBQ $0x8, SP                           
      0x10515f3               48892c24                MOVQ BP, 0(SP)                          
      0x10515f7               488d2c24                LEAQ 0(SP), BP                          
      0x10515fb               488b442410              MOVQ 0x10(SP), AX                       
      0x1051600               833d19d8080000          CMPL $0x0, runtime.writeBarrier(SB)     
      0x1051607               7402                    JE 0x105160b                            
      0x1051609               eb24                    JMP 0x105162f                           
      0x105160b               4889051e220700          MOVQ AX, main.b(SB)                     
      0x1051612               4889050f220700          MOVQ AX, main.a(SB)                     
      0x1051619               48c7050c22070000000000  MOVQ $0x0, main.b(SB)                   
      0x1051624               eb00                    JMP 0x1051626                           
      0x1051626               488b2c24                MOVQ 0(SP), BP                          
      0x105162a               4883c408                ADDQ $0x8, SP                           
      0x105162e               c3                      RET                                     
      0x105162f               488d3dfa210700          LEAQ main.b(SB), DI                     
      0x1051636               e8a59effff              CALL runtime.gcWriteBarrier(SB)         
      0x105163b               488d3de6210700          LEAQ main.a(SB), DI                     
      0x1051642               e8999effff              CALL runtime.gcWriteBarrier(SB)         
      0x1051647               488d3de2210700          LEAQ main.b(SB), DI                     
      0x105164e               31c0                    XORL AX, AX                             
      0x1051650               e88b9effff              CALL runtime.gcWriteBarrier(SB)         
      0x1051655               ebcf                    JMP 0x1051626                           
      0x1051657               e8a480ffff              CALL runtime.morestack_noctxt(SB)       
      0x105165c               eb82                    JMP main.simple(SB)       


  