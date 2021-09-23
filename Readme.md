# Twittor

## Estructura de directorios en go

- Si nuestros proyectos se crearan en una carpeta diferente a donde se instalo el programa por defecto, creamos una nueva carpeta donde queramos harerlo con el nombre ***go***
- Dentro de la carpeta **go** creamos las carpetas **bin, pkg y src**
- Dentro de la carpeta **src** creamos la carpeta **github.com**
- Dentro de la carpeta **github.com** creamos una carpeta con el nombre de nuestro usuario de *github*
- Dentro de la carpeta **nombreUsuarioGithub** creamos una carpeta con el nombre de nuestro repositorio de github  **nombreRepositorioGithub**
- En esta carpeta nueva iran todas nuestras carpetas base de ya nuestro proyecto.
  - bd
  - handlers
  - jwt
  - middlew
  - models
  - routers
  - uploads

## Variables de entorno

Nos vamos a configurar las variables de entorno y agregamos:

- Varibles de usuario
  - Creamos la variable ***GOPATH*** y le asignamos la ruta donde esten los proyectos que realizaremos ***go*** que en mi caso es `D:\practica\go`
  - En ***Path*** agregamos la ruta del ***bin*** de nuestros proyectos de go `D:\practica\go\bin`
- Variables de sistema
  - Igual que en caso anterior creamos la variable ***GOPATH*** y le asignamos la ruta donde esten los proyectos que realizaremos ***go*** que en mi caso es `D:\practica\go`
  - En ***Path*** igual que antes agregamos la ruta del ***bin*** de nuestros proyectos de go `D:\practica\go\bin`
  - En ***Path*** tambien agregamos la ruta ***bin*** de donde se instalo el programa `C:\Program Files\Go\bin`

## Iniciamos el repositorio git

- Creamos el archivo ***Readme.md***
- continuamos con los siguiente codigos en consola

```git
git init
git add Readme.md
git commit -m "Primer commit"
git branch -M main
git remote add origin git@github.com:miguelmalagaortega/twittor.git
git push -u origin main
```

## Iniciamos el repositorio de Heroku

- Primer abrimos sesion desde la pagina de Heroku para luego poder logearnos desde VSCode.
- Creamos la aplicacion en Heroku, al terminar no olvidar ir  a:
  - Nuestra aplicacion, en este caso ***twittordcn***
  - De ahi a la seccion ***Setting***
  - En las opciones elegimos ***Add buildpack*** y en la ventana elegimos el lenguaje, en ese caso ***go***
- Despues descargamos el cliente de heroku.
- Luego ponemos en consola el siguiente codigo: `heroku login`, esto nos llevara a cargar las credenciales de Heroku.
- Antes de esto ya se debio haber hecho la creacion de git init, lo cual se debio realizar en la parte *Iniciamos el repositorio git*.
- Agregamos el ***remote*** de Heroku `heroku git:remote -a twittordcn`
- Subimos la aplicacion a Heroku `git push heroku main`

***Nota:*** Antes de poder subir todo a Heroku debemos de crear algunos ***archivos de go*** para que se pueda hacer una configuracion adecuada y no tener errores.

## Creacion de los archivos base

- Desde consola creamos el archivo go.mod con el siguiente codigo

`go mod init github.com/nombreUsuarioGithub/nombreRepositorioGithub`

tendra una estructura parecida a esta luego de crearce

```go
module github.com/miguelmalagaortega/twittor

go 1.16

require (

)
```

- Tambien creamos el archivo ***Procfile*** que es un archivo de configuracion de HEROKU, dentro colocaremos lo siguiente

```go
web: nombreRepositorioGithub
<!-- Ejemplo -->
web: twittor
```

- Por ultimo crearemos el archivo ***main.go*** y pondremos como base

```go
package main

func main(){
  // codigo
}
```

***Nota:*** aun asi tengamos estos archivos creados es posible que aun nos de un error al subir los archivos, si es asi esto se resolvera luego de agregar algo mas de codigo. En ese momento volvamos a intentar subir todo.

## Dependencias que usaremos

- Iniciamos agregando las dependencias a go con el comando **go get** y seguido los url

  - go.mongodb.org/mongo-driver/mongo, *El paquete mongo proporciona una API de controlador MongoDB para Go*
  - go.mongodb.org/mongo-driver/mongo/options
  - go.mongodb.org/mongo-driver/bson, *Package bson es una biblioteca para leer, escribir y manipular BSON. BSON es un formato de serialización binario que se utiliza para almacenar documentos y realizar llamadas a procedimientos remotos en MongoDB.*
  - go.mongodb.org/mongo-driver/bson/primitive, *La primitiva de paquete contiene tipos similares a las primitivas Go para tipos BSON que no tienen representaciones primitivas Go directas.*
  - golang.org/x/crypto/bcrypt, *El paquete bcrypt implementa el algoritmo de hash adaptativo bcrypt de Provos y Mazières*
  - github.com/gorilla/mux, *El paquete gorilla/mux implementa un enrutador de solicitudes y un despachador para hacer coincidir las solicitudes entrantes con su respectivo controlador*
  - github.com/rs/cors, *CORS es un net/httpcontrolador que implementa la especificación Cross Origin Resource Sharing W3 en Golang*
  - github.com/dgrijalva/jwt-go, *Esta biblioteca admite el análisis y la verificación, así como la generación y firma de JWT. Los algoritmos de firma admitidos actualmente son HMAC SHA, RSA, RSA-PSS y ECDSA, aunque hay ganchos para agregar los suyos propios*, **NOTA** Este paquete ya no esta en manenimiento, se recomienda usar en su lugar github.com/golang-jwt/jwt

