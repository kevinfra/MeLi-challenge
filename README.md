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

## Dificultades encontradas:

- Nunca habia usado Go, así que tuve que aprender desde 0.
- La correcta configuración de NGINX para que funcione entre los containers de Docker me tomó más tiempo del que esperaba, pues tuve problemas con la redirección de los requests.