## Kurulum

### Gereklilikler

- **[Go](https://go.dev/)**'nun bilgisayarınızda yüklü olması.

### Paketi projenize entegre etme

```bash
$ go get github.com/yusuftalhaklc/tc-kimlik-dogrulama
```

### Kullanım

```go
package main

import (
	"fmt"

	tckimlik "github.com/yusuftalhaklc/tc-kimlik-dogrulama"
	"github.com/yusuftalhaklc/tc-kimlik-dogrulama/model"
)

func main() {
	identity := &model.TCKimlik{
		TCKimlikNo: 11111111111,
		Ad:         "Yusuf Talha",
		Soyad:      "Kılıç",
		DogumYili:  1111,
	}

	verify := tckimlik.Dogrula(identity) // returns bool
	fmt.Println(verify)
}
```
```bash
$ go run example.go
true
```
