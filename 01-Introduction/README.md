# Introducción a Micro-Servicios

Los microservicios son una arquitectura de sofware que organiza una aplicación como un conjunto de servicios pequeños e independientes que se comunican entre sí. Cada microservicio está diseñado para realizar una tarea específica o manejar una funcionalidad particular dentro de una aplicación más grande. Estos servicios son autónomos, lo que significa que pueden ser desarrollados, desplegados y escalados de forma independiente.

En lugar de construir un aplicación monolítica en la que todas las funcionalidades están interconectadas y ejecutándose en un único proceso, los microservicios permiten que cada componente funcion por separado. Esto facilita la acrualización, el mantenimiento y el despliegue de nuevas versiones de la aplicación sin afectar a toda la arquitectura.

Las caracteristicas clave de los microservicios incluyen:
1. **Descentralización:** Los microservicios pueden ser desarrollados utilizando diferentes tecnologías y lenguajes de programación, lo que permite elegir la herramienta adecuada para cada tarea.
2. **Escalabilidad:** Cada microservicio puede ser escalado de manera independiente según la demanda, lo que optimiza el uso de recursos.
3. **Despliegue independiente:** Los microservicios pueden ser desplegados sin necesidad de actualizar toda la aplicación, reduciendo los tiempos de inactividad.
4. **Resiliencia:** Si un microservicio falla, no necesariamente afenata a toda la aplicación, ya que los demás servicios pueden seguir funcionando.
5. **Comunicación a través de APIs:** Los microservicios se comunican entre sí a través de APIs bien definidas, generalmente urilizando protocolos como HTTP/REST o mensajería asincrónica.

Esta arquitectura es popular en entornos de desarrrollo ágil y DevOps, ya que facilita la implementación conrinua y la mejora continua del sorware. Sin embargo, tambien introduce desafíos como la gestion de la complejidad de la comunicación entre servicios, la orquestación y la obsevabilidad de la aplicación.

# Construcción de un servidor web simple con net/http

El paquete net/http provee todas las caracteristicas necesarias para escribir clientes y servidores HTTP.
Nos da la capacidad de enviar peticiones a otros servidores comunicandose bajo el protocolo HTTP asi comola abilidad de correr un servidor HTTP que pueda enrutar cada peticion a funciones independientes de go, archivos estáticos y mucho mas.

La sintaxis para crear un servidor basico es:
```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080

	http.HandleFunc("/holamundo", holaMundoHandler)
	
	log.Println("Iniciando el servidor en el puerto:", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port),nil))
}

func holaMundoHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Hola Mundo")
}
```

Lo primero que hacemos es llamar el metodo HandleFunc del paquete http. El método HandleFunc crea un objeto tipo Handler (manejador) en el multiplexor del DefaultServeMux handler, mapeando el path pasado en el primer argumento y ejecuta la función pasada como segundo argumento:
`func HandleFunc(pattern string, handler func(ResponseWriter, *Request))`

Luego iniciamos el servidor invocando el metodo **ListenAndServe** el cual toma dos parámetros, la dirección TCP y el handler que va a manejar las peticiones: `func ListenAndServe(addr string, handler Handler) error`

Como segundo parámetro en este caso le estamos pasando **nil** ya que queremos que use el DefaultServeMux handler o el enrutador por defecto de la instacia de servidor.

Ahora simplemtente ejecutamos el servidor
```bash
go run basic_http_server.go
```
Para validar el funcionamiento desde un navegador accede a http://localhost:8080/holamundo

# Lectura y Escritura de JSON

En Go, el manejo de JSON se realiza principalmente a través del paquete **encoding/json**, que proporciona funciones para convertir datos a y desde el formato JSON. Las dos funciones más comunes en este paquete son Marshal y Unmarshal.

## json.Marshal
La función json.Marshal se utiliza para convertir estructuras de datos de Go (como structs, mapas, slices, etc.) en una representación JSON. Cuando se llama a Marshal, devuelve los datos en formato JSON como un slice de bytes y un posible error si ocurre algún problema.

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Persona struct {
	Nombre   string `json:"nombre"`
	Edad     int    `json:"edad"`
	EsCasado bool   `json:"es_casado"`
}

func main() {
	p := Persona{
		Nombre:   "Juan",
		Edad:     30,
		EsCasado: true,
	}

	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	fmt.Println(string(jsonData)) // {"nombre":"Juan","edad":30,"es_casado":true}
}

