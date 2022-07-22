* Create a folder say __utils__
* Navigate to __utils__ folder in terminal
* Run the following command to initialize it as a module

```
go mod init "mycompany.com/utils"
```

* Create a go file say __math_constants_.go__ with the following code

``` go
package utils

const PI = 3.14
const Length = 12
const radius = 34 //Will not be exported

```

* Navigate out of __utils__ folder
* Create a file main.go with the following code. You will use utils package items

``` go
package main

import (
	"fmt"
	"mycompany.com/utils"
)

func main() {
	fmt.Println(utils.PI, utils.Length)
}

```

* Let's create a new module __mycompany.com/main__

```
go mod init "mycompany.com/main"
```

* You need to specify the path of the __utils__ package by running the following commands

```
go mod edit --replace mycompany.com/utils=./utils
go mod tidy
```

* Take a look at the go.mod file generated
* You can run __go run main.go__ file
