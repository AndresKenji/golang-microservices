# Diseñando una API

# APIs RESTful

El término REST fue sugerido por Roy Fielding en su disertación doctoral en el año 2000. Significa **Transferencia de Estado Representacional**(representational state transfer) y se describe como: 
 >"REST enfatiza la escalabilidad de las interacciones de componentes, la generalidad de las interfaces, el despliegue independiente de componentes, y los componentes intermediarios para reducir la latencia de interacción, aplicar seguridad y encapsular sistemas heredados."

Tener una API que se ajuste a los principios de REST es lo que la hace RESTful.

# URIs

Una URI (Uniform Resource Identifier) es una cadena de caracteres que identifica de manera única un recurso en Internet o en un sistema. Existen dos tipos principales de URIs: URL (Uniform Resource Locator) y URN (Uniform Resource Name).

Diferencias entre URI, URL y URN:
1. URI (Uniform Resource Identifier):
- Es el concepto general, una categoría que engloba tanto URLs como URNs.
- Puede identificar un recurso de manera directa o indirecta, y no necesariamente indica cómo acceder a él.
- Ejemplo: https://example.com/resource, urn:isbn:0451450523

2. URL (Uniform Resource Locator):
- Es un tipo de URI que localiza un recurso describiendo su método de acceso, como el protocolo, dirección y ruta.
- Básicamente, un URL te dice dónde está un recurso y cómo acceder a él.
- Ejemplo: https://example.com/index.html (protocolo https, dominio example.com, ruta /index.html).

3. URN (Uniform Resource Name):
- Es otro tipo de URI que nombra un recurso de manera única, pero no indica cómo acceder a él.
- Un URN garantiza que el recurso tenga una identidad única, independientemente de su ubicación.
- Ejemplo: urn:isbn:0451450523 (una URN para un libro identificado por su ISBN).

En resumen, todas las URLs y URNs son URIs, pero no todas las URIs son URLs o URNs.

# Formato de URI

