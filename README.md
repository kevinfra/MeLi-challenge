# Mercadolibre Challenge

## Descripción del Challenge:

El objetivo fue hacer funcionar una API con MySQL y NGINX, las cuales se encontraban en containers de Docker y podían inicializarse con Docker Compose.

Teniendo todo levantado, NGINX recibe los requests en el puerto 8080 y los redirige a la API en el puerto 3000.

La API contiene los siguientes endpoints separados por secciones:

### Items

- `GET /item/:id`: Recibe un id de item y devuelve el item con ese ID. En caso de no existir, devuelve error.
- `POST /item`: Recibe por POST el  (`name`) y descripción (`description`) de un item y lo agrega a la base de datos. Ejemplo:
`curl -d '{"name":"item1", "description":"first item"}' -H "Content-Type: application/json" -X POST http://localhost:8080/item`
- `DELETE /item/:id`: Recibe el id de un item y lo borra de la base de datos.
- `GET /item`: Devuelve un item nulo (`"id": 0, "name": "", "description": ""`)

## Google Drive

- `POST /auth-for-drive`: La primera vez devuelve un JSON donde la key es una url para autorizar a la app. Luego de autorizada, hay que enviar el token como parametro a este mismo endpoint de la forma `/auth-for-drive?token=TOKEN`. Ese post devuelve otro JSON indicando si el token era valido o no. Las llamadas posteriores solo devuelven un JSON indicando que el usuario ya está autenticado. Es necesario hacer el request a este endpoint para que los otros endpoints de Drive no devuelvan error de autenticación.
- `GET /search-in-doc/:id`: Recibe el id y un parámetro `word`. Toma el id recibido y lo usa para buscar en Drive el archivo con ese id. Si lo encuentra, busca en el contenido del archivo si existe la palabra pasada por parámetro. Si encuentra la palabra, retorna 200. Si no, 404.
- `POST /file`: Recibe como parámetros `"titulo"` y `"descripcion"`, datos con los cuales crea un archivo en el drive del usuario con el correspondiente título y descripción. En caso de ser exitoso, devuelve un json con los datos del archivo creado y su id correspondiente en la base de datos.

## Dificultades encontradas:

- Nunca había usado Go.
- La correcta configuración de NGINX para que funcione entre los containers de Docker me tomó más tiempo del que esperaba, pues tuve problemas con la redirección de los requests.
- Tampoco había usado nunca la API de Drive, lo que implicó que parte del tiempo lo dedique a entender su funcionamiento.
- Estos 3 problemas fueron las mayores dificultades encontradas, y todas se solucionaron con tiempo, lectura de las documentaciones, y pruebas.
