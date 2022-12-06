#### Notas: 
+ se desarrolla con la premisa de no entregar información sobre posibles errores y mas bien manejarlos por logs.
+ se desarrolla test unitarios para punto 1 y 2 
+ No se desarrolla el punt 4
+ dentro del repositorio estara el Modelo entidad relacion, archivo docker para levantar un servidor y la BD, un archivo para inicializar la BD, test unitarios,
arquitectura del proyecto y collección de postman.
+ El punto 3 del crud, solo se desarrolla el obtener usuarios e insertalos, la eliminación se planeaba como una eliminación sinbolica en
la cual solo se desactivara el usuario... por ende se pensaba un PACTH tanto para eliminar como para actualizar la info del usuario...
* iniciar comando... $  docker-compose up -d --build y luego comprobar los correspondientes puertos denotados en el .yml

## DOCUMENTACIÓN

## Endpoint 1: classifiedList

Endpoint: __/classified__ <br>
Método: __POST__

* Definir un api el cual reciba una matriz de números y devuelva esta misma de una manera ordenada, tenga en cuenta que
  los números duplicados se deben mover al final de la lista ordenada.


### Criterios y/o Restricciones:

1. Solo se permiten enteros (positivos y negativos)
2. Una lista vaciá se responderá con sí misma dando a entender que ya esta clasificada
3. La longitud maxima permitida de la lista sera de 100
4. La forma de agregar los elementos repetidos __NO ES ALEATORIA__, su comportamiento sigue el patron de que el número
   repetido (de izquierda a derecha) que ya esté clasificado, debera irse agregando al final de la lista (después de los
   clasificados)
5. Implementar test Unitarios

#### Ejemplo REQUEST:

~~~json
{
  "sin clasificar": [
    3,
    5,
    5,
    6,
    8,
    3,
    4,
    4,
    7,
    7,
    1,
    1,
    2
  ]
}
~~~

#### Ejemplo RESPONSE:

~~~json
{
  "sin-clasificar": [
    3,
    5,
    5,
    6,
    8,
    3,
    4,
    4,
    7,
    7,
    1,
    1,
    2
  ],
  "clasificado": [
    1,
    2,
    3,
    4,
    5,
    6,
    7,
    8,
    5,
    3,
    4,
    7,
    1
  ]
}
~~~

## Endpoint 2: balance

Endpoint: __/balance/{filterMes}__ <br>
Método: __POST__

* Crear un api que reciba un objeto con los meses del año con las ventas y gastos asociados. El formato de petición sera
  el siguiente:

~~~json
{
  "meses": [
    "Enero",
    "Febrero",
    "Marzo",
    "Abril"
  ],
  "ventas": [
    30500,
    35600,
    28300,
    33900
  ],
  "gastos": [
    22000,
    23400,
    18100,
    20700
  ]
}
~~~

* La respuesta del Api debera dar el balance de cada mes ingresado. El formato de cada mes será el siguiente:

~~~json
{
  "mes": "Enero",
  "ventas": 30500,
  "gastos": 35600,
  "balance": -5100
}
~~~

* Teniendo esto en cuenta, La respuesta del Api según el ejemplo será la siguiente

~~~json
{
  "Balances": [
    {
      "Mes": "Enero",
      "Ventas": 30500,
      "Gastos": 35600,
      "Balance": -5100
    },
    {
      "Mes": "Febrero",
      "Ventas": 35600,
      "Gastos": 23400,
      "Balance": 12200
    },
    {
      "Mes": "Marzo",
      "Ventas": 28300,
      "Gastos": 18100,
      "Balance": 10200
    },
    {
      "Mes": "Abril",
      "Ventas": 33900,
      "Gastos": 20700,
      "Balance": 13200
    }
  ]
}
~~~

* __Valor Agregado:__
    * Si se decide filtrar por un mes en específico, será permitido y esto retornará solo el balance de dicho mes
        * Ejemplo:
            * si filterMes = febrero, entonces se hará el balance únicamente del mes febrero
            * ~~~json 
              { 
                  "Mes": "Febrero",
                  "Ventas": 35600,
                  "Gastos": 23400,
                  "Balance":12200
              } 
              ~~~

### Criterios y/o Restricciones:

1. Todas las listas deben tener la misma longitud y la longitud máxima permitida será 100
2. Valores permitidos en meses = `[enero,febrero,marzo,abril,mayo,junio,julio,agosto,septiembre,octubre,noviembre,diciembre]` en cualquier formato 
3. Si en los meses ingresados hay un valor diferente a los meses del año, no se permite hacer el balance
4. Si filterMes es diferente a un mes del año, se hará el balance de todos los meses
5. El response con los Balances estará siempre ordenado por los meses
6. El valor en las Ventas y Gastos será un entero positivo o cero
7. Se puede incluir N veces un mes en la petición para hacer su balance 
8. Si se filtra por un mes que no se incluya en la petición, no se responderá ningún contenido
9. implementar test unitarios

