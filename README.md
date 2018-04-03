# Hyperloglog Counter
We know that Hyperloglog is powerfull for cardinality count. We can estimate how many unique user that visit home page. But Hyperloglog cannot handle how many times user X visit page Home. This counter will solve that problem.

# How to Use?

For Add new action, use Add

For counting, use HowMany

Function HowMany is more expensive than Add, use it wisely

```go
package main

import (
	"fmt"

	"github.com/ariefsam/hllcounter"
)

func main() {
	counter := hllcounter.NewCounter()
	b := counter.Add("User 4569 visit home")
	fmt.Println(b) // 1
	b = counter.Add("User 4569 visit home")
	fmt.Println(b) // 2
	b = counter.Add("User 4569 visit home")
	fmt.Println(b) // 3
	b = counter.Add("User 888 visit home")
	fmt.Println(b) // 1
	b = counter.HowMany("User 888 visit home")
	fmt.Println(b) // 1
	b = counter.HowMany("User 4569 visit home")
	fmt.Println(b) // 3
	b = counter.HowMany("User 4560 visit home")
	fmt.Println(b) // 0
}

```

# Credit
Thanks to [https://github.com/influxdata/influxdb] for Hyperloglog Library. When we test their library, error rate max only at 0.2%

# Milestone
Now we only count max 8 every key. Next we will Upgrade Capacity to max integer