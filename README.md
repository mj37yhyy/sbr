# sbr
Golang Struct By Reflect

```go
package sbr

import (
	"fmt"
	"testing"
)

type Children struct {
	Name string
	Age  int
	Sex  bool
}

func TestStructByReflect(t *testing.T) {
	sbr := Sbr{}

	child := Children{}
	m := map[string]any{"Name": "Tom", "Age": 15, "Sex": true}
	sbr.StructByReflect(m, &child)

	fmt.Printf("child: %v", child)
}
```
run this program:
```shell
=== RUN   TestStructByReflect
child: {Tom 15 true}--- PASS: TestStructByReflect (0.00s)
PASS
```