## Endpoint 3: users

Endpoint: __/users__ <br>
Método: __GET__

* NOTA: Correr los scripts de DB ubicados en data_base_init.sql antes de ejecutar los endpoints relacionados con los usuarios

* crear un esquema de base de datos con un mínimo de tres tablas relacionadas, en el cual una de ellas será el modelo de usuario al cual se le debe crear un servicio que permita añadir, consultar, modificar y eliminar su información.

#### Descripción
* Se crea un endpoint para consultar la información del usuario permitiendo filtrar por distintos valores simulando listas de selección y campo de búsqueda
* Un usuario cuenta con su información básica almacenada en la tabla users, además se crea una tabla countries que indica a que pais pertenece el usuario, de los países existentes. Al igual que con el tipo de documento (tabla document_type) y Area profesional (tabla department)
* Este endpoint responde con todos los usuarios que coincidieron con los filtros o búsquedas, los filtros y busquedas pueden mezclarse para obtener resultados más precisos.

### Criterios y/o Restricciones:
* Filtro por varios países, retorna la información de los usuarios de los países solicitados.
* Filtro por tipo de documento, retorna la información de los usuarios de los tipos solicitados.
* Filtro por áreas profesionales, retorna la información de los usuarios de las áreas seleccionadas.
* Filtro por estado del usuario, retorna la información de los usuarios del estado solicitado. Estados permitidos = "enable","disable","stand-by"
  + Esto ayudara a filtrar a los usuarios que estén en estado activo, inactivos o en suspeción en nuestra DB
* Campo de búsqueda (search). Esta búsqueda se realiza a nivel de:
    + Primer nombre
    + Segundo nombre
    + Primer apellido
    + Segundo apellido
    + Correo electrónico
    + Número de documento
* * Todo con la finalidad de permitir buscar correos, números de documentos o nombres en específico, se retorna la información de los usuarios que coincida con la búsqueda en algunos de los mencionados campos
* Se debe permitir combinar filtros entre sí más el campo de búsqueda.
* Paginación por el método de búsqueda

### USO DE QUERY PARAMS PARA LOS FILTROS Y EL CAMPO DE BÚSQUEDA
#### Campos en QueryParams

~~~
* limit (indica el limite de registros que tendra la página,requerido)
cursor (indica el ultimo registro de la petición anterios para saber que pagina es la siguiente)
search (campo de busqueda)
* status (indica el estado de los usuarios a consultar, requerido)
countries (lista de id de paises por los cuales filtrar)
identifications_types (lista de id de tipos de documentos por los cuales filtrar)
departments (lista de dareas profesionales por las cuales filtrar)
~~~

#### Ejemplo queryparams:  
?limit=2&search=gian&status=enable
+ indica querer ver 2 resultados (si los hay) de un usuario en estado enable y que en sus nombres, correo o número de documento exista el valor "gian" sin importar la ubicación de la coincidencia.

#### Ejemplo de Response
~~~json
{
  "last_cursor": "KCdHaWFuJywnOTBiMzcwOTYtMjUxOC00YWRlLWJkYzMtZjI1MzkzZTY1MTI0Jyk=",
  "total_registers": 12,
  "users": [
    {
      "id": "03fdc57f-e9bf-47d0-a25a-255f4a7f894b",
      "name": "Gian",
      "others_names": "Lucca",
      "last_name": "Aguado",
      "second_last_name": "Rendon",
      "country": "Estados Unidos",
      "identification_type": "Cedula de Ciudadania",
      "identification_number": "1116238356",
      "email": "gian.aguado@correo.com.us",
      "department": "Infraestructura",
      "status": "enable"
    },
    {
      "id": "90b37096-2518-4ade-bdc3-f25393e65124",
      "name": "Alex",
      "last_name": "Lopez",
      "country": "Estados Unidos",
      "identification_type": "Cedula de Ciudadania",
      "identification_number": "1116238356",
      "email": "gian.lopez@correo.com.us",
      "department": "Infraestructura",
      "status": "enable"
    }
  ]
}
~~~

## Endpoint 4: users

Endpoint: __/users__ <br>
Método: __POST__

### creación de usuarios
### Criterios y/o Restricciones:
* el numero de identificación no se repite en una mismo tipo de documento
* el email del usuario es generado automaticamente y garantiza unicidad
* el momento de la creación del usuario quedara asociado en su registro
* el estado del usuario creado sera habilitado
+ validaciones de campos requeridos (name,last_name,country_id,identification_type_id,identification_number,department_id)
+ no se admiten caracteres especiales

### ejemplo request 
~~~json
{
  "name": "Ejemplo",
  "others_names": "Lucca",
  "last_name": "Apellido",
  "second_last_name": "",
  "country_id": 1,
  "identification_type_id": 1,
  "identification_number": "11223344",
  "department_id": 1
}
~~~


### FIN :3 