- Luego de agregar todas las dependencias veremos que se crea el archivo ***go.sum*** y tambien se veran agregados al archivo ***go.mod***

## Programando el archivo de conexion a la base de datos

1. En la carpeta ***bd*** creamos nuestros archivos para la base de datos.

- El primer archivo sera ***conexionBD.go***
- El codigo sera el siguiente:

```go
package bd

import (
  // En go no existen variables globales, no se usan
  // Los contextos nos sirven para comunicar informacion entre ejecucion y ejecucion ya demas nos permite setear una seria de valor como por ejemplo un timeout, que nos ayuda a usar un tiempo de espera y seguir con el programa sin que este se cuelgue
  "context"
  // para grabar nombre y texto dentro del log de ejecucion
  "log"

  // importamos algunos de los paquetes que agregamos antes con go get, en este caso para manejar la base de datos son necesarios estos dos
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

// ----------------------------------------------------------
// crearemos algunas variables

// Variable de uso externo, por eso la iniciamos con MAYUCULA
// con esta variable que toma el valor de una funcion se va a ejecutar la conexion a la base de datos y poder realizar las operaciones
var MongoCN = ConectarBD()

// Variable de uso interno, por eso la iniciamos con MINUSCULA
// Cadena de conexion a la base de datos
var clientOptions = options.Client().ApplyURI("mongodb+srv://usuario:passwod@cluster0.w6vsr.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

// ----------------------------------------------------------
// implementamos la funcion
// Esta funcion devolvera un objeto del tipo mongo.Client
func ConectarBD() *mongo.Client {
  // dos parametros de respuesta: client y err
  // al realizar la operacion dependiendo de que pase uno devolvera algo y el otro nil (null)
  // mongo.Connect(param1, param2) hace una conexion a la base de datos
  // param2: es la conexion que contiene la URL de la base de datos
  // param1: usamos el contexto basico, context.TODO()
  // nota: usamos el simbolo := porque estamos asignando y creando a la vez las variables
  client, err := mongo.Connect(context.TODO(), clientOptions)

  // Recordar que en go el null se escribe como nil
  if err != nil {
    // Si err es diferente de nulo significa que si hubo un error
    // usamos .Error() para que convierta el err a un String y pueda agregarse correctamente al log
    log.Fatal(err.Error())
    // retornamos el client y termina esta funcion, aunque client sea nul
    return client
  }

  // Si el err = nil entonces no hubo erro y seguimos con la funcion

  // esta instruccion Ping() la usamos para ver si la Base de datos esta arriba
  // si hubo un error con el ping devolvera un error de lo contrario un nil
  err = client.Ping(context.TODO(), nil)

  // Entonces debemos de volver ha hacer la misma pregunta de antes con el if
  if err != nil {
    log.Fatal(err.Error())
    return client
  }

  // en caso err haya sido nul en todos los casos, significa que el client esta correcto y la conexion a la base de datos esta funcionando bien
  log.Println("Conexion exitosa con la BD")
  return client
}

// creamos una segunda funcion
// aqui devolveremos 
func ChequeoConnection() int {
  // chequeamos el Ping de la base de datos, devolvera nil si no hay error
	err := MongoCN.Ping(context.TODO(), nil)

  // si hubo error err sera diferente de nil asi que retornamos 0
	if err != nil {
		return 0
	}

  // si no hubo error y este es nil, nos devolvera 1
	return 1
}

```

## Programando el archivo handlers

- Creamos el archivo ***handlers.go*** en la carpeta ***handlers***

