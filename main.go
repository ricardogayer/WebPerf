package main

// ecobiomatrica ultrasonica

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

var Produtos []Produto

func main() {

	produto1 := Produto{
		Nome:  "Caneta",
		Preco: 1.99,
	}

	produto2 := Produto{
		Nome:  "Lapis",
		Preco: 2.99,
	}

	produto3 := Produto{
		Nome:  "Caderno",
		Preco: 3.99,
	}

	produto4 := Produto{
		Nome:  "Borracha",
		Preco: 4.99,
	}

	produto5 := Produto{
		Nome:  "Apontador",
		Preco: 5.99,
	}

	produto6 := Produto{
		Nome:  "Borracha 2",
		Preco: 6.99,
	}

	produto7 := Produto{
		Nome:  "Apontador 2",
		Preco: 7.99,
	}

	produto8 := Produto{
		Nome:  "Borracha 3",
		Preco: 8.99,
	}

	produto9 := Produto{
		Nome:  "Apontador 2",
		Preco: 9.99,
	}

	produto10 := Produto{
		Nome:  "Borracha 4",
		Preco: 10.99,
	}

	Produtos = append(Produtos, produto1)
	Produtos = append(Produtos, produto2)
	Produtos = append(Produtos, produto3)
	Produtos = append(Produtos, produto4)
	Produtos = append(Produtos, produto5)
	Produtos = append(Produtos, produto6)
	Produtos = append(Produtos, produto7)
	Produtos = append(Produtos, produto8)
	Produtos = append(Produtos, produto9)
	Produtos = append(Produtos, produto10)

	for _, produto := range Produtos {
		fmt.Printf("%s - R$ %.2f\n", produto.Nome, produto.Preco)
	}

	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r.GET("/produtos", func(c *gin.Context) {
		c.JSON(200, Produtos)
	})
	r.Run(":8080")

	// app := fiber.New(fiber.Config{
	// 	Prefork:     true,
	// 	JSONEncoder: json.Marshal,
	// 	JSONDecoder: json.Unmarshal,
	// })

	// app.Get("/produtos", func(c *fiber.Ctx) error {
	// 	return c.JSON(Produtos)
	// })
	// app.Listen(":8080")

}

type Produto struct {
	Nome  string  `json:"Nome"`
	Preco float64 `json:"Preco"`
}
