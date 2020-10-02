# CitacoesAPI
API criada para retornar citações famosas. No momento estão disponíveis 1035 citações.

## Go
Agora a CitacoesAPI também esta disponivel em Golang. [Clique Aqui](https://github.com/WillianRod/CitacoesAPI/tree/master/go)

Implementação em Golang por [Thiago May](https://github.com/maythiago).

## Node.Js
Agora a CitacoesAPI também esta disponivel em Node.js. [Clique Aqui](https://github.com/WillianRod/CitacoesAPI/tree/master/node)

Implementação em node por [Sergio Junior](https://github.com/SergioJrDev).

## Flask
Agora a CitacoesAPI também esta disponivel em Flask. [Clique Aqui](https://github.com/WillianRod/CitacoesAPI/tree/master/flask)

Implementação em Flask por [Dennis Urtubia](https://github.com/dennisurtubia).

## Implementação
Para implementar o api ao código, basta adicionar as seguintes linhas de código:
```sh
require_once("getquote.php"); // Chama o arquivo das citações
$quote = getQuote(); // cria uma variavel para o array retornado pela função
```
Como a função retorna uma Array, basta usar os valores abaixo para retornar os valores:
```sh
echo $quote['id']; // Printa o ID da citação (Apenas para fins de depuração)
echo $quote['author']; // Printa o autor da citação
echo $quote['tema']; // Printa o tema da citação
echo $quote['quote']; // Printa a citação
```
## JSON
Caso preferir, a api também retorna uma citação aleatória em formato json. Basta colocar o parametro "type" com o valor "json".
```sh
getquote.php?type=json
```

## XML
Se preferir retornar a citação em xml, a api também esta preparada para isso. Ao invés de colocar o valor "json" no parametro "type" basta colocar o valor "xml".
```sh
getquote.php?type=xml
```