```go
package handlers

import (
  // para grabar nombre y texto dentro del log de ejecucion
	"log"
	"net/http"
  // Sitema operativo
	"os"

  // importamos algunos de los paquetes que agregamos antes con go get
  // mux para amnejar la respuesta y enviado de informacion
	"github.com/gorilla/mux"
  // cors para tener permisos a la API desde cualquier lugar
  "github.com/rs/cors"
  // Para usar archivos de otras carpetas debemos seguir este ruta
  // github.com/nombreUsuarioGithub/nombreRepositorioGithub/nombreCarpetaDeDondeUsaremosArchivos
	"github.com/miguelmalagaortega/twittor/middlew"
	"github.com/miguelmalagaortega/twittor/routers"
	
)

// Manejadores seteo mi puerto, el handler y pongo a escuchar al servidor
// Desde aqui manejaremos todas las rutas de nuestra aplicacion
func Manejadores() {
  // Iniciamos con la variable router, creando un nuevo Router, veremos con esto si hay informacion en la cabecera(header) o en el cuerpo (body), tambien vera lo del token.
	router := mux.NewRouter()

  // aqui colocamos todas las rutas que vamos a manejar

  // Obtenemos el puerto, con esto vemos si el sistema operativo ya tiene un puerto por defecto
	PORT := os.Getenv("PORT")

  // si la varibale PORT esta vacia, nosotros le asignamos un puerto en este caso el 8080
	if PORT == "" {
		PORT = "8080"
	}

  // cuando sea subido a HEROKU y necesitemos entrar a la API desde otro lugar necesitaremos de cors
  // AllowAll() le da permiso a cualquier por ahora
  //  Handler(router) 
	handler := cors.AllowAll().Handler(router)

  // creamos un Fatal esto pondra a la consola a escuchar todos los cambios en la BD
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
```

## Programando el archivo main.go

1. Iniciaremos la construccion de esta API con el archivos ***main.go***

- Se inicia importando el paquete ***fmt*** para poder mostrar textos por consola.

```go
import (
  "log"

  // Para usar archivos de otras carpetas debemos seguir este ruta
  // github.com/nombreUsuarioGithub/nombreRepositorioGithub/nombreCarpetaDeDondeUsaremosArchivos
  "github.com/miguelmalagaortega/twittor/bd"
	"github.com/miguelmalagaortega/twittor/handlers"
)

func main() {

  // ahora entramos a la bd y usamos la funcion ChequeoConnection()
	if bd.ChequeoConnection() == 0 {
    // si es 0 significaba que hubo un error con la BD
		log.Fatal("Sin conexion a la BD")
		return
	}

  // Con los handles veremos las rutas
	handlers.Manejadores()

}
```

## Compilando por primera vez

- Hasta ahora ya podemos realizar la compilacion de nuestra aplicacion, para eso hacemos en consola `go build main.go`
- Ahora para correr el programa usamos `go run main.go`
- Si todo salio bien debe salir *Conexion Exitosa con la BD*

## Creacion del endPoint - REGISTRO

### Programando el primer modelo - usuario

- Creamos el archivo ***usuario.go*** en la carpeta ***models***

```go
package models

import (
  // lo usamos para la fecha
	"time"

  // llamamos este paquete para la estructura de usuario
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Definimos los atributos que tendra el usuario
type Usuario struct {
  // Nombre       tipo                bson:"Como se guardara en MONGODB, omitempty" json:"nombre que devolvera en el Json"`
  // con el omitempty si lo colocamos en el bson, significa que lo mandamos
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre          string             `bson:"nombre" json:"nombre,omitempty"`
	Apellidos       string             `bson:"apellidos" json:"apellidos,omitempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password,omitempty"`
	Avatar          string             `bson:"avatar" json:"avatar,omitempty"`
	Banner          string             `bson:"banner" json:"banner,omitempty"`
	Biografia       string             `bson:"biografia" json:"biografia,omitempty"`
	Ubicacion       string             `bson:"ubicacion" json:"ubicacion,omitempty"`
	SitioWeb        string             `bson:"sitioweb" json:"sitioWeb,omitempty"`
}
```

### Programando el middleware

- Primero agregamos en el archivo ***handlers.go*** y agregamos la ruta para registrar

```go
package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
  "github.com/rs/cors"
	"github.com/miguelmalagaortega/twittor/middlew"
	"github.com/miguelmalagaortega/twittor/routers"
)

func Manejadores() {
	router := mux.NewRouter()

  // Agregamos esta linea que es para registrar a lso usuarios
  // HandleFunc("ruta", deLaCARPETAmiddlew.funcionChequeBD(De ser correcta la conexion le devuelve el control al routers.Registro)).Methods("tipoDeMetodo")
  // routers es la carpeta u Registro el archivo en la carpeta routers
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

```

- Segundo creamos el archivo ***chequeoBD.go*** en la carpeta ***middlew***

```go
package middlew

import (
	"net/http"

  // para relacionarnos con la carpeta bd
	"github.com/miguelmalagaortega/twittor/bd"
)

// los middlewares reciben algo y devuelven lo mismo
// recibo la conexion http y debo devolver la misma conexion ya que esta conexion viene conlos parametros y cabeceras
// como http.HandlerFunc es una funcion, entonces retornamos un funcion
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
  // func(w response, r *request) * significa que es un puntero
	return func(w http.ResponseWriter, r *http.Request) {
    // chequeamos si se conecta a la base de datos
    // en caso devuelva cero no hay conexion
		if bd.ChequeoConnection() == 0 {
      // mandamos en el http un reponse como Error, ademas ponemos al final el status po ejemplo 500
			http.Error(w, "Conexion perdida con la Base de Datos", 500)
      // con el return terminamos esto
			return
		}

    // En caso de que no haya habido error no vamos al siguiente paso
    // usamos next.ServeHTTP(RESPONSE, REQUEST)
		next.ServeHTTP(w, r)
	}
}
```

### Programando la ruta - registro

- Creamos el archivo ***registro.go*** en la carpeta ***routers***

```go
package routers

