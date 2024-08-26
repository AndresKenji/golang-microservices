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

Puede que haya tomado la mayoría de las descripciones para estos handlers directamente de la documentación de Go, y probablemente ya las hayas leído porque, ¿has leído la documentación, verdad? Con Go, la documentación es excelente y escribir documentación para tus propios paquetes está muy recomendado, incluso es obligatorio si usas el comando golint que viene con el paquete estándar; esto reportará áreas de tu código que no se ajustan a los estándares. Realmente recomiendo pasar un tiempo navegando por la documentación estándar cuando uses uno de los paquetes; no solo aprenderás el uso correcto, sino que también puedes descubrir un enfoque mejor. Sin duda estarás expuesto a buenas prácticas y estilos, y quizás incluso puedas seguir trabajando en el triste día en que Stack Overflow deje de funcionar y toda la industria se detenga.





