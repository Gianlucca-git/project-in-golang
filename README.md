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