import (
  // propio de go para formatos de json
	"encoding/json"
  // para la conexion
	"net/http"

  // traemos los archivos de las carpetas bd y models
	"github.com/miguelmalagaortega/twittor/bd"
	"github.com/miguelmalagaortega/twittor/models"
)

// Metodo que recibe un response, request , es un metodo ya que no devuelve nada
func Registro(w http.ResponseWriter, r *http.Request) {
  // llamamos al modelo 
	var t models.Usuario
  // json.NewDecoder(r.Body)
  // NewDecoder() decodificar lo que viene en el body
  // El body es un objeto de tipo string
  // r.body solo se lee una vez, y luego se destruye en memoria
  // Decode() lo decodificamos en el puntero t, para eso usamos el &
	err := json.NewDecoder(r.Body).Decode(&t)

  // VALIDACIONES

  // En caso err es diferente de nil osea que si hubo un erro
	if err != nil {
    // a la conexion le pasamos el Error, en el response, texto enviado, status)
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
    // terminamos todo con el return
		return
	}

  // si el largo del email es cero no mandaron el email que es indispensable
	if len(t.Email) == 0 {
    // a la conexion le pasamos el Error, en el response, texto enviado, status)
		http.Error(w, "El email de usuario es requerido", 400)
    // terminamos todo con el return
		return
	}

  // si el largo del password es menor a 6
	if len(t.Password) < 6 {
    // a la conexion le pasamos el Error, en el response, texto enviado, status)
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", 400)
    // terminamos todo con el return
		return
	}

  // vamos a hacer una validacion contra los datos
  // tenemos una funcion que devuelve tres valores pero solo me intereza uno, los que no me interezan los ponemos con _
  // crearemos la funcion ChequeoYaExisteUsuario para validar si el correo ya existe
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con el email", 400)
		return
	}

  // llamamos a la funcion InsertoRegistro() que nos devuelve tres valores, pero solo nos importan dos de ellos
  // con esta funcion InsertoRegistro() guardamos los datos en la BD
	_, status, err := bd.InsertoRegistro(t)

  // si err es diferente de nil, hubo un error al guardar los datos
	if err != nil {
		http.Error(w, "Ocurio un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}

  // si status es false, no se pudo relizar el registro, ya que no ocurrio un error pero pudo haber otro problema
	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
		return
	}

  // por el response w devolveresmo por la cabecera WriteHeader el status creado
	w.WriteHeader(http.StatusCreated)

}
```

### Programando la rutina para insertar usuario en la BD

- Creamos el archivo ***insertoregistro.go*** en la carpeta ***bd***

```go
package bd

import (
  // Trabajaremos con contextos
	"context"
  // con el tiempo
	"time"

  // importamos bson
  "go.mongodb.org/mongo-driver/bson/primitive"
  // traemos los archivos de la carpeta models
	"github.com/miguelmalagaortega/twittor/models"
	
)

// funcion InsertoRegistro(nombre tipoDato) devulve tres cosas (string, bool, error)
func InsertoRegistro(u models.Usuario) (string, bool, error) {

  // nos aseguramos que la base de datos no se quede colgada creando un contexto
  // context.WithTimeout(contexto que traigO de la base de datos TODO(), tiempo en segundos)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
  // instruccion que se setea al inicio, pero se ejecuta al final
  // este cancel() cansela el context al final de la rutina por el defer
	defer cancel()

  // la conexion a la base de datos
	db := MongoCN.Database("twittor")
  // que coleccion es la que usaremos, tabla de la bd
	col := db.Collection("usuarios")

  // La password se debe encriptar para eso usamos la funcion EncriptarPassword() a la cual le pasaremos la password que se mando en el modelo
	u.Password, _ = EncriptarPassword(u.Password)

  // A la coleccion le asemos el insert para eso le debemos pasar el contexto y el modelo
	result, err := col.InsertOne(ctx, u)

  // si el err es diferente de nil, significa que hubo un error
	if err != nil {
    // retornamos las tres cosas que devuelve esta funcion
		return "", false, err
	}

  // obtenemos el id, para eso usamos el InsertedID.() y obtenemos el ObjID y el otro valor no lo necesitamos
	ObjID, _ := result.InsertedID.(primitive.ObjectID)

  // retornamos las tres cosas que devuelve esta funcion
  // el id con conversion en string, true, nil
	return ObjID.String(), true, nil
}

