# API estrategia concursos

Olá sou o Breno Andrade, queria agradecer pela oportunidade de fazer o teste e espero conhecer vocês em breve!

Nesse projeto utilizei o Redis como banco em memória e a api do IBM Watson que serviu para recomendar as tags.

[documentação da api](https://estrategia.docs.apiary.io/#)

###  Iniciando o projeto

Para que o projeto seja iniciado basta rodar o comando `make init-install`, os demais comandos que foram escritos no Makefile, foram utilizados para compor este comando principal, eles são o `make deps` que faz o download das dependencias, o `make redis` que inicia a imagem docker [redis:3.2.5-alpine](https://hub.docker.com/_/redis) e o `make init` que somente inicia o projeto.