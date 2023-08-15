Desafio com base nas seguintes requisições:

## Problema

Nesse desafio você tem a responsabilidade de criar uma Restful API que provê formas de transferir dinheiro de uma pessoa para outra. Cada usuário no sistema tem um saldo que é atualizado em cada transferência.

## Topicos:

* Crie um enpoint que recebe o dois ids de usuários, e um valor monetário representando a transferência entre eles.
* Crie um endpoint que recebe um id de usuário e retorna o saldo dele.
* Valide se o usuário de origem tem saldo suficiente antes da transferência.
* É preciso pensar na possibilidade de concorrência de transferências onde duas pessoas tranferem dinheiro ao mesmo tempo para uma terceira.
* Se uma transferência falhar, o saldo do usuário de origem deve ser restaurado.
* Não é necessário endpoints para criar usuários, popule o banco de forma com que os dois usuários existam e que transferências possam ser feitas entre eles.