```

### Programando el archivo encriptarPassword

- Creamos el archivo ***encriptarPassword.go*** en la carpeta ***bd***

```go
package bd

// solo traemos el paquete necesario para encriptar
import "golang.org/x/crypto/bcrypt"

// creamos la funcion para encriptar el cual recibe un pass como string y devulve dos cosas el password encriptado y el error
func EncriptarPassword(pass string) (string, error) {
  // el costo es el grado de encriptacion se evalua como 2^8 y con eso sera la cantidad de pasadas de la password, mintras mas grande la encriptacion demorara mas
  // para un admin se recomienda un 8
  // para un usuario normal se recomienda un 6
	costo := 8

  // usamos una funcion de bcrypt llamada GenerateFromPassword([])
  // [] este es un slice, que es un vector sin numero de elementos
  // el tipo de ese slice es byte => []byte
  // le pasamos el parametro que en este caso es el pass
  // y el segundo parametro de la funcion GenerateFromPassword es el costo
  // nos devulve tambien en bytes
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)

  // lo convertimos a string y el err si hubiera
	return string(bytes), err

}

```

### Programando el archivo chequeoYaExisteUsuario

- Creamos el archivo ***chequeoYaExisteUsuario.go*** en la carpeta ***bd***

```go
package bd

import (
	"context"
	"time"

	"github.com/miguelmalagaortega/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

// con esta funcion conprobaremos si el email ya existe, le pasamos el email y devolvemos un Usuario, un bool, y un string
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
  // hacemos un contexto de tiempo como hicimos en casos anteriores
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
  // igual cancelamos cuando termina la busqueda en la base de datos
	defer cancel()

  // conectamos con la base de datos
	db := MongoCN.Database("twittor")
  // traemos la coleccion con la cual trabajaremos
	col := db.Collection("usuarios")

  // buscaremos el email en la base de datos para eso usamos el bson
  // bson.M{} para poder usar el formato json
	condicion := bson.M{"email": email}

  // creo una variable de tipo Usuario para grabar los resultados
	var resultado models.Usuario

  // Ahora si buscamos un solo registro en la bd con el FindOne, en la coleccion elegida
  // el FindOne(contexto, condicion) y el Decode para convertirlo en Json con un puntero
  // si hay error se graba en err sino trae resultados graba en la variable resultado
	err := col.FindOne(ctx, condicion).Decode(&resultado)

  // del modelo de usuario extraemos el ID en formato string hexadecimal
	ID := resultado.ID.Hex()

  // si el error es diferente de nil entonces si existio un error
	if err != nil {
    // devolvemos lo siguiente
		return resultado, false, ID
	}

  // si el error es nil devolvemos lo siguiente
	return resultado, true, ID

}
```

## Compilando nuevamente

- Realizamos la compilacion de nuestra aplicacion, para eso hacemos en consola `go build main.go`
- Ahora para correr el programa usamos `go run main.go`
- Si todo salio bien debe salir *Conexion Exitosa con la BD*

## Subiendo a github y HEROKU

- primero subimos a github

> git add .
> git commit -m "Registro de usuario"
> git push origin main

- segundo lo subimos a Heroku

> git push heroku main

## Preparando POSTMAN para las pruebas

1. Creamos una nueva coleccion
1.1. Presionamos sobre el simbolo + y nos aparecera la venta para crear la colecion
![imagen5](/img/5.png)
1.2. En este caso le pondremos el nombre ***twittor***
![imagen6](/img/6.png)
2. Crearemos los ***Environmet***
2.2. Seguimos la imagen para seleccionar ***No Environmet***
![imagen1](/img/1.png)
2.3. Ahora seleccionamos el ojito y luego Add
![imagen2](/img/2.png)
3. Con estos pasos crearemos dos ***Environmet***
3.1. El ***Local***
![imagen4](/img/4.png)
3.2. El ***Heroku***
![imagen3](/img/3.png)
4. Seleccionamos en la coleccion para agregar un ***request***
![imagen7](/img/7.png)
5. Creamos el ***request*** con el siguiente formato
![imagen8](/img/8.png)
6. La parte que dice ***{{Ruta}}*** cambiara al realizar el cambio del ***Environmet*** y apuntara hacia donde se le indico en la Ruta
7. Para probar presionamos ***Send*** y si nos devulve un 201 esta todo correcto
8. Podemos ver en la base de datos que el usuario ya esta registrado
![imagen9](/img/9.png)

## Creacion del endPoint - LOGIN

### Rutina intentoLogin

- Creamos el archivo ***intentoLogin.go*** en la carpeta ***bd***

```go
package bd

