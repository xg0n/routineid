# routineid

`routineid` help you to get ids of the current go routine and its parent go routine, using implementation from library [gohack](https://github.com/timandy/gohack)

## Usage
The library only exposes `func GetRoutineIds() (curRoutineId uint64, parentRoutineId uint64)`

**Example**
`go_routine_ids_example.go`
```go
package main

import (
	"fmt"
	"sync"

	"github.com/xg0n/routineid"
)

var wg sync.WaitGroup

func printGoRoutineId(routineName string) {
	curId, parentId := routineid.GetRoutineIds()
	fmt.Printf("%4s go routine - current id: %2d, parent id: %3d\n", routineName, curId, parentId)
}

func main() {
	printGoRoutineId("main")

	wg.Add(1)
	go foo()

	/* Wait for all go routines complete */
	wg.Wait()
}

func foo() {
	defer wg.Done()
	printGoRoutineId("foo")

	wg.Add(1)
	go bar()
}

func bar() {
	defer wg.Done()

	printGoRoutineId("bar")
}
```

**Result**:

```shell
$ go run go_routine_ids_example.go
main go routine - current id:  1, parent id:   0
 foo go routine - current id: 18, parent id:   1
 bar go routine - current id: 19, parent id:  18
```

# Support Grid

|                | **`darwin`** | **`linux`** | **`windows`** | **`freebsd`** | **`js`** |                |
|---------------:|:------------:|:-----------:|:-------------:|:-------------:|:--------:|:---------------|
|      **`386`** |              |      ✅      |       ✅       |       ✅       |          | **`386`**      |
|    **`amd64`** |      ✅       |      ✅      |       ✅       |       ✅       |          | **`amd64`**    |
|    **`armv6`** |              |      ✅      |               |               |          | **`armv6`**    |
|    **`armv7`** |              |      ✅      |               |               |          | **`armv7`**    |
|    **`arm64`** |      ✅       |      ✅      |               |               |          | **`arm64`**    |
|  **`loong64`** |              |      ✅      |               |               |          | **`loong64`**  |
|     **`mips`** |              |      ✅      |               |               |          | **`mips`**     |
|   **`mipsle`** |              |      ✅      |               |               |          | **`mipsle`**   |
|   **`mips64`** |              |      ✅      |               |               |          | **`mips64`**   |
| **`mips64le`** |              |      ✅      |               |               |          | **`mips64le`** |
|    **`ppc64`** |              |      ✅      |               |               |          | **`ppc64`**    |
|  **`ppc64le`** |              |      ✅      |               |               |          | **`ppc64le`**  |
|  **`riscv64`** |              |      ✅      |               |               |          | **`riscv64`**  |
|    **`s390x`** |              |      ✅      |               |               |          | **`s390x`**    |
|     **`wasm`** |              |             |               |               |    ✅     | **`wasm`**     |
|                | **`darwin`** | **`linux`** | **`windows`** | **`freebsd`** | **`js`** |                |

✅: Supported

# *License*

`gohack` is released under the [Apache License 2.0](LICENSE).

```
Copyright 2021-2024 TimAndy

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
