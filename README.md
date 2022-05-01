# Comparativo de Performance Fiber vs Gin

## Gin

[Gin](https://gin-gonic.com)

```go
r := gin.New() -- No lugar de gin.Default() para evitar o log das chamadas...
gin.SetMode(gin.ReleaseMode)
gin.DefaultWriter = ioutil.Discard

r.GET("/produtos", func(c *gin.Context) {
   c.JSON(200, Produtos)
})

r.Run(":8080")
```

## Fiber

[Fiber](https://docs.gofiber.io)

```go
import "github.com/goccy/go-json" -- Parse de JSON mais performático

app := fiber.New(fiber.Config{
   Prefork:     true, -- Aproveitamento de threads
   JSONEncoder: json.Marshal,
   JSONDecoder: json.Unmarshal,
})

app.Get("/produtos", func(c *fiber.Ctx) error {
   return c.JSON(Produtos)
})

app.Listen(":8080")
```

## Vegata para realizar os comparativos

[Vegeta](https://github.com/tsenart/vegeta)

Instalação do Vegeta

```sh
brew update && brew install vegeta
```

Executação

```sh
echo "GET http://localhost:8080/produtos" | vegeta attack -rate=0 -duration=30s -max-workers=8 | tee results-fiber.bin | vegeta report

cat results-fiber.bin | vegeta report -type='hist[0,2ms,4ms,6ms,12ms,24ms]'

echo "GET http://localhost:8080/produtos" | vegeta attack -rate=0 -duration=30s -max-workers=8 | tee results-gin.bin | vegeta report

cat results-gin.bin | vegeta report -type='hist[0,2ms,4ms,6ms,12ms,24ms]'
```

* Ricardo Gayer