import (
	"github.com/miguelmalagaortega/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

// Funcion para hacer login y que devuelve el usuario y el bool
func IntentoLogin(email string, password string) (models.Usuario, bool){
  // usamos la funcion chequeoYaExisteUsuario para ver si exite o no el usuario
  usu, encontrado, _ := chequeoYaExisteUsuario(email)

  if encontrado == false {
    return usu, false
  }

  // guardamos en la variable el password enviado por el usuario
  // esta password esta encriptada
  passwordBytes := []byte(password)

  // guardamos en la variable el passwoord obtenido del usuario al hacer el chequeo si existe
  // esta password no esta encriptada
  passwordBD := []byte(usu.Password)

  // si falla devuelve un error
  err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

  if err != nil {
    return usu, false
  }

  return usu, true

}
```

### Creamos el archivo login

- Creamos el archivo ***login.go*** en la carpeta ***routers***

```go
package routers

import (
  "encoding/json"
  "net/http"
  "time"

  "github.com/miguelmalagaortega/twittor/bd"
  "github.com/miguelmalagaortega/twittor/jwt"
  "github.com/miguelmalagaortega/twittor/models"
)

func Login(w http.ResponseWriter, r *http.Request){
  // Enviamos a la cabecera el formato tipo json
  w.Header().Add("content-type","application/json")

  var t models.Usuario

  // resivimos los dos datos, email y password por medio del Body luego lo decodificamos a json y los guardamos en t
  err := json.NewDecoder(r.Body).Decode(&t)

  if err != nil {
    // devolvemos el error en caso halla
    http.Error(w, "Usuario y/o contraseña invalida " + err.Error(),400)
    return
  }

  if len(t.Email)==0 {
    http.Error(w, "El email del usuario es requerido",400)
    return
  }

  // llamamos a la funcion IntentoLogin para ver si nos devulve un usuario o un error
  documento, existe := bd.IntentoLogin(t.Email, t.Password)

  if existe == false {
    http.Error(w, "Usuario y/o contraseña invalida",400)
    return
  }

  // esto devolvera el token o el error
  jwtKey, err := jwt.GeneroJWT(documento)

  if err != nil {
    http.Error(w, "Ocurrio un error al intentar generar el token correspondiente " + err.Error(),400)
    return
  }

  // Armamos un json con el token para luego devolverlo al navegador
  resp := models.RespuestaLogin {
    Token : jwtKey
  }

  // Con esto devolveremos el token al navegador
  w.Header().Set("content-type","application/json")
  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(resp)

  // Grabar una cookie
  expirationTime := time.Now().Add(24*time.Hour)
  http.SetCookie(w, &http.Cookie{
    Name: "token",
    Value: jwtKey,
    Expires: expirationTime
  })

}
```

### Creamos el archivo jwt

- Creamos el archivo ***jwt.go*** en la carpeta ***jwt***

```go
package jwt

import (
  "time"

  jwt "github.com/dgrijalva/jwt-go"
  "github.com/miguelmalagaortega/twittor/models"
)

// GeneroJWT genera el encriptado con JWT
func GeneroJWT(t models.Usuario) (string, error){

  miClave := []byte("MasterdelDesarrollo_grupodeFacebook")

  // el payload son los datos que estaran en el token
  payload := jwt.MapClaims{
    "email": t.Email,
    "nombre": t.Nombre,
    "apellidos": t.Apellidos,
    "fecha_nacimiento": t.FechaNacimiento,
    "biografia": t.Biografia,
    "ubicacion": t.Ubicacion,
    "sitioweb": t.SitioWeb,
    "_id": t.ID.Hex(),
    "exp": time.Now().Add(time.hour*24).Unix()
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
  tokenStr, err := token.SignedString(miClave)

  if err != nil {
    return tokenStr, err
  }

  return tokenStr, nil

}
```

### Creamos el archivo jwt

- Creamos el archivo ***respuestaLogin.go*** en la carpeta ***models***

```go
package models

type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
```

### Añadir la nueva ruta al archivo handlers

- Abrimos el archivo ***handlers.go*** y agregamos lo siguiente

```go
package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/miguelmalagaortega/twittor/middlew"
	"github.com/miguelmalagaortega/twittor/routers"
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
  // Agregamos esta linea para la nueva ruta
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
```

## Compilando nuevamente

- Realizamos la compilacion de nuestra aplicacion, para eso hacemos en consola `go build main.go`
- Ahora para correr el programa usamos `go run main.go`
- Si todo salio bien debe salir *Conexion Exitosa con la BD*

## Subiendo a github y HEROKU

- primero subimos a github

> git add .
> git commit -m "Registro de usuario"
> git push origin main

- segundo lo subimos a Heroku

> git push heroku main

## Preparando POSTMAN para las pruebas

1. Seleccionamos en la coleccion para agregar un ***request***
![imagen7](/img/7.png)
2. Creamos el ***request*** con el siguiente formato
![imagen10](/img/10.png)
3. La parte que dice ***{{Ruta}}*** cambiara al realizar el cambio del ***Environmet*** y apuntara hacia donde se le indico en la Ruta
4. Para probar presionamos ***Send*** y si nos devulve un 201 esta todo correcto
5. Nso devera devolver un token
![imagen11](/img/11.png)
6. Abrimos el  ***Environmet*** ***Local***
![imagen12](/img/12.png)
7. Agregamos el token
![imagen13](/img/13.png)
8. Si queremos ver lo que sucede con los datos del token nos vamos a la pagina <https://jwt.io/> y ahi escribimos lo siguiente
![imagen14](/img/14.png)

## Middleware validacion de jwt

### Creacion del archivo claim

- Creamos el archivo ***claim.go*** en la carpeta ***models***

```go
package models

import (
  jwt "github.com/dgrijalva/jwt-go"
  "go.mongodb.org/mongo-driver/bson/primitive"
)

type Claim struct {
  Email string `json:"email"`
  ID    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
  jwt.StandardClaims
}
```

### Creacion de rutina de middleware, para validar jwt

- Creamos el archivo ***validoJWT.go*** en la carpeta ***middlew***

```go
package middlew

import (
  "net/http"
  "github.com/miguelmalagaortega/twittor/routers"
)

func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {

  return func(w http.ResponseWriter, r *http.Request){
    // leemos de la cabecera el valor Authorization y lo mandamos a ProcesoToken
    _, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))

    if err != nil {
      http.Error(w, "Error en el token! " + err.Error(), http.StatusBadRequest)
      return
    }

    next.ServeHTTP(w, r)

  }

}
```

### Añadir la nueva ruta al archivo handlers

- Abrimos el archivo ***handlers.go*** y agregamos lo siguiente

```go
package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/miguelmalagaortega/twittor/middlew"
	"github.com/miguelmalagaortega/twittor/routers"
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
  // Agregamos esta linea para la nueva ruta
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.Login))).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
```

### Añadir la rutina de Routers

- Creamos el archivo ***procesoToken.go*** en la carpeta ***routers***

```go
package routers

