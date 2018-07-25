If Go was installed under /usr/local/go and your GOPATH was
set to /home/myproject:/home/mylibraries, the compiler would look for
the net/http package in the following order:

```
/usr/local/go/src/pkg/net/http
/home/myproject/src/net/http
/home/mylibraries/src/net/http
```
