# Go Intermedio
@ingporto

## Introducción
### Visualización de las Go Routines
https://www.instagram.com/reel/CvPqKr2v-qX/?igshid=MzRlODBiNWFlZA==

## Programación Orientada a Objetos

### Herencia:
No se extiende de una clase (o `struct`)
* Se crean un `struct`
* Se crea otro `struct` que tiene como atributo el `struct` anterior, solo que se define como campo anónimo (un atributo son nombre, solo con tu tipo)

A partir de ahí, el segundo `struct` cuenta con las propiedades del primero.

### Interfaces: 
No se implementan de forma explícita. 
* Se crea un `type` interface donde se definen los métodos que debe tener cualquier `struct` que implmente esta interfaz. 
* Se crea el o los `structs` con sus respectivos métodos que cumplan con la interfaz.
* Se debe de crear un método que reciba la interface, y ejecute el método que ésta tiene definido.

## Go Modules

## Testing

## Concurrencia

## Proyecto: Servidor con worker pools