import (
  "errors"
  "strings"

  jwt "github.com/dgrijalva/jwt-go"
  "github.com/miguelmalagaortega/twittor/bd"
  "github.com/miguelmalagaortega/twittor/models"
)

// Email valor de Email usado en todos los EndPoints
var Email string

// IDUsuario es el ID devuelto del modelo, que se usara en todos los EndPoints
var IDUsuario string

// se recomienda poner el error como parametro devuelto final
func ProcesoToken(tk string) (*models.Claim, bool, string, error){
  miClave := []byte("MasterdelDesarrollo_grupodeFacebook")

  // creamos la variable relacionandola con el puntero
  claims := &models.Claim{}

  // En versiones anteriores el token venia con la palabra "Bearer" al inicio, ahora ya no lo hace asi que estas lineas ya no son necesarias

  // splitToken := strings.Split(tk,"Bearer")

  // if len(splitToken) != 2 {
    // como no tenemos un error que devolver usamos el paquete errors
    // los errores generados no pueden tener ni mayusculas ni signos
    // return claims, false, string(""), errors.New("formato de token invalido")
  // }

  // de los dos elementos que trae el vector splitToken traemos el segundo que es el token
  // ademas eliminamos los espacios en blanco
  // tk = strings.TrimSpace(splitToken[1])

  tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token)(interface{}, error){
    return miClave, nil
  })

  if err == nil {
    _, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)

    if encontrado == true {
      Email = claims.Email
      IDUsuario = claims.ID.Hex()
    }

    return claims, encontrado, IDUsuario, nil
  }

  if !tkn.Valid {
    return claims, false, string(""), errors.New("token invalido")
  }

  return claims, false, string(""), err

}
```

## EndPoint VerPerfil

### Creacion de buscoPefil

- Creamos el archivo ***buscoPerfil.go*** en la carpeta ***bd***

```go
package bd

