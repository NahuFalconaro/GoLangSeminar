# GoLangSeminario

Hicimos la estructura donde retornamos si la cadena es valida, si no retornamos un error.
Arrancamos analizando al estructura de la cadena, como conclucion sacamos que:
-Los dos primeros digitos son el tipo, siendo NN para los numericos, y TX para los de alfabetico
-Los siguientes dos digitos, son el indicador de que tan larga es la cadena de valores
-Los digitos a partir de ahi, son los valores.

Preguntamos si la longitud de la cadena es mayor o igual a 4, para evaluar si inicialmente es valida.

Chequeamos el tipo de la cadena, para asi preguntar, si tiene todos los valores correspondientes.
Es decir, si el tipo es NN, los valores sean todos numericos.
Y luego chequeamos que el valor indicado en la longitud de la cadena, es igual a la cantidad de digitos que tiene el valor


Para el testing, utilizamos assert para comprobar que los resultados son los esperados.