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
- **Controller** as
- **Entity**  contiene las estructuras que representan las request y response
- **lib / trilateration** Es mi propia implementación de trilateración para ubicar la posición de la nave espacial
- **Router** obtiene los endpoints y los envia al controlador.
- **Service**  contiene la capa de negocio. Aqui se define toda la lógica necesaria para calcular el mensaje secreto de la Nave del imperio.


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
			"message": ["", "", "", "", "", "es", "", "challenge", "", "mercado", "libre"]
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