import (
  "context"
  "fmt"
  "time"

  "github.com/miguelmalagaortega/twittor/models"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscoPerfil(ID string) (models.Usuario, error){

  ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
  defer cancel()

  db := MongoCN.Database("twittor")
  col := db.Collection("usuarios")

  var perfil models.Usuario

  objID, _ := primitive.ObjectIDFromHex(ID)

  // Condicion de busqueda
  condicion := bson.M{
    "_id" : objID,
  }

  // realizamos la busqueda

  err := col.FindOne(ctx, condicion).Decode(&perfil)

  perfil.Password = ""

  if err != nil {
    fmt.Println("Registro no encontrado " + err.Error())
    return perfil, err
  }

  return perfil, nil

}
```

### Creacion de la rutina verPerfil

- Creamos el archivo ***verPerfil.go*** en la carpeta ***routers***

```go
package routers

import (
  "encoding/json"
  "net/http"
  "github.com/miguelmalagaortega/twittor/bd"
)

func VerPerfil(w http.ResponseWriter, r *http.Request){

  ID := r.URL.Query().Get("id")

  if len(ID) < 1 {
    http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
    return
  }

  perfil, err := bd.BuscoPerfil(ID)

  if err != nil {
    http.Error(w, "Ocurrio un error al intentar buscar el registro " + err.Error(), 400)
    return
  }

  w.Header().Set("context-type","application/json")
  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(perfil)

}
```

## Probando el nuevo EndPoint

- Realizamos la compilacion de nuestra aplicacion, para eso hacemos en consola `go build main.go`
- Ahora para correr el programa usamos `go run main.go`
- Si todo salio bien debe salir *Conexion Exitosa con la BD*

1. Cargamos al Github los cambios

> git add .
> git commit -m "verPerfil"
> git push origin main

2. Hacemos lo mismo con heroku

> git push heroku main

3. creamos eun nuevo request en Postman

![imagen15](/img/15.png)

![imagen16](/img/16.png)

4. Ahora ya podemos hacer el send y ver si nos devuelde los datos

## EndPoint modificarPerfil

### Rutina modificarRegistro

- Creamos el archivo ***modificoRegistro.go*** en la carpeta ***bd***

```go
package bd

import (
	"context"
	"time"

	"github.com/miguelmalagaortega/twittor/models"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
)

func ModificoRegistro(u models.Usuario, ID string) (bool, error){

  ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
  defer cancel()

  db := MongoCN.Database("twittor")
  col := db.Collection("usuarios")

  // make, permite crear slice o mapas
  registro := make(map[string]interface{})

  if len(u.Nombre) > 0 {
    registro["nombre"] = u.Nombre
  }

  if len(u.Apellidos) > 0 {
    registro["apellidos"] = u.Apellidos
  }

  registro["fechaNacimiento"] = u.FechaNacimiento

  if len(u.Avatar) > 0 {
    registro["avatar"] = u.Avatar
  }

  if len(u.Banner) > 0 {
    registro["banner"] = u.Banner
  }

  if len(u.Biografia) > 0 {
    registro["biografia"] = u.Biografia
  }

  if len(u.Ubicacion) > 0 {
    registro["ubicacion"] = u.Ubicacion
  }

  if len(u.SitioWeb) > 0 {
    registro["sitioWeb"] = u.SitioWeb
  }

  // con esto tenemos la sentencia para actualizar
  updString := bson.M{
    "$set" : registro,
  }

  objID, _ := primitive.ObjectIDFromHex(ID)

  // ponemos el filtro para decir a que id le hara la actualizacion
  filtro := bson.M{
    "_id": bson.M{"$eq":objID},
  }

  // realizamos la actualizacion
  _, err := col.UpdateOne(ctx, filtro, updString)

  if err != nil {
    return false, err
  }

  return true, nil

}
```

### Funcion enrouters para modificar

- Creamos el archivo ***modificarPerfil.go*** en la carpeta ***routers***

```go
package routers

import (
	"encoding/json"
	"net/http"

	"github.com/miguelmalagaortega/twittor/models"
	"github.com/miguelmalagaortega/twittor/bd"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

  var t models.Usuario

  err := json.NewDecoder(r.Body).Decode(&t)

  if err != nil {
    http.Error(w, "Datos incorrectos " + err.Error(), 400)
    return
  }

  // En esta funcion pasamos el modelo de usuario y la variable global IDUsuario
  status, err := bd.ModificoRegistro(t, IDUsuario)

  if err != nil {
    http.Error(w, "Ocurrio un error al intentar modificar el registro. Reintente nuevamente " + err.Error(), 400)
    return
  }

  if !status {
    http.Error(w, "No se ha logrado modificar el registro del usuario", 400)
    return
  }

  w.WriteHeader(http.StatusCreated)

}
```

## Probando el ENDPOINT

### Actualizamos el handlers

- Abrimos el archivo ***handlers.go*** de la carpeta ***handlers***

```go
package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/miguelmalagaortega/twittor/middlew"
	"github.com/miguelmalagaortega/twittor/routers"
	"github.com/rs/cors"
)

// Manejadores seteo mi puerto, el handler y pongo a escuchar al servidor
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
  // Agregamos esta linea
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")


	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

```

### Creamos un nuevo request en Postman

1. Creamos el nuevo request

![imagen17](/img/17.png)

2. En el body del request ponemos la siguiente informacion

```js
{
  "nombre":"Miguel Angelo",
  "apellidos":"Malaga Ortega",
  "fechaNacimiento":"1988-05-02T00:00:00z",
  "banner":"",
  "ubicacion":"Ciudad de Lima",
  "biografia":"Estudiante de ingenieria de sistemas y electrocina que se encuentra en proceso de aprender la programacion en go",
  "sitioWeb":"https://www.google.com"
}
```

![imagen18](/img/18.png)

3. Ahora ya podemos hacer un Send y ver lo que ocurre