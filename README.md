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

Execução

```sh
echo "GET http://localhost:8080/produtos" | vegeta attack -rate=0 -duration=30s -max-workers=8 | tee results-fiber.bin | vegeta report

Requests      [total, rate, throughput]         1143992, 38133.56, 38132.71
Duration      [total, attack, wait]             30s, 30s, 667.76µs
Latencies     [min, mean, 50, 90, 95, 99, max]  44.434µs, 162.069µs, 142.785µs, 239.841µs, 285.309µs, 520.779µs, 7.184ms
Bytes In      [total, mean]                     387813288, 339.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:1143992  
Error Set:

cat results-fiber.bin | vegeta report -type='hist[0,2ms,4ms,6ms,12ms,24ms]'

Bucket         #        %        Histogram
[0s,    2ms]   1143936  100.00%  ##########################################################################
[2ms,   4ms]   41       0.00%    
[4ms,   6ms]   9        0.00%    
[6ms,   12ms]  6        0.00%    
[12ms,  24ms]  0        0.00%    
[24ms,  +Inf]  0        0.00%  

echo "GET http://localhost:8080/produtos" | vegeta attack -rate=0 -duration=30s -max-workers=8 | tee results-gin.bin | vegeta report

Requests      [total, rate, throughput]         868156, 28938.55, 28938.33
Duration      [total, attack, wait]             30s, 30s, 228.573µs
Latencies     [min, mean, 50, 90, 95, 99, max]  55.586µs, 220.429µs, 169.789µs, 310.49µs, 458.086µs, 1.092ms, 23.638ms
Bytes In      [total, mean]                     294304884, 339.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:868156  

cat results-gin.bin | vegeta report -type='hist[0,2ms,4ms,6ms,12ms,24ms]'

[0s,    2ms]   865154  99.65%  ##########################################################################
[2ms,   4ms]   2214    0.26%   
[4ms,   6ms]   514     0.06%   
[6ms,   12ms]  242     0.03%   
[12ms,  24ms]  32      0.00%   
[24ms,  +Inf]  0       0.00%  

```

* Ricardo Gayer
