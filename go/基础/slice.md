slice底层数据结构是一个结构体, array存储数组指针，这样在传递参数时避免（大数据量）的值传递

在修改切片值会直接修改数组，导致原始切片值改变

append时会copy一份，不会更改原始切片

切片截取数组也使用同一块内存地址，即使赋予新变量也会更改原始值

切片容量是指原始数组长度

```go
type slice struct {
    array unsafe.Pointer 
    len int 
    cap int
}
```


```go

package main

import "fmt"

func main() {
    a := []int{1, 2, 3, 4, 5, 6}
    setv(a)
    fmt.Println(a)
}

func setv(val []int) {
    val = append(val, 7)
    // val[0] = 2
    fmt.Println(val)
}
```

扩容 append slice
```shell

$ dlv debug main.go
Type 'help' for list of commands.
(dlv) break main.go:7
Breakpoint 1 set at 0x49cc5c for main.main() ./main.go:7
(dlv) break main.go:8
Breakpoint 2 set at 0x49cc69 for main.main() ./main.go:8
(dlv) break main.go:13
Breakpoint 3 set at 0x49cd96 for main.setv() ./main.go:13
(dlv) continue
> main.main() ./main.go:7 (hits goroutine(1):1 total:1) (PC: 0x49cc5c)
     2:
     3: import "fmt"
     4:
     5: func main() {
     6:         a := []int{1, 2, 3, 4, 5, 6}
=>   7:         setv(a)
     8:         fmt.Println(a)
     9: }
    10:
    11: func setv(val []int) {
    12:         val = append(val, 7)
(dlv) print a 
[]int len: 6, cap: 6, [1,2,3,4,5,6]
(dlv) examinemem -count 3 -size 8 -x &a
0xc00006ff40:   0x000000c00001c0c0   0x0000000000000006   0x0000000000000006   
(dlv) examinemem -count 6 -size 8 -x 0x000000c00001c0c0
0xc00001c0c0:   0x0000000000000001   0x0000000000000002   0x0000000000000003   0x0000000000000004   0x0000000000000005   0x0000000000000006   




(dlv) continue
> main.setv() ./main.go:13 (hits goroutine(1):1 total:1) (PC: 0x49cd96)
     8:         fmt.Println(a)
     9: }
    10:
    11: func setv(val []int) {
    12:         val = append(val, 7)
=>  13:         fmt.Println(val)
    14: }
(dlv) print val
[]int len: 7, cap: 12, [1,2,3,4,5,6,7]
(dlv) examinemem -count 3 -size 8 -x &val
0xc00006ff00:   0x000000c00008a060   0x0000000000000007   0x000000000000000c   
   
(dlv) examinemem -count 7 -size 8 -x 0x000000c00008a060
0xc00008a060:   0x0000000000000001   0x0000000000000002   0x0000000000000003   0x0000000000000004   0x0000000000000005   0x0000000000000006   0x0000000000000007  

```



修改slice值

```shell

$ dlv debug main.go                           
Type 'help' for list of commands.
(dlv) break main.go:7
Breakpoint 1 set at 0x49cc5c for main.main() ./main.go:7
(dlv) break main.go:8
Breakpoint 2 set at 0x49cc69 for main.main() ./main.go:8
(dlv) break main.go:14
Breakpoint 3 set at 0x49cd58 for main.setv() ./main.go:14
(dlv) continue
> main.main() ./main.go:7 (hits goroutine(1):1 total:1) (PC: 0x49cc5c)
     2:
     3: import "fmt"
     4:
     5: func main() {
     6:         a := []int{1, 2, 3, 4, 5, 6}
=>   7:         setv(a)
     8:         fmt.Println(a)
     9: }
    10:
    11: func setv(val []int) {
    12:         // val = append(val, 7)
(dlv) print a 
[]int len: 6, cap: 6, [1,2,3,4,5,6]
(dlv) examinemem -count 3 -size 8 -x &a
0xc00006ff40:   0x000000c00001c0c0   0x0000000000000006   0x0000000000000006   
 
(dlv) examinemem -count 7 -size 8 -x 0x000000c00001c0c0
0xc00001c0c0:   0x0000000000000001   0x0000000000000002   0x0000000000000003   0x0000000000000004   0x0000000000000005   0x0000000000000006   0x0000000000000000   
(dlv) continue
> main.setv() ./main.go:14 (hits goroutine(1):1 total:1) (PC: 0x49cd58)
     9: }
    10:
    11: func setv(val []int) {
    12:         // val = append(val, 7)
    13:         val[0] = 2
=>  14:         fmt.Println(val)
    15: }
(dlv) print val
[]int len: 6, cap: 6, [2,2,3,4,5,6]
(dlv) examinemem -count 3 -size 8 -x &val
0xc00006ff00:   0x000000c00001c0c0   0x0000000000000006   0x0000000000000006   
(dlv) examinemem -count 7 -size 8 -x 0x000000c00001c0c0
0xc00001c0c0:   0x0000000000000002   0x0000000000000002   0x0000000000000003   0x0000000000000004   0x0000000000000005   0x0000000000000006   0x0000000000000000   
(dlv) continue
[2 2 3 4 5 6]
> main.main() ./main.go:8 (hits goroutine(1):1 total:1) (PC: 0x49cc69)
     3: import "fmt"
     4:
     5: func main() {
     6:         a := []int{1, 2, 3, 4, 5, 6}
     7:         setv(a)
=>   8:         fmt.Println(a)
     9: }
    10:
    11: func setv(val []int) {
    12:         // val = append(val, 7)
    13:         val[0] = 2
(dlv) print a
[]int len: 6, cap: 6, [2,2,3,4,5,6]
(dlv) examinemem -count 3 -size 8 -x &a
0xc00006ff40:   0x000000c00001c0c0   0x0000000000000006   0x0000000000000006   
  
(dlv) examinemem -count 7 -size 8 -x 0x000000c00001c0c0
0xc00001c0c0:   0x0000000000000002   0x0000000000000002   0x0000000000000003   0x0000000000000004   0x0000000000000005   0x0000000000000006   0x0000000000000000   

```


copy一份
```shell
$ dlv debug main.go
Type 'help' for list of commands.
(dlv) break main.go:12
Breakpoint 1 set at 0x49cd13 for main.main() ./main.go:12
(dlv) continue
> main.main() ./main.go:12 (hits goroutine(1):1 total:1) (PC: 0x49cd13)
     7:         // setv(a)
     8:         // fmt.Println(a)
     9:         // fmt.Println(cap(a))
    10:         var a1 = make([]int, 3)
    11:         copy(a1, a)
=>  12:         fmt.Println(a1, a)
    13:
    14: }
    15:
    16: func setv(val []int) {
    17:         val = append(val, 7)
(dlv) print a
[]int len: 3, cap: 3, [1,2,3]
(dlv) print a1
[]int len: 3, cap: 3, [1,2,3]
(dlv) examinemem -count 3 -size 8 -x &a
0xc00006fef0:   0x000000c00001e0c0   0x0000000000000003   0x0000000000000003   
(dlv) examinemem -count 3 -size 8 -x &a1
0xc00006fed8:   0x000000c00001e0d8   0x0000000000000003   0x0000000000000003   
(dlv) examinemem -count 3 -size 8 -x 0x000000c00001e0c0
0xc00001e0c0:   0x0000000000000001   0x0000000000000002   0x0000000000000003   
(dlv) examinemem -count 3 -size 8 -x 0x000000c00001e0d8
0xc00001e0d8:   0x0000000000000001   0x0000000000000002   0x0000000000000003
```