El [RFC 3986](https://www.ietf.org/rfc/rfc3986.txt), publicado en 2005, define el formato que hace válidas las URIs:

```bash
    URI = scheme "://" authority "/" path [ "?" query] ["#" fragment"]
    URI = http://myserver.com/mypath?query=1#document
```
Usaremos el elemento de ruta para localizar un endpoint que esté ejecutándose en nuestro servidor. 
En un endpoint REST, esto puede contener tanto parámetros como la ubicación de un documento. 
La cadena de consulta es igualmente importante, ya que se utilizará para pasar parámetros como el número de página o el orden para controlar los datos que se devuelven.

Algunas reglas generales para el formato de URI:

- Una barra inclinada **/** se usa para indicar una relación jerárquica entre recursos.
- Una barra inclinada final **/** no debe incluirse en las URIs.
- Los guiones **-** deben usarse para mejorar la legibilidad.
- No se deben usar guiones bajos **_** en las URIs.
- Se prefieren las letras en minúsculas, ya que la sensibilidad a mayúsculas y minúsculas es un diferenciador en la parte del path de una URI.

El concepto detrás de muchas de las reglas es que una URI debe ser fácil de leer y de construir. También debe ser consistente en la forma en que se construye, por lo que se debe seguir la misma taxonomía para todos los endpoints en tu API.

# Diseño de rutas URI para servicios REST
Las rutas se dividen en documentos, colecciones, stores y controladores.

## Colecciones

Una colección es un directorio de recursos típicamente dividido por parámetros para acceder a un documento individual. Por ejemplo:
```bash
    GET /cats -> Todos los gatos en la colección
    GET /cats/1 -> Documento individual para el gato 1
```
Al definir una colección, siempre debemos usar un sustantivo en plural, como "cats" o "people", para el nombre de la colección.

## Documentos

Un documento es un recurso que apunta a un objeto individual, similar a una fila en una base de datos. Puede tener recursos hijos que pueden ser tanto subdocumentos como colecciones. Por ejemplo:
```bash
GET /cats/1 -> Documento individual para el gato 1
GET /cats/1/kittens -> Todos los gatitos pertenecientes al gato 1
GET /cats/1/kittens/1 -> Gatito 1 del gato 1
```
## Controlador
Un recurso controlador es como un procedimiento, y generalmente se usa cuando un recurso no puede mapearse a las funciones estándar de CRUD (crear, recuperar, actualizar y eliminar).

Los nombres de los controladores aparecen como el último segmento en una ruta URI sin recursos hijos. Si el controlador requiere parámetros, estos típicamente se incluirían en la cadena de consulta. Ejemplos:
```bash
    POST /cats/1/feed -> Alimentar al gato 1
    POST /cats/1/feed?food=fish -> Alimentar al gato 1 con un pez
```
Al definir un nombre de controlador, siempre debemos usar un verbo, que indica una acción o un estado, como "feed" o "send".

## Store

Una Store es un repositorio de recursos administrado por el cliente. Permite al cliente agregar, recuperar y eliminar recursos. A diferencia de una colección, una Store nunca generará una nueva URI; usará la que especifique el cliente. Un ejemplo que agrega un nuevo gato a nuestra Store sería:
```bash
    PUT /cats/2 -> Esto agregaría un nuevo gato a la Store con un ID de 2.
```
Si hubiéramos publicado el nuevo gato omitiendo el ID en una colección, la respuesta necesitaría incluir una referencia al documento recién definido para que pudiéramos interactuar con él más tarde. Al igual que con los controladores, debemos usar un sustantivo en plural para los nombres de las Stores.

# Nombres de funciones CRUD

Al diseñar URIs REST óptimas, nunca usamos un nombre de función CRUD como parte de la URI; en su lugar, usamos un verbo HTTP. Por ejemplo:

```bash
DELETE /cats/1234
```
No incluimos el verbo en el nombre del método, ya que este está especificado por el verbo HTTP. Las siguientes URIs se considerarían un anti-patrón:

```bash
GET /deleteCat/1234
DELETE /deleteCat/1234
POST /cats/1234/delete
```

# Verbos HTTP

Los verbos HTTP comúnmente utilizados son:
- GET
- POST
- PUT
- PATCH
- DELETE
- HEAD
- OPTIONS

Cada uno de estos métodos tiene una semántica bien definida dentro del contexto de nuestra API REST, y su correcta implementación ayudará a los usuarios a entender la intención.

## GET

El método GET se usa para recuperar un recurso y nunca debe usarse para modificar una operación, como actualizar un registro. Típicamente, no se pasa un cuerpo con una solicitud GET; sin embargo, no es inválido hacerlo en una solicitud HTTP.

Solicitud:

```bash
GET /v1/cats HTTP/1.1
```
Respuesta:

```bash
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: xxxx
{"name": "Fat Freddie's Cat", "weight": 15}
```
## POST 

El método POST se usa para crear un nuevo recurso en una colección o para ejecutar un controlador. Es una acción típicamente no idempotente, lo que significa que múltiples posts para crear un elemento en una colección crearán múltiples elementos, no se actualizarán después de la primera llamada.

El método POST siempre se usa cuando se llaman controladores, ya que estas acciones se consideran no idempotentes.

Solicitud:

```bash
POST /v1/cats HTTP/1.1
Content-Type: application/json
Content-Length: xxxx
{"name": "Felix", "weight": 5}
```
Respuesta:

```bash
HTTP/1.1 201 Created
Content-Type: application/json
Content-Length: 0
Location: /v1/cats/12343
```
## PUT 

El método PUT se usa para actualizar un recurso mutable y siempre debe incluir el localizador del recurso. Las llamadas al método PUT también son idempotentes, es decir, múltiples solicitudes no modificarán el recurso a un estado diferente del primer llamado.

Solicitud:

```bash
PUT /v1/cats HTTP/1.1
Content-Type: application/json
Content-Length: xxxx
{"name": "Thomas", "weight": 7}
```
Respuesta:

```bash
HTTP/1.1 201 Created
Content-Type: application/json
Content-Length: 0
```
## PATCH 

El verbo PATCH se usa para realizar una actualización parcial. Por ejemplo, si solo quisiéramos actualizar el nombre de nuestro gato, podríamos hacer una solicitud PATCH solo con los detalles que deseamos cambiar.

Solicitud:

```bash
PATCH /v1/cats/12343 HTTP/1.1
Content-Type: application/json
Content-Length: xxxx
{"weight": 9}
```
Respuesta:

```bash
HTTP/1.1 204 No Body
Content-Type: application/json
Content-Length: 0
```
En mi experiencia, las actualizaciones PATCH rara vez se usan. La convención general es usar PUT y actualizar el objeto completo, lo que no solo facilita escribir el código, sino también hacer que la API sea más sencilla de entender.

## DELETE 

El verbo DELETE se usa cuando queremos eliminar un recurso. Generalmente, pasamos el ID del recurso como parte de la ruta en lugar de en el cuerpo de la solicitud. De esta manera, tenemos un método consistente para actualizar, eliminar y recuperar un documento.

Solicitud:

```bash
DELETE /v1/cats/12343 HTTP/1.1
Content-Type: application/json
Content-Length: 0
```
Respuesta:

```bash
HTTP/1.1 204 No Body
Content-Type: application/json
Content-Length: 0
```
## HEAD 

Un cliente usaría el verbo HEAD cuando desea recuperar solo los encabezados de un recurso sin el cuerpo. El verbo HEAD se usa típicamente en lugar de un GET cuando un cliente solo quiere verificar si un recurso existe o leer los metadatos.

Solicitud:

```bash
HEAD /v1/cats/12343 HTTP/1.1
Content-Type: application/json
Content-Length: 0
```
Respuesta:

```bash
HTTP/1.1 200 OK
Content-Type: application/json
Last-Modified: Wed, 25 Feb 2004 22:37:23 GMT
Content-Length: 45
```
## OPTIONS 

El verbo OPTIONS se usa cuando un cliente desea recuperar las interacciones posibles para un recurso. Típicamente, el servidor devolverá un encabezado "Allow", que incluirá los verbos HTTP que pueden usarse con este recurso.

Solicitud:

```bash
OPTIONS /v1/cats/12343 HTTP/1.1
Content-Length: 0
```
Respuesta:

```bash
HTTP/1.1 200 OK
Content-Length: 0
Allow: GET, PUT, DELETE
```

# Diseño de consulta URI

Es perfectamente aceptable utilizar una cadena de consulta (query string) como parte de una llamada API; sin embargo, se recomienda no usarla para pasar datos al servicio. En su lugar, la consulta debe usarse para realizar acciones como:
- Paginación
- Filtrado
- Ordenamiento

Si necesitamos hacer una llamada a un controlador, discutimos anteriormente que deberíamos usar una solicitud POST ya que es probable que sea una solicitud no idempotente. Para pasar datos al servicio, deberíamos incluir los datos dentro del cuerpo. Sin embargo, podríamos usar una cadena de consulta para filtrar la acción del controlador:

Solicitud:

```bash
POST /sendStatusUpdateEmail?$group=admin
{
  "message": "All services are now operational\nPlease accept our apologies for any inconvenience caused.\nThe Kitten API team"
}
```
En el ejemplo anterior, enviaríamos un correo electrónico de actualización de estado con el mensaje incluido en el cuerpo de la solicitud. Debido a que estamos utilizando el filtro de grupo pasado en la cadena de consulta, podríamos restringir la acción de este controlador a enviar solo al grupo de administradores.

Si hubiéramos agregado el mensaje a la cadena de consulta y no hubiéramos pasado un cuerpo de mensaje, podríamos estar causando dos problemas. El primero es que la longitud máxima para una URI es de `2083 caracteres`. El segundo es que, generalmente, una solicitud POST siempre incluiría un cuerpo de solicitud. Aunque esto no es un requisito de la especificación HTTP, se esperaría este comportamiento por parte de la mayoría de los usuarios.

# Códigos de respuesta

Cuando diseñamos una buena API, deberíamos usar códigos de estado HTTP para indicar al cliente el éxito o fracaso de la solicitud.
Actualmente, hay un consenso general de que esta es una buena práctica, ya que permite al cliente determinar de inmediato el estado de una solicitud sin tener que analizar el cuerpo de la solicitud para obtener más información. En caso de fallo, las APIs que siempre devuelven una respuesta 200 OK al usuario con un cuerpo de mensaje que contiene información adicional no son una buena práctica, ya que requieren que el cliente inspeccione el cuerpo para determinar el resultado. También significa que el cuerpo del mensaje contendrá información adicional más allá del objeto que debería representar. 

Considera la siguiente mala práctica:

- Cuerpo de solicitud malo:

```bash
POST /kittens
RESPONSE HTTP 200 OK
{
  "status": 401,
  "statusMessage": "Bad Request"
}
```
- Solicitud exitosa:

```bash
POST /kittens
RESPONSE HTTP 201 CREATED
{
  "status": 201,
  "statusMessage": "Created",
  "kitten": {
    "id": "1234334dffdf23",
    "name": "Fat Freddy's Cat"
  }
}
```
Imagina que estás escribiendo un cliente para la solicitud anterior, necesitas agregar lógica a tu aplicación para verificar el nodo de estado en la respuesta antes de poder leer y procesar el gato devuelto.

Ahora considera algo aún peor:

- Fallo aún peor:

```bash
POST /kittens
RESPONSE HTTP 200 OK
{
  "status": 400,
  "statusMessage": "Bad Request"
}
```
- Éxito aún peor:

```bash
POST /kittens
RESPONSE HTTP 200 OK
{
  "id": "123434jhjh3433",
  "name": "Fat Freddy's Cat"
}
```
Si el autor de tu API hubiera hecho algo como el ejemplo anterior, necesitarías verificar si la respuesta que se ha devuelto es un error o el gato que esperabas. La cantidad de "WTFs" por minuto que pronunciarías mientras codificas un cliente para esta API no te haría ganar cariño del autor. Estos ejemplos pueden parecer extremos, pero hay casos como este en el mundo real. En algún momento de mi carrera, estoy bastante seguro de haber cometido semejante error, pero entonces no había leído este libro.

Lo que el autor, con la mejor intención, ha hecho es intentar tomar los códigos de estado HTTP demasiado literalmente. La W3C [RFC2616](https://www.w3.org/Protocols/rfc2616/rfc2616-sec6.html#sec6.1.1) establece que el código de estado HTTP se relaciona con el intento de entender y satisfacer la solicitud ; sin embargo, esto es un poco ambiguo cuando se observan algunos de los códigos de estado individuales. 
El consenso moderno es que está bien usar los códigos de estado HTTP para indicar el estado de procesamiento de una solicitud API, no solo la capacidad del servidor para procesar la solicitud. Considera cómo podríamos mejorar estas solicitudes implementando este enfoque.

- Un buen ejemplo de un fallo:

```bash
POST /kittens
RESPONSE HTTP 400 BAD REQUEST
{
  "errorMessage": "Name should be between 1 and 256 characters in length and only contain [A-Z] - ['-.]"
}
```
- Un buen ejemplo de éxito:

```bash
POST /kittens
RESPONSE HTTP 201 CREATED
{
  "id": "123hjhjh2322",
  "name": "Fat Freddy's Cat"
}
```
Esto es mucho más semántico; el usuario solo necesita leer la respuesta en caso de fallo si requiere más información. Además, podemos proporcionar un objeto de error estándar que se utilice en todos los puntos finales de nuestra API, lo que proporciona información adicional pero no obligatoria para determinar por qué falló una solicitud. Veremos los objetos de error en breve, pero por ahora, veamos los códigos de estado HTTP con más detalle.

## 2xx Éxito
Los códigos de estado 2xx indican que la solicitud del cliente ha sido recibida y comprendida con éxito.

### 200 OK
Este es un código de respuesta genérico que indica que la solicitud ha tenido éxito. La respuesta que acompaña este código generalmente es:

  - GET: Una entidad que corresponde al recurso solicitado.
  - HEAD: Los campos de cabecera correspondientes al recurso solicitado sin el cuerpo del mensaje.
  - POST: Una entidad que describe o contiene el resultado de la acción.

### 201 Created
La respuesta Created se envía cuando una solicitud tiene éxito y el resultado es que se ha creado una nueva entidad. Junto con la respuesta, es común que la API devuelva un encabezado Location con la ubicación de la nueva entidad creada:

```bash
201 Created  
Location: https://api.kittens.com/v1/kittens/123dfdf111
```
Es opcional devolver un cuerpo de objeto con este tipo de respuesta.

### 204 No Content
Este estado informa al cliente que la solicitud ha sido procesada con éxito; sin embargo, no habrá un cuerpo de mensaje en la respuesta. Por ejemplo, si el usuario realiza una solicitud DELETE a la colección, la respuesta puede devolver un estado 204.

## 3xx Redirección
Los códigos de estado de la clase 3xx indican que el cliente debe realizar una acción adicional para completar la solicitud. Muchos de estos códigos son utilizados por CDNs y otras técnicas de redirección de contenido; sin embargo, el código 304 puede ser especialmente útil al diseñar nuestras APIs para proporcionar retroalimentación semántica al cliente.

### 301 Moved Permanently
Este código indica al cliente que el recurso solicitado ha sido movido permanentemente a una ubicación diferente. Aunque se utiliza tradicionalmente para redirigir una página o recurso desde un servidor web, también puede ser útil al crear nuestras APIs. En el caso de que renombremos una colección, podríamos usar un redireccionamiento 301 para enviar al cliente a la ubicación correcta. Sin embargo, esto debe usarse como una excepción y no como la norma. Algunos clientes no siguen automáticamente los redireccionamientos 301, lo que añade complejidad para los consumidores.

### 304 Not Modified
Esta respuesta es generalmente utilizada por un CDN o servidor de caché y se establece para indicar que la respuesta no ha sido modificada desde la última llamada a la API. Está diseñada para ahorrar ancho de banda y la solicitud no devolverá un cuerpo, pero sí devolverá un encabezado Content-Location y Expires.

## 4xx Error del Cliente
En caso de un error causado por el cliente, no por el servidor, este último devolverá una respuesta 4xx y siempre incluirá una entidad que proporcione más detalles sobre el error.

### 400 Bad Request
Este código indica que el servidor no pudo entender la solicitud debido a un formato incorrecto o un fallo en la validación del dominio (datos faltantes o una operación que causaría un estado no válido).

### 401 Unauthorized
Este código indica que la solicitud requiere autenticación del usuario e incluirá un encabezado WWW-Authenticate que contiene un desafío aplicable al recurso solicitado. Si el usuario ha incluido las credenciales necesarias en el encabezado WWW-Authenticate, la respuesta debería incluir un objeto de error que contenga información diagnóstica relevante.

### 403 Forbidden
El servidor ha entendido la solicitud, pero se niega a cumplirla. Esto puede deberse a un nivel de acceso incorrecto al recurso, no a que el usuario no esté autenticado.
Si el servidor no desea hacer público que una solicitud no puede acceder a un recurso debido al nivel de acceso, es permisible devolver un estado 404 Not Found en lugar de esta respuesta.

### 404 Not Found
Esta respuesta indica que el servidor no ha encontrado nada que coincida con el URI solicitado.
No se da ninguna indicación de si la condición es temporal o permanente. Es permisible que el cliente realice múltiples solicitudes a este punto final, ya que el estado podría no ser permanente.

### 405 Method Not Allowed
El método especificado en la solicitud no está permitido para el recurso indicado por el URI. Esto puede suceder cuando el cliente intenta modificar una colección enviando un POST, PUT o PATCH a una colección que solo permite la recuperación de documentos.

### 408 Request Timeout
El cliente no produjo una solicitud dentro del tiempo que el servidor está dispuesto a esperar. El cliente puede repetir la solicitud sin modificaciones más tarde.

## 5xx Server Error
Los códigos de estado de respuesta dentro del rango 500 indican que algo ha salido mal, el servidor lo sabe y lamenta la situación.
El RFC aconseja que se devuelva una entidad de error en la respuesta que explique si esto es permanente o temporal, y que contenga una explicación del error. Cuando examinemos el capítulo sobre seguridad, veremos la recomendación de no dar demasiada información en los mensajes de error, ya que este estado podría haber sido provocado por un usuario con la intención de comprometer tu sistema, y al devolver cosas como una traza de pila u otra información interna con un error 5xx, podrías ayudar a comprometer tu sistema. Con esto en mente, actualmente es común que un error 500 simplemente devuelva algo muy genérico.

### 500 Internal Server Error
Un mensaje de error genérico que indica que algo no salió como se esperaba.

### 503 Service Unavailable
El servidor no está disponible temporalmente debido a sobrecarga o mantenimiento.
Existe un patrón útil que puedes implementar para evitar fallos en cascada en caso de un mal funcionamiento, donde el microservicio monitoriza su estado interno y, en caso de fallo o sobrecarga, se niega a aceptar la solicitud y señala esto al cliente de inmediato. Veremos este patrón en más detalle en el capítulo xx; sin embargo, este es probablemente el escenario donde querrás devolver un estado 503. Esto también puede usarse como parte de tus verificaciones de salud.

# HTTP Headers
Los encabezados de solicitud son una parte muy importante del proceso de solicitud y respuesta HTTP, y al implementar un enfoque estándar ayudas a tus usuarios a hacer la transición de una API a otra.
En esta subsección, no cubriremos todos los encabezados posibles que puedes usar en tu API, pero analizaremos los encabezados más comunes. Para obtener información completa sobre el protocolo HTTP, consulta el RFC 7231 RFC 7231, que contiene una descripción completa del estándar actual.

## Encabezados de solicitud estándar
Los encabezados de solicitud proporcionan información adicional para la solicitud y la respuesta de tu API. Piénsalos como metadatos para la operación. Pueden utilizarse para ampliar otros datos de la respuesta que no pertenecen al cuerpo en sí, como la codificación del contenido. También pueden ser utilizados por el cliente para proporcionar información que ayude al servidor a procesar la respuesta. Siempre que sea posible, deberíamos usar los encabezados estándar, ya que esto ofrece consistencia a tu usuario y les proporciona un estándar común en múltiples puntos finales de diferentes proveedores.



