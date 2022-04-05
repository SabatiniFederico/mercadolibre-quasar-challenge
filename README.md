# Mercadolibre Quasar

Desafío "Fuego de Quasar" - de Federico Sabatini

## Inicio
requiere versión **GO 1.8**

Para iniciar la app en local

> $ go run server.go
> 
Para ejecutar los tests

> $ go test ./... -cover

## Estructura
```
Estructura del proyecto
├── controller
├── entity
├── lib
│ └── trilateration
├── router
└── service
```
- **Controller** recibe y le otorga el primer tratamiento a los request http, para luego enviarlos al servicio apropiado.
- **Entity** contiene las estructuras que representan las request y response
- **lib / trilateration** Es mi propia implementación de trilateración para ubicar la posición de la nave espacial
- **Router** obtiene los endpoints y los envía al controlador.
- **Service**  contiene la capa de negocio. Aquí se define toda la lógica necesaria para calcular el mensaje secreto de la Nave del imperio.

## Consideraciones 
### GetLocation ### 
Recibe una lista de distancias y a partir de un algoritmo de trilateración resuelve la posición. 

Se decidió realizar una implementación propia en *2D* del algoritmo, en vez de utilizar una librería externa. En caso de que se quiera cambiar a futuro, se puede remover la librería y hacer un `go get` de la librería deseada. 

Se basó en el siguiente artículo de la wiki: https://en.wikipedia.org/wiki/True-range_multilateration#Two_Cartesian_dimensions

Aquí se propone una forma mucho más fácil de calcular trilateración que la estándar, con la excepción de que se requiere que el punto 1 de interés este en el origen y el segundo punto en eje de abscisas.

Cualquier conjunto de puntos pueden ser trasladado para cumplir dichas condiciones usando transformaciones lineales. La idea matemática es la siguiente:

1. Trasladar el conjunto para que el primer punto (kenobi) se encuentre en el *(0,0)*.
2. Rotar el conjunto para que el segundo punto (skywalker) quede ubicado en el eje de las abscisas.
3. Calculamos *(x,y)* con las fórmulas propuestas por el artículo:

![{\displaystyle {\begin{aligned}x&={\frac {r_{1}^{2}-r_{2}^{2}+U^{2}}{2U}}\\[4pt]y&=\pm {\sqrt {r_{1}^{2}-x^{2}}}\end{aligned}}}](https://wikimedia.org/api/rest_v1/media/math/render/svg/ebcc6eb379df69ed08e8e83b5c4488c83481b3e3)

4. Validamos con el tercer punto (sato) cuál de las 2 soluciones de *(y)* es la correcta.
5. En caso de que haya solución rotamos y trasladamos el resultado en sentido inverso.

### GetMessage ###
Recibe mensajes incompletos y resuelve el mensaje original

El defasaje del mensaje se decidió interpretarlo el cómo un *"Ruido"* de strings vacíos que llegan a la izquierda del mensaje. 

La solución involucra buscar el slice más corto, usarlo como parámetro para determinar el largo del mensaje y finalmente cortar los otros mensajes. 
Luego se revisa que el mensaje sea descifrable es decir que no tenga contradicciones, por ejemplo los mensajes:

```
["Luke","yo","soy","tu","padre]
["Usa","la","fuerza","Anakin","Skywalker]
```

No tienen solución válida, ya que tendrías que elegir si descartar el primer mensaje o el segundo, contradicciones en una palabra se interpretan como basura y descartan el mensaje de ser una respuesta válida. 

### Servicio "Topsecret_split ###

Para el servicio topsecret_split se posee una lista vacía de satélites y se utiliza *validator* para revisar que los nombres sean efectivamente *"kenobi", "skywalker o "sato"* en minúscula

Si se agrega información de un satélite que ya existía en la lista, este mismo se actualizara con los nuevos datos (distancia y mensaje)

## Endpoints
  
**Topsecret**

Obtiene la posición y mensaje completo de la nave destinataria que haya enviado el mensaje. Algunas posiciones no son válidas.
A continuación se provee una request de ejemplo.

`POST -> /topsecret`

Ejemplo de request body:
```json
{
	"satellites": [
		{
			"name": "kenobi",
			"distance": 650.0,
			"message": ["", "", "", "", "es", "", "challenge", "", "mercado", "libre"]
		},
		{
			"name": "skywalker",
			"distance": 350.0,
			"message": ["este", "es", "", "", "de", "", "libre"]
		},
		{
			"name": "sato",
			"distance": 680.0735,
			"message": ["", "", "este", "", "un", "","de", "mercado", "libre"]
		}
	]
}
```

**Topsecret_split**

Similar a TopSecret, solo que aquí el nombre llega por request param.

`POST -> /topsecret_split/kenobi`

```
{
    "distance": 650.0,
    "message": ["join", "the", "dark", "side"]
}
```

`GET -> /topsecret_split/`
Ejemplo de una respuesta exitosa:

```json
{
	"position": {
		"x": 100,
		"y": -450.00000000000006
	},
	"message": "este es un challenge de mercado libre"
}
```

### Iteraciones futuras ###
Esta sección destaca una lista de posibles mejoras a la hora de iterar el proyecto productivo.

- Se podría agregar algún Serverless Redis o DynamoDB para gestionar los nombres de los satélites y sus posiciones constantes.
- Siempre es un plus agregar documentación con Swagger API
- La capa de servicio tiene un coverage de 95.2%, pero estaría bueno tener tests de integración en la capa del controlador para completar.
- Crear algún custom Logger para tener mejor control del flujo de la aplicación.


