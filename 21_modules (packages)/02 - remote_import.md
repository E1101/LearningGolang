For Example:

```
import "github.com/spf13/viper"
```

When you try to build a program with this import path, the go build command will
search the GOPATH for this package location on disk.

When an import path contains a URL , the Go tooling can be used to fetch the package
from the DVCS and place the code inside the GOPATH at the location that matches the
URL . This fetching is done using the go get command.

Since go get is recursive, it can walk down the source tree for a package and
fetch all the dependencies it finds.
