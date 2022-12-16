### 原子加法
```go
uint32 volatile *val, int32 delta
 
TEXT ·Xadd(SB), NOSPLIT, $0-20
    MOVQ	ptr+0(FP), BX  // 8字节
	MOVL	delta+8(FP), AX   // 4字节
	MOVL	AX, CX
	LOCK
	XADDL	AX, 0(BX)
	ADDL	CX, AX
	MOVL	AX, ret+16(FP)
	RET
```

---
AX：寄存器称为累加器，常用于存放算术、逻辑运算中的操作数或结果。另外，所有的I/O指令都要使用累加器与外设接口传递数据。

BX：寄存器称为基址寄存器，常用来存放访问内存时的地址。

CX：寄存器称为计数寄存器，在循环、串操作指令中用作计数器。

DX：寄存器称为数据寄存器，在寄存器间接寻址中的I/O指令中存放I/O端口的地址。

--- 
参考文献
Plan9操作系统：https://plan9.io/plan9/

Plan9汇编器手册：http://doc.cat-v.org/plan_9/4th_edition/papers/asm

https://golang.design/under-the-hood/zh-cn/part1basic/ch01basic/

cas https://segmentfault.com/a/1190000039785918