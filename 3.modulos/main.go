/**
Inicializar el uso de módulos con el repositorio de nuestro proyecto
en github:

> go mod init github.com/ingporto/go-intermediate

Se crea automaticamente el archivo go.mod. Ahí se almacenarán todas
las dependencias.

Al momento de instalar un módulo se crea automaticamente
un archivo go.sum:

> go install github.com/donvito/hellomod

go.sum verifica que no se instale malware al verificar el repo y su
versión, con un hash único.

**************************************************

Eliminar dependencias que no se usan en el build

> go mod tidy

Revisar la caché de las descargas previas

> go mod download -json

***************************************************

Instalar una versión diverente de un módulo

> go get github.com/donvito/hellomod/v2


****************************************************
PARA EL ERROR: "gopls requires a module at the root of your workspace"

* ir a vs code settings -> buscar gopls y en el objeto gopls agregar la variable “experimentalWorkspaceModule” con el valor de true.

o 

* go env -w GO111MODULE=on

*/

package main

import (
	"fmt"

	hellomod "github.com/donvito/hellomod"
	hellomod2 "github.com/donvito/hellomod/v2"
)

func main () {
	fmt.Println("Hello World")
	hellomod.SayHello()
	hellomod2.SayHello("People")
}