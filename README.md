# payutils-http-gin

[Gin](https://github.com/gin-gonic/gin) HTTP driver for
[payutils](https://go.gh.ink/payutils/v3).

Adapts a Gin router so payutils can register provider callback routes on it.
Standard-library `http.HandlerFunc`s are bridged to Gin via `gin.WrapH`.

## Module

```
go.gh.ink/payutils/http/gin/v3
```

```bash
go get go.gh.ink/payutils/http/gin/v3
```

## Usage

Blank-import the driver (it self-registers under the name `gin`) and pass a Gin
router (`*gin.Engine` or a route group) in `Config.Instances`:

```go
import (
    "github.com/gin-gonic/gin"

    "go.gh.ink/payutils/v3/client"
    "go.gh.ink/payutils/v3/model"

    httpGin "go.gh.ink/payutils/http/gin/v3"
    _ "go.gh.ink/payutils/pay/alipay/v3"
)

r := gin.New()

c, err := client.NewClient(model.Config{
    Endpoint:    "https://api.example.com",
    Instances:   model.I{httpGin.Name: r}, // *gin.Engine or *gin.RouterGroup
    Credentials: model.C{ /* ... */ },
    Contract:    myContract{},
})
```

payutils then registers `POST /{provider}/callback` on the router for each
configured provider.

## Accepted instance

Anything implementing `gin.IRoutes` — `*gin.Engine` and `*gin.RouterGroup`.
Passing an unsupported value makes `NewInstance` return
`errors.ErrUnsupportedInstance`.

## Supported verbs

`Get` · `Post` · `Put` · `Patch` · `Delete` · `Head` · `Options` · `Any`.

## License

See [LICENSE](LICENSE).
