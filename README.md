# Hyperloglog Counter
We know that Hyperloglog is powerfull for cardinality count. We can estimate how many unique user that visit home page. But Hyperloglog cannot handle how many times user X visit page Home. This counter will solve that problem.

# How to Use?
```go
package main

import (
	"fmt"

	"github.com/ariefsam/hllcounter"
)

func main() {
	counter := hllcounter.NewCounter()
	b := counter.Add("User 4569 visit home")
	fmt.Println(b)
	b = counter.Add("User 4569 visit home")
	fmt.Println(b)
	b = counter.Add("User 4569 visit home")
	fmt.Println(b)
	b = counter.Add("User 888 visit home")
	fmt.Println(b)
	b = counter.HowMany("User 888 visit home")
	fmt.Println(b)
	b = counter.HowMany("User 4569 visit home")
	fmt.Println(b)
}

```