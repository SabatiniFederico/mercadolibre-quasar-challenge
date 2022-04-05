# mercadolibre Quasar

Desafio "Fuego de Quasar" - de Federico Sabatini

## Inicio
requiere versión **GO 1.8**

Para iniciar la app

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
- **Entity**  contiene las estructuras que representan las request y response
- **lib / trilateration** Es mi propia implementación de trilateración para ubicar la posición de la nave espacial
- **Router** obtiene los endpoints y los envia al controlador.
- **Service**  contiene la capa de negocio. Aqui se define toda la lógica necesaria para calcular el mensaje secreto de la Nave del imperio.

## Consideraciones 
**GetLocation**: Recibe una lista de distancias y a partir de un algoritmo de trilateración resuelve la posición. 

Se decidio realizar una implementación propia en *2D* del algoritmo, en ves de utilizar una libreria externa. En caso de que se quiera cambiar a futuro, se puede remover la libreria y hacer un `go get` de la libreria deseada. la solución esta basada en el siguiente articulo de la wiki:

https://en.wikipedia.org/wiki/True-range_multilateration#Two_Cartesian_dimensions

Aqui se propone una forma mucho más fácil de calcular trilateración que la estandar, con la excepción de que se requiere que el punto 1 de interés este en el origen y el segundo punto en eje de abscisas.

Cualquier conjunto de puntos pueden ser transladado para cumplir dichas condiciones utilizando transformaciones lineales. La idea matematica es la siguiente:

1. Transladar el conjunto para que el primer punto (kenobi) se encuentre en el *(0,0)*.
2. Rotar el conjunto para que el segundo punto (skywalker) quede ubicado en el eje de las abscisas.
3. Calculamos *(x,y)* con las formulas propuestas por el articulo:

![{\displaystyle {\begin{aligned}x&={\frac {r_{1}^{2}-r_{2}^{2}+U^{2}}{2U}}\\[4pt]y&=\pm {\sqrt {r_{1}^{2}-x^{2}}}\end{aligned}}}](https://wikimedia.org/api/rest_v1/media/math/render/svg/ebcc6eb379df69ed08e8e83b5c4488c83481b3e3)

4. Validamos con el tercer punto (sato) cual de las 2 soluciones de *(y)* es la correcta.
5. En caso de que haya solución rotamos y transladamos el resultado en sentido inverso.

**GetMessage**: Recibe mensajes incompletos y resuelve el mensaje original

Se decidio interpretar el defasaje como strings vacios que llegan a la izquierda del mensaje. 


## Endpoints
  
**Topsecret**

Obtiene la posición y mensaje completo de la nave destinataria que haya enviado el mensaje. algunas posiciones no son validas

a continuación se provee una request de ejemplo.

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