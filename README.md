DNA(脱氧核糖核酸)碱基有4种：腺嘌呤（A）、鸟嘌呤（G）、胸腺嘧啶（T）和胞嘧啶（C）。

RNA(核糖核酸)碱基主要有4种，即A（腺嘌呤）、G（鸟嘌呤）、C（胞嘧啶)、U(尿嘧啶)，其中，U（尿嘧啶）取代了DNA中的T。

# BaseCode 密码本加密

BaseCode 是一个密码本加密包，可以很方便创建自己的密码本。默认使用RNA碱基对（因为DNA是双链🤔）作为密码本

你说啥？百家姓，价值观...    什么莫名其妙的,哼！拍桌.jpg


好孩子看不懂：GCUGGCAGGCGUGCUCGCGGGUGAAUCCAUUUGUCAGUGAAUUGGUGGGUACGCUCAUCCGCACGUGAGCCGGCCAAUCCGCAGGCGCAUGAAUAAAUACAUGUAUGCAUCGGCGAAUAGGCGCAUACAUAUAUAGAUCAGCACGCGAGCGCAUGGAUAAAUGAAUGUGCAUAUCAGCAUAUAAGCAUGCGCAUAGGCACAUGAAUGCGCGGAUCGAUACAUCGAUGGAUCAGCAGGCAU


## Installation
```
go get github.com/Sakurasan/basecode
```

Import it in your code:
```
import "github.com/Sakurasan/basecode"
```

## Quick start
```
$ cat basecode_test.go  | 👀 👇
```
```
package main

import (
	"github.com/Sakurasan/basecode"
	"fmt"
)

func main() {
	fmt.Println(basecode.StdEncoding.Encode("DNA"))
}
```


## BenchMark

![testimg.png](https://i.loli.net/2021/05/14/31HQjW786NJMPFs.png)