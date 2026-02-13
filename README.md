# url-shortener

Proyecto de practica para desarrollo de backends con Golang.

La forma de acortar las url es mediante una base de datos que relacionara la url pasada por el usuario con un id de 5 caracteres en base62. Se eligio 5 caracteres dado que permiten un total de 916.132.832 combinaciones posibles en dicha base. Esta cantidad es mas que suficiente para incluso un caso real de acortadores. Si se expandiera a 6 caracteres se obtendrian 56.8B de combinaciones, lo que aseguraria nunca quedarse sin posibles id.

No se usara ningun framework de enrutamiento como son Gin, Echo o Fiber. Por fines educativos se realizara todo desde cero, a excepcion de los drivers de bases de datos.

Este tipo de servicios proveen la capacidad de analiticas sobre, por ejemplo, visitas en los links.

Cosas clave a manejar por el backend:

1. Validacion de entrada:
- Que el formato sea de URL ('http://' o 'https://')
- Evitar bucles, es decir no acortar url's ya acortadas por el propio sistema. Blacklist del dominio propio.
- Limite de longitud, para evitar que alguien guarde el texto completo de El Quijote en la base de datos haciendolo pasar por una URL.

2. Manejo de concurrencia:
- Indices unicos en la db. No confiar solo en el codigo y que la db sea la ultima linea de defensa.
- Transacciones atomicas, para evitar perder datos si existiera un trafico alto.

3. Async para las analiticas:
Cuando alguien hace un request sobre un link la redireccion sera inmediata, mientras que el contar la visita se realice en una goroutine, es decir en segundo plano.

4. Timeouts:
- Las consultas a la db deben ser con 'context.WithTimeout' para evitar esperas y liberar recursos en unos 5 segundos por ejemplo.
- Para verificar que la url que el usuario pida acortar exista se le hara un ping, con un timeout de por ejemplo 2000ms.

5. Shutdown:
Cuando el servidor deba apagarse, por el motivo que sea, debera hacerlo en los siguientes pasos:
- Dejar ded aceptar nuevas conexiones.
- Esperar que terminen todas las conexiones activas (con un limite de tiempo, ej: 10s).
- Cerrar la conexion con la db
- Apagarse.


Otros conceptos:
- configuracion con variables de entorno.
- logs estructurados.
- rate limiting por ip.
- eliminar enlaces sin uso.