/**
Creación de un módulo para que cualquiera lo descargue
*/

// No puede ser package main porque no puede ser usuado
// Por los demás
package utils

import "fmt"

// Primera letra en mayúscula para función pública
func HellowWorld() {
	fmt.Println("Hello World from my custom module")
}

// SUBIR EL MÓDULO A GITHUB
// git init 
// git add .
// git commit -m "first commit"
// git remote add origin https://github.com/ingporto/go-intermediate.git
// git branch -M main
// git push -u origin main

// -------------------------------------------------------
/**

Ya en otro proyecto puedo instalar mi módulo usando

> go get https://github.com/ingporto/go-intermediate.git

Importar el módulo utils

```go
import (
	utils "github.com/ingporto/go-intermediate"
)
```

y ejecutar el metodo HellowWorld

> utils.HellowWorld()

*/