```

En este ejemplo, la estructura **Persona** se convierte en un objeto JSON. Los campos de la estructura pueden ser etiquetados con etiquetas json para definir cómo deben ser serializados.

## json.Unmarshal
La función `json.Unmarshal` se utiliza para convertir datos en formato JSON de vuelta a estructuras de datos de Go. Toma un slice de bytes que contiene el JSON y un puntero a una estructura en la que los datos deben deserializarse.
`func Unmarshal(data []byte, v interface{}) error` 

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Persona struct {
	Nombre   string `json:"nombre"`
	Edad     int    `json:"edad"`
	EsCasado bool   `json:"es_casado"`
}

func main() {
	jsonData := `{"nombre":"Juan","edad":30,"es_casado":true}`

	var p Persona

	err := json.Unmarshal([]byte(jsonData), &p)
	if err != nil {
		fmt.Println("Error al convertir desde JSON:", err)
		return
	}

	fmt.Printf("Nombre: %s, Edad: %d, EsCasado: %t\n", p.Nombre, p.Edad, p.EsCasado)
}

```
En este ejemplo, el JSON en la cadena jsonData se deserializa en la estructura Persona. Unmarshal requiere un puntero a la estructura porque modifica directamente los datos en la memoria.

## Consideraciones
Etiquetas JSON: Como se ve en los ejemplos, las etiquetas JSON (json:"nombre") permiten personalizar los nombres de los campos en el JSON y controlar su visibilidad.

Campos Exportados: Solo los campos exportados (es decir, los que comienzan con una letra mayúscula) pueden ser serializados y deserializados. Si un campo no es exportado (empieza con minúscula), será ignorado por Marshal y Unmarshal.

Manejo de Errores: Es importante manejar los posibles errores devueltos por Marshal y Unmarshal, especialmente cuando se trabaja con datos JSON no confiables o complejos.

Este manejo de JSON en Go es poderoso y flexible, permitiendo trabajar con datos estructurados de manera eficiente y segura, podemos usar etiquetas para controlar la salida, ignorar campos etc.

```go
type Persona struct {
	Nombre   string `json:"nombre"`
	Edad     int    `json:"edad"`
    //No mostrar este campo
	EsCasado bool   `json:"-"`
    // No mostrar este campo si está vacio
    Direccion string `json:",ommitempty"`
    // Convertir el campo a string y renombrarlo
    ID int `json:"id, string"`
}
```

Ahora si llevamos esto a nuestro servidor no seria muy eficiente decodificar una structura en bytes y luego escribirla destro de la respuesta. Go provee codificadores y decodificadores (encoders, decoders) los cuales pueden escribir directamente a un stream de datos como lo es el ResponseWriter de los handlers

## ResponseWriter

El ResponseWriter es una interfaz que define tres métodos
```go
// Regresa el mapa de encabezados el cual debe er enviado por el método WriteHeader
Header()

//Escribe la data a la conexión. si el WriteHeader no ha sido invocado entonces se invoca
// WriteHeader(http.StatusOK)
Write([]byte) (int, error)

// Envia un cabecero de respuesta HTTP con el codigo de status
WriteHeader(int)
```
Si tenemos una interfaz **ResponseWriter** cómo podemos usarla con `fmt.Fprint(w io.Writer, a ...interface{})`? Este método requiere una interfaz Writer como parámetro y nosotros tenemos una RespnseWriter, si miramos la firma del Writer se ve de la siguiente manera: 
- `Write(p []byte) (n int, err error)`
Como la interfaz del ResponseWriter implementa este método tambien satisface la interfaz Writer y asi cualquier objeto que implemente ResponseWriter puede ser pasado a cualquier función que espere un Writer.

El paquete encoding/json tiene una función llamada NewEncoder esto retorna un objeto de tipo Encoder el cual puede ser usado para escribir directamente un json en un Writer abierto, asi que en vez de almacenar la salida del marshall en un arreglo de bytes podemos escribirlo directamente en la respuesta HTTP.

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type holaMundoResponse struct {
	Message string `json:"message"`
}

func main() {
	port := 8080

	http.HandleFunc("/holamundo", holaMundoHandler)
	
	log.Println("Iniciando el servidor en el puerto:", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port),nil))
}

