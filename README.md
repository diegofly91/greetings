# Saludos en Go

Este es un módulo de Go que provee una función para saludar.

## Instalación
Ejecuta el siguiente comando para instalar el módulo:

```bash
go get -u github.com/diegofly91/greetings
``` 
## Uso
```go
package main

import (
    "fmt"
    "github.com/diegofly91/greetings"
)

func main() {
    message, err := greetings.Hello("Diego")
    if err != nil {
        fmt.Println("Ocurrio un error: ", err)
        return
    }
    fmt.Println(message)
}
```