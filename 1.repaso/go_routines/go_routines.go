// go mod init      <--- Inicializar la lectura de módulos
package go_routines

import (
	"fmt"
	"time"
)

func DoSomething(c chan int) {
	// Dormir por 3 segundos.
	time.Sleep( 3 * time.Second) 
	fmt.Println("Done")
	c <- 1
}