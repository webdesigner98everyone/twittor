# twittor
para la ejecucion del siguiente Software enfocado en Backend necesitaremos el apoyo de dos herramientas fundamentales

1. MogoBD Compass
2. Postman

Postman nos ayudara a la ejecucion de pruebas backend de nuestro software llamado twittor.

Recomendaciones MONGODB:

1. Una vez descarguemos nuestro MongoDB Compass, vamos a realizar la conexion a la BD que el software tiene y es el siguiente: (mongodb+srv://root:root@cluster0.1xrlq.mongodb.net/admin?authSource=admin&replicaSet=atlas-h7w62z-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true)
2. cuando estemos dentro de nnuestra BD, observaremos que existe una BD llamada twittor, esta contiene 3 tablas las cuales por el momento estan disponibles 2 que son las de usuario y tweet
3. actualemente el software esta probado en backend con las opciones de registro,Login, Ver perfil, ModificarPerfil en la tabla usuarios 
4. en la tabla Tweet el backend esta probado con agregar tweets, LeerTweets o buscarlos y borrarlos por medio del IDTweets que se almacena en esta tabla.

Recomendaciones POSTMAN:
1. Nuestro postman contiene por ahora 7 colecciones las cuales se mencionaron en los puntos 3 y 4 de las recomendaciones de mongo, estas colecciones van por medio de POST, GET, PUT y DEL
2. En el repos esta una carpeta llamda POSTMAN en la que se ubica nuestro Postman en forma de JSON para la importacion en la app local

TENER ENCUENTA QUE TAMBIEN EL SOFTWARE ESTA DISPONIBLE EN HEROKU.