func holaMundoHandler(w http.ResponseWriter, r *http.Request){
	response := holaMundoResponse{Message: "Hola mundo!!"}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
```

# Enrutamiento

Incluso un simple microservicio va a necesitar la capacidad de enrutar las peticiones a diferentes handlers dependiento en la ruta de la petición o en el método, para facilitar el manejo de manejar multiples rutas el paquete http tiene un objeto especial llamado **ServerMux**
el cual implementa la interfaz http.Handler.

Hay dos funciones para agregar handlers al ServerMux:
- func HandlerFunc(pattern string, handler func(ResponseWriter, *Request))
- func Handle(pattern string, handler Handler)

**HandlerFunc:** es una función de conveniencia que crea un manejador cuyo método ServeHTTP llama a una función ordinaria con la firma func(ResponseWriter, *Request) que se pasa como parámetro.

**Handle:** requiere que pases dos parámetros: el patrón con el que deseas registrar el manejador y un objeto que implemente la interfaz Handler:
```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

## Paths
Ya hemos hablado sobre que el ServerMux es el encargado de enrutar las peticiones entrantes a los handlers regeistrados, sin embargo, la forma en la que se hace match con las rutas puede ser confusa.

Puedes registrar tanto rutas fijas como /images/cat.jpg, o subárboles enraizados como /images/. La barra diagonal final en el subárbol enraizado es importante, ya que cualquier solicitud que comience con /images/, por ejemplo, /images/happy_cat.jpg, se direccionará al manejador asociado con /images/.

Si registramos la ruta /images/ con el manejador foo, y un usuario hace una solicitud a nuestro servicio en /images (sin la barra diagonal final), entonces ServerMux reenviará la solicitud al manejador de /images/, agregando una barra diagonal al final.

Si también registramos la ruta /images (sin la barra diagonal final) con el manejador bar, y el usuario solicita /images, entonces esta solicitud se dirigirá a bar; sin embargo, /images/ o /images/cat.jpg se dirigirán a foo:

```go
http.Handle("/images/", newFooHandler())
http.Handle("/images/persian/", newBarHandler())
http.Handle("/images", newBuzzHandler())
```
- /images => Buzz
- /images/ => Foo
- /images/cat => Foo
- /images/cat.jpg => Foo
- /images/persian/cat.jpg => Bar
Las rutas más largas siempre tendrán prioridad sobre las más cortas, por lo que es posible tener una ruta explícita que apunte a un manejador diferente de una ruta general.

También podemos especificar el nombre del host. Por ejemplo, podríamos registrar una ruta como search.google.com/ y /ServerMux reenviaría cualquier solicitud a http://search.google.com y http://www.google.com a sus respectivos handlers.

## Convenience Handler (Manejador de conveniencia)

Un convenience handler (o manejador de conveniencia) es un término que se refiere a una función o método en un framework o biblioteca que simplifica el proceso de crear y registrar handlers para solicitudes HTTP. Estos handlers de conveniencia están diseñados para reducir el código repetitivo o hacer que las tareas comunes sean más fáciles de implementar.

En el contexto de Go, un ejemplo típico de un convenience handler es la función http.HandleFunc. En lugar de requerir que implementes manualmente la interfaz http.Handler y su método ServeHTTP, http.HandleFunc te permite pasar directamente una función con la firma func(http.ResponseWriter, *http.Request). Esta función se convierte automáticamente en un manejador que se puede registrar para manejar solicitudes a una ruta específica.

Por ejemplo, en lugar de escribir:

```go
type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, world!")
}

http.Handle("/hello", MyHandler{})
```
Puedes usar un convenience handler como http.HandleFunc:

```go
http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, world!")
})
```
Esto simplifica la creación de handlers y es especialmente útil para rutas simples o aplicaciones pequeñas, donde no es necesario implementar toda la interfaz http.Handler.

El paquete net/http implementa varios métodos que crean diferentes tipos de convenience handlers.

## FileServer

Una función FileServer devuelve un manejador que atiende solicitudes HTTP con el contenido del sistema de archivos. Esto se puede utilizar para servir archivos estáticos, como imágenes u otro contenido almacenado en el sistema de archivos:

```go
func FileServer(root FileSystem) Handler
```
Observa el siguiente código:

```go
http.Handle("/images", http.FileServer(http.Dir("./images")))
```
Esto nos permite mapear el contenido de la ruta del sistema de archivos ./images a la ruta del servidor /images. Dir implementa un sistema de archivos que está restringido a un árbol de directorios específico, y el método FileServer utiliza esto para poder servir los recursos.


## NotFoundHandler
La función NotFoundHandler devuelve un manejador de solicitudes simple que responde a cada solicitud con una respuesta de "404 página no encontrada":

```go
func NotFoundHandler() Handler
```
## RedirectHandler
La función RedirectHandler devuelve un manejador de solicitudes que redirige cada solicitud que recibe a la URI dada, utilizando el código de estado proporcionado. El código proporcionado debe estar en el rango 3xx y generalmente es StatusMovedPermanently, StatusFound, o StatusSeeOther:

```go
func RedirectHandler(url string, code int) Handler
```
## StripPrefix
La función StripPrefix devuelve un manejador que atiende solicitudes HTTP eliminando el prefijo dado de la ruta del URL de la solicitud y luego invocando el manejador h. Si una ruta no existe, entonces StripPrefix responderá con un error HTTP 404 "no encontrado":

```go
func StripPrefix(prefix string, h Handler) Handler
```

## TimeoutHandler
La función TimeoutHandler devuelve una interfaz Handler que ejecuta h con el límite de tiempo dado. 

```go
func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler
```
El nuevo manejador llama a h.ServeHTTP para manejar cada solicitud, pero si una llamada dura más tiempo que el límite establecido, el manejador responde con una respuesta 503 "Servicio no disponible" con el mensaje proporcionado (msg) en su cuerpo.

Los dos últimos handlers son especialmente interesantes, ya que, en efecto, encadenan handlers. Esta es una técnica que abordaremos más a fondo en un capítulo posterior, ya que permite practicar código limpio y mantener tu código DRY (Don't Repeat Yourself).


## Crear Handlers

Para finalizar el ejemplo vamos a crear un Handler en lugar de solo usar **HandleFunc**. Vamos a dividir el código que realiza la validación de solicitudes para nuestro endpoint helloworld y el código que devuelve la respuesta en manejadores separados para ilustrar cómo es posible encadenar manejadores.

```go
type validationHandler struct {
    next http.Handler
}

func newValidationHandler(next http.Handler) http.Handler {
    return validationHandler{next: next}
}
```

Lo primero que necesitamos hacer al crear nuestro propio Handler es definir un campo en la estructura que implementará los métodos en la interfaz Handler. En este ejemplo vamos a encadenar Handlers, el primer Handler, que es nuestro manejador de validación, necesita tener una referencia al siguiente en la cadena, ya que tiene la responsabilidad de llamar a **ServeHTTP** o devolver una respuesta.

Para mayor comodidad, hemos añadido una función que nos devuelve un nuevo manejador; sin embargo, podríamos haber configurado simplemente el campo **next**. Este método, sin embargo, es una mejor práctica, ya que hace que nuestro código sea un poco más fácil de leer y, cuando necesitamos pasar dependencias complejas al manejador, usar una función para crear el manejador mantiene las cosas un poco más ordenadas:

```go
func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
    var request helloWorldRequest
    decoder := json.NewDecoder(r.Body)

    err := decoder.Decode(&request)
    if err != nil {
        http.Error(rw, "Bad request", http.StatusBadRequest)
        return
    }

    h.next.ServeHTTP(rw, r)
}
```
El bloque de código anterior ilustra cómo implementaríamos el método ServeHTTP. Lo único interesante que cabe destacar aquí es la instrucción que comienza en la línea del return. Si se devuelve un error al decodificar la solicitud, escribimos un error 500 en la respuesta, y la cadena de manejadores se detendría aquí. Solo cuando no se devuelve ningún error llamamos al siguiente manejador en la cadena, y lo hacemos simplemente invocando su método ServeHTTP. Para pasar el nombre decodificado de la solicitud, simplemente estamos configurando una variable:

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

func main() {
	port := 8080

	handler := newValidationHandler(newHelloWorldHandler())

	http.Handle("/helloworld", handler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

type validationHandler struct {
	next http.Handler
}

func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}

	h.next.ServeHTTP(rw, r)
}

type helloWorldHandler struct{}

func newHelloWorldHandler() http.Handler {
	return helloWorldHandler{}
}

func (h helloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello"}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}
```
El tipo **helloWorldHandler** que escribe la respuesta no se ve muy diferente de cuando estábamos usando una función simple. lo único que realmente hemos hecho es eliminar la decodificación de la solicitud.

Ahora, este código es puramente para ilustrar cómo se puede hacer algo, no necesariamente es la mejor manera de hacerlo. 
En este caso simple, dividir la validación de la solicitud y el envío de la respuesta en dos manejadores agrega mucha complejidad innecesaria y realmente no está haciendo nuestro código más DRY (Don't Repeat Yourself). Sin embargo, la técnica es útil. 



# Context

El problema con el patrón anterior es que no hay forma de pasar la solicitud validada de un Handler al siguiente sin romper la interfaz http.Handler. Pero Go tiene una solución. 
El paquete **context**; el tipo Context implementa un método seguro para acceder a datos con alcance de solicitud que es seguro para usar simultáneamente por múltiples rutinas de Go. 

## Background

El método Background devuelve un contexto vacío que no tiene valores; típicamente es utilizado por la función main y como el contexto de nivel superior:
`func Background() Context`

## WithCancel

El método WithCancel devuelve una copia del contexto padre con una función de cancelación. Llamar a la función de cancelación libera los recursos asociados con el contexto y debe ser llamada tan pronto como las operaciones que se ejecutan en el tipo Context se completen:
`func WithCancel(parent Context) (ctx Context, cancel CancelFunc)`

## WithDeadline

El método WithDeadline devuelve una copia del contexto padre que expira después de que el tiempo actual sea mayor que el tiempo (deadline). En este punto, el canal **Done** del contexto se cierra y los recursos asociados se liberan. 
También devuelve un método **CancelFunc** que permite la cancelación manual del contexto:
`func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)`


## WithTimeout
El método WithTimeout es similar a WithDeadline, excepto que en lugar de un plazo específico, le pasas una duración por la cual el tipo Context debería existir. Una vez que esta duración ha transcurrido, el canal **Done** se cierra y los recursos asociados con el contexto se liberan: 
`func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)`

## WithValue

El método WithValue devuelve una copia del contexto padre en el que el valor val está asociado con la clave key. Los valores de Context son perfectos para usarse con datos con alcance de solicitud:
`func WithValue(parent Context, key interface{}, val interface{}) Context`

## Uso de contextos

El método **Context()** nos da acceso a una estructura context.Context, que siempre es no nula, ya que se inicializa cuando se crea originalmente la solicitud. Para solicitudes entrantes, el **http.Server** gestiona automáticamente el ciclo de vida del contexto, cancelándolo cuando la conexión del cliente se cierra. Para solicitudes salientes, Context controla la cancelación; es decir, si cancelamos el método Context(), podemos cancelar la solicitud saliente. Este concepto se ilustra en el siguiente ejemplo:

```go
func fetchGoogle(t *testing.T) {
    r, _ := http.NewRequest("GET", "https://google.com", nil)

    timeoutRequest, cancelFunc := context.WithTimeout(r.Context(), 1*time.Millisecond)
    defer cancelFunc()

    r = r.WithContext(timeoutRequest)

    _, err := http.DefaultClient.Do(r)
    if err != nil {
        fmt.Println("Error:", err)
    }
}
```

En la línea `timeoutRequest, cancelFunc := context.WithTimeout(r.Context(), 1*time.Millisecond)`, estamos creando un contexto con tiempo de espera (timeout) a partir del original en la solicitud, y a diferencia de una solicitud entrante donde el contexto se cancela automáticamente, en una solicitud saliente debemos realizar este paso manualmente.

La línea `r = r.WithContext(timeoutRequest)` implementa el segundo de los dos nuevos métodos de contexto que se han agregado al objeto http.Request:

```go
func (r *Request) WithContext(ctx context.Context) *Request
```
El método WithContext devuelve una copia superficial de la solicitud original en la que el contexto se ha cambiado al contexto ctx proporcionado.

Cuando ejecutamos esta función, veremos que después de 1 milisegundo, la solicitud se completará con un error:

```txt
Error: Get https://google.com: context deadline exceeded
```
El contexto se agota antes de que la solicitud tenga la oportunidad de completarse, y el método Do devuelve inmediatamente. Esta es una excelente técnica para usar en conexiones salientes.

---

En el ejemplo anterior, podemos actualizar la conexión entrante para aprovechar el paquete context y así implementar un acceso seguro a objetos mediante goroutines. 

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

type validationContextKey string

type helloWorldResponse struct {
	Message string `json:"message"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

func main() {
	port := 8080

	handler := newValidationHandler(newHelloWorldHandler())
	http.Handle("/helloworld", handler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

type validationHandler struct {
	next http.Handler
}

func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), validationContextKey("name"), request.Name)
	r = r.WithContext(c)

	h.next.ServeHTTP(rw, r)
}

type helloWorldHandler struct {
}

func newHelloWorldHandler() http.Handler {
	return helloWorldHandler{}
}

func (h helloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	name := r.Context().Value(validationContextKey("name")).(string)
	response := helloWorldResponse{Message: "Hello " + name}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

func fetchGoogle(t *testing.T) {
	r, _ := http.NewRequest("GET", "https://google.com", nil)

	timeoutRequest, cancelFunc := context.WithTimeout(r.Context(), 1*time.Millisecond)
	defer cancelFunc()

	r = r.WithContext(timeoutRequest)

	_, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
```
En este código, cuando tenemos una solicitud válida, estamos creando un nuevo contexto para esta solicitud y luego establecemos el valor del campo Name en la solicitud dentro del contexto.

En la `c := context.WithValue(r.Context(), validationContextKey("name"),request.Name)`, cuando agregas un elemento a un contexto usando WithValue, el método devuelve una copia del contexto anterior. Para ahorrar tiempo y simplificar el código, estamos utilizando un puntero al contexto. Para pasar esto como una copia a WithValue, debemos desreferenciar el puntero. Luego, para actualizar nuestro puntero, también debemos establecer el valor devuelto como el valor referenciado por el puntero, lo que significa que debemos desreferenciarlo nuevamente.

Otro detalle a observar es la clave que estamos utilizando, validationContextKey, que es un tipo explícito declarado de string:
`type validationContextKey string`
Esta clave se utiliza para evitar colisiones en el contexto, ya que cada tipo de clave es único. De esta manera, podemos asegurarnos de que nuestro valor almacenado en el contexto sea seguro y no interfiera con otros valores potenciales.

En el ejemplo anterior, al no usar una simple cadena de texto como clave para el contexto, evitamos posibles colisiones. Esto es especialmente importante cuando el contexto fluye a través de diferentes paquetes. Si se usara una cadena común como "name", podría haber un conflicto si otro paquete también usa el mismo nombre de clave. Al declarar un tipo a nivel de paquete como validationContextKey y utilizarlo, aseguramos que nuestros valores en el contexto sean únicos y no sean sobrescritos inadvertidamente por otro código.

A continuación, en el helloWorldHandler, recuperamos el valor del contexto de la siguiente manera:

```go
func (h helloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
    name := r.Context().Value(validationContextKey("name")).(string)
    response := helloWorldResponse{Message: "Hello " + name}

    encoder := json.NewEncoder(rw)
    encoder.Encode(response)
}
```
Aquí, obtenemos el contexto y llamamos al método Value con nuestra clave validationContextKey("name"), que devuelve el valor asociado. Luego, hacemos un type assertion para convertirlo a una cadena.

# Soporte RPC en la biblioteca estándar de Go
Go ofrece un excelente soporte para RPC (Remote Procedure Call) en su biblioteca estándar. Aquí veremos cómo crear un cliente y un servidor que se comunican a través de RPC utilizando una interfaz compartida.

## Ejemplo simple de RPC

Veamos un ejemplo básico en el que utilizamos el paquete estándar rpc para construir una API basada en RPC en Go. Seguimos un ejemplo típico de "Hello World":

```go
// rpc/server/server.go
type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
    reply.Message = "Hello " + args.Name
    return nil
}
```

Al igual que en el ejemplo de creación de APIs REST usando la biblioteca estándar para RPC, también definiremos un handler. La diferencia entre este handler y http.Handler es que no necesita ajustarse a una interfaz; mientras tengamos un campo de struct con métodos, podemos registrarlo con el servidor RPC:

```go
func Register(rcvr interface{}) error
```
La función Register, que se encuentra en el paquete rpc, publica los métodos que forman parte de la interfaz dada en el servidor predeterminado y permite que sean llamados por los clientes que se conectan al servicio. El nombre del método usa el nombre del tipo concreto, por lo que, en este caso, si un cliente quiere llamar al método HelloWorld, se accedería a él usando HelloWorldHandler.HelloWorld. Si no deseamos usar el nombre del tipo concreto, podemos registrarlo con un nombre diferente usando la función RegisterName, que utiliza el nombre proporcionado en su lugar:

```go
func RegisterName(name string, rcvr interface{}) error
```
Esto permitiría mantener el nombre del campo de struct como sea significativo para mi código; sin embargo, para el contrato de cliente podría decidir usar algo diferente como Greet:

```go
func StartServer() {
    helloWorld := &HelloWorldHandler{}
    rpc.Register(helloWorld)

    l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
    if err != nil {
        log.Fatal(fmt.Sprintf("No se pudo escuchar en el puerto dado: %s", err))
    }

    for {
        conn, _ := l.Accept()
        go rpc.ServeConn(conn)
    }
}
```
En la función StartServer, primero creamos una nueva instancia de nuestro handler y luego la registramos con el servidor RPC predeterminado.

A diferencia de la comodidad de **net/http**, donde solo necesitamos crear un servidor con ListenAndServe, cuando usamos RPC necesitamos hacer un poco más de trabajo manual. En la línea `l, err := net.Listen("("tcp",", fmt.Sprintf(":%(":%v",", port))`, estamos creando un socket usando el protocolo dado y vinculándolo a la dirección IP y puerto. Esto nos da la capacidad de seleccionar específicamente el protocolo que queremos usar para el servidor: tcp, tcp4, tcp6, unix o unixpacket:

```go
func Listen(net, laddr string) (Listener, error)
```
La función Listen() devuelve una instancia que implementa la interfaz Listener:

```go
type Listener interface {
    // Accept espera y devuelve la siguiente conexión al listener.
    Accept() (Conn, error)
    // Close cierra el listener.
    // Cualquier operación Accept bloqueada se desbloqueará y devolverá errores.
    Close() error
    // Addr devuelve la dirección de red del listener.
    Addr() Addr
}
```
Para recibir conexiones, debemos llamar al método **Accept** en el listener. Si observas la línea `conn, _ := l.Accept()`, verás que tenemos un bucle infinito, esto es porque, a diferencia de **ListenAndServe** que bloquea para todas las conexiones, con un servidor RPC manejamos cada conexión de forma individual y, tan pronto como manejamos la primera conexión, necesitamos continuar llamando a Accept para manejar las conexiones subsiguientes o la aplicación saldrá. Accept es un método bloqueante, por lo que si no hay clientes intentando conectar al servicio, *Accept bloqueará hasta que uno lo haga*. Una vez que recibimos una conexión, necesitamos llamar al método Accept nuevamente para procesar la siguiente conexión. Si miras la línea `go rpc.ServeConn(conn)` en el código de ejemplo, verás que se llama al método **ServeConn**:
```go
func ServeConn(conn io.ReadWriteCloser)
```
El método ServeConn ejecuta el método DefaultServer en la conexión dada y bloqueará hasta que el cliente termine. En el ejemplo, usamos la declaración go antes de ejecutar el servidor para que podamos procesar inmediatamente la siguiente conexión en espera sin bloquear la conexión del primer cliente.
---
En términos de protocolo de comunicación, ServeConn utiliza el formato de serialización **gob**. El formato gob fue diseñado específicamente para facilitar la comunicación basada en Go, siendo más fácil de usar y posiblemente más eficiente que otros formatos como los **protocol buffers**, aunque con la desventaja de no ser ideal para la comunicación entre lenguajes diferentes.

Con gob, los valores y tipos entre el origen y el destino no necesitan coincidir exactamente. Si un campo está presente en el origen pero no en la estructura receptora, el decodificador lo ignorará y continuará procesando sin error. Del mismo modo, si un campo está presente en el destino pero no en el origen, el decodificador lo ignorará y procesará el resto del mensaje. Esta flexibilidad contrasta con otros métodos RPC más antiguos, como JMI, que requerían que las interfaces del cliente y del servidor fueran idénticas, lo que generaba un acoplamiento estrecho entre las bases de código y complicaba la implementación de actualizaciones en la aplicación.

Para realizar una solicitud a nuestro cliente, ya no podemos simplemente usar curl, ya que no estamos usando el protocolo HTTP ni el formato JSON. En su lugar, debemos crear un cliente RPC en Go:

```go
func CreateClient() *rpc.Client {
    client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%v", port))
    if err != nil {
        log.Fatal("dialing:", err)
    }
    return client
}
```
En el bloque anterior, se muestra cómo configurar **rpc.Client**. En la línea `client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%v", port))`, creamos el cliente usando la función **Dial()** del paquete rpc. Luego, utilizamos esta conexión devuelta para hacer una solicitud al servidor:

```go
func PerformRequest(client *rpc.Client) contract.HelloWorldResponse {
    args := &contract.HelloWorldRequest{Name: "World"}
    var reply contract.HelloWorldResponse

    err := client.Call("HelloWorldHandler.HelloWorld", args, &reply)
    if err != nil {
        log.Fatal("error:", err)
    }

    return reply
}
```
En la línea `err := client.Call("HelloWorldHandler.HelloWorld", args, &reply)`, estamos utilizando el método **Call()** del cliente para invocar la función nombrada en el servidor. Call es una función bloqueante que espera hasta que el servidor envíe una respuesta y escribe la respuesta en la referencia de HelloWorldResponse pasada al método. Si ocurre un error al procesar la solicitud, este se devuelve y se puede manejar adecuadamente.

# RPC sobre HTTP

En caso de que necesites usar HTTP como protocolo de transporte, el paquete rpc puede facilitarlo llamando al método **HandleHTTP**. 

Este método configura dos endpoints en tu aplicación: 
```go
const (
	// Defaults used by HandleHTTP
	DefaultRPCPath = "/_"/_goRPC_"_"
	DefaultDebugPath = "/"/debug/rpc""
)
```

pero ten en cuenta que los mensajes siguen estando codificados en gob, por lo que sería necesario escribir un codificador y decodificador gob en JavaScript para comunicarse desde un navegador web, lo cual no es el propósito del paquete y no se recomienda. También, el endpoint de depuración no proporciona documentación generada automáticamente para tu API.
Aquí hay un ejemplo de un servidor RPC sobre HTTP:

```go
func StartServer() {
    helloWorld := &HelloWorldHandler{}
    rpc.Register(helloWorld)
    rpc.HandleHTTP()

    l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
    if err != nil {
        log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
    }

    log.Printf("Server starting on port %v\n", port)
    http.Serve(l, nil)
}
```
En la línea 4, llamamos al método **rpc.HandleHTTP**, que es necesario cuando se usa HTTP con RPC, ya que registra los manejadores HTTP con el método DefaultServer. Luego, llamamos al método **http.Serve** y le pasamos el listener que estamos creando en la línea 6 y el segundo parametro en nil ya que queremos que use el método **DefaultServer**.

# JSON-RPC sobre HTTP

Este último ejemplo, veremos el paquete **net/rpc/jsonrpc** que proporciona un códec integrado para serializar y deserializar al estándar JSON-RPC. 

El método StartServer no contiene nada que no hayamos visto antes; es la configuración estándar del servidor RPC. La principal diferencia está en la línea 10, donde en lugar de iniciar el servidor RPC, estamos iniciando un servidor HTTP y pasándole el listener junto con un manejador:

```go
func StartServer() {  
  helloWorld := new(HelloWorldHandler)  
  rpc.Register(helloWorld)  
 
  l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))  
  if err != nil {  
    log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))  
  }  
 
  http.Serve(l, http.HandlerFunc(httpHandler))  
}
```
El manejador que estamos pasando al servidor es donde ocurre la "magia":

```go
func httpHandler(w http.ResponseWriter, r *http.Request) {  
  serverCodec := jsonrpc.NewServerCodec(&HttpConn{in: r.Body, out: w})  
  err := rpc.ServeRequest(serverCodec)  
  if err != nil {  
    log.Printf("Error while serving JSON request: %v", err)  
    http.Error(w, "Error while serving JSON request, details have been logged.", 500)  
    return  
  }  
}
```
En la línea 2, estamos llamando a la función **jsonrpc.NewServerCodec** y pasándole un tipo que implementa **io.ReadWriteCloser**. El método **NewServerCodec** devuelve un tipo que implementa **rpc.ClientCodec**, que tiene los siguientes métodos:

```go
type ClientCodec interface {  
  WriteRequest(*Request, interface{}) error  
  ReadResponseHeader(*Response) error  
  ReadResponseBody(interface{}) error  
  Close() error  
}
```
Un tipo **ClientCodec** implementa la escritura de solicitudes RPC y la lectura de respuestas RPC. Para escribir una solicitud en la conexión, un cliente llama al método **WriteRequest**. Para leer la respuesta, el cliente debe llamar a **ReadResponseHeader** y **ReadResponseBody** como un par. Una vez que el cuerpo ha sido leído, es responsabilidad del cliente llamar al método Close para cerrar la conexión. Si se pasa una interfaz nil a **ReadResponseBody**, el cuerpo de la respuesta debe leerse y luego descartarse.

```go
type HttpConn struct {  
  in  io.Reader  
  out io.Writer  
}
func (c *HttpConn) Read(p []byte) (n int, err error) { return c.in.Read(p) }  
func (c *HttpConn) Write(d []byte) (n int, err error) { return c.out.Write(d) }  
func (c *HttpConn) Close() error { return nil }
```
El método **NewServerCodec** requiere que le pasemos un tipo que implemente la interfaz **ReadWriteCloser**. Como no tenemos tal tipo pasado como parámetro en el método **httpHandler**, hemos definido nuestro propio tipo, **HttpConn**, que encapsula el cuerpo de **http.Request**, que implementa **io.Reader**, y el método **ResponseWriter**, que implementa **io.Writer**. Luego, podemos escribir nuestros propios métodos que hagan proxy de las llamadas al lector y escritor, creando un tipo que tenga la interfaz correcta.






