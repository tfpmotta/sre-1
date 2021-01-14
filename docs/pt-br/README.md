# Olá!

Você está então agora entrando na jornada de uma pessoa SRE, vamos apenas realizar alguns testes para entender como você avalia problemas e desenvolve soluções criativas e eficientes para problemas modernos.

Esse repositório (e quando citarmos *"esse repositório"*, estamos nos referindo ao diretório raíz dele) é um projeto. Um projeto de automação de teste de uma aplicação. Porém temos alguns desafios para que ele alcance o objetivo de **trazer confiabilidade para o negócio garantindo qualidade de desenvolvimento e boas práticas de aplicações cloud native**. Claro, também podemos citar que um diferencial nessas metas seria **reduzir o número, tempo e responsabilidade de interações humanas para que o pipeline tenha garantias de sua execução**.

Algumas regras:

- Você está livre para pesquisar e utilizar as ferramentas que está acostumado para resolver o problema desde que mantenha as tecnologias obrigatórias: **GitHub Actions, Docker, Linux**.

Como é feita a avaliação:

- Os critérios serão avaliados por peso e demonstração de competência, exemplo: *sobre o assunto X (com peso Z) o projeto atendeu Y%, logo sua pontuação será: Y%Z*.
- Critérios podem ter o mesmo peso, alguns são mais importantes do que outros, queremos avaliar e entender suas decisões.
- A soma de da pontuação de todos os critérios (demonstrações de capacidades) será sua nota final.
- Partindo da pontuação teórica máxima realizaremos uma avaliação interna de qual perfil e nível podemos te oferecer.

**Uma dica importante:** se não conseguir realizar as várias tarefas dessa avaliação, não se preocupe, foque em fazer as que você se sente confortável.

## Contexto

Para dar um pouco de contexto ao projeto vamos falar sobre o cenário atual: faltam apenas 30 dias para a data de uma campanha promocional que será feita pelo time de marketing em conjunto com o time comercial e pelo histórico sabemos que o volume de usuários aumenta muito.

Essa promoção envolve um serviço de sorteio que depende de um serviço para gerar números randômicos. Como o negócio cresceu, no último ano fizemos a transição de um data center privado para a nuvem e agora não temos mais o serviço antigo que era baseado em geração por hardware. Precisamos de uma nova solução.

## O que fizemos até agora

Temos uma aplicação em **Golang** nesse repositório e uma declaração de **workflow do GitHub Action**. Essa aplicação expõe os seguintes endpoints:

- `HTTP, porta 8080, GET /random-number` - um número randômico em um body em JSON.
- `HTTP, porta 9090, GET /metrics` - métricas sobre a aplicação.

## Próximos passos

Crie um repositório privado na sua conta no GitHub e adicione os usuários que informamos. Nesse repositório faça o push para a branch **main** do conteúdo desse `.zip` que te enviamos. As tarefas abaixo devem ser feitas nesse repositório que você criou.

- Comandos utilizados
	$ git remote add origin https://github.com/tfpmotta/sre-1.git
	$ git branch -M main
	$ git push -u origin main

## Para configurar seu repositório

- [X] Realize a substituição de todas as strings `testing/sre-test-1` por `SEU_USUARIO_GIT/NOME_DO_SEU_REPOSITÓRIO` criando um script para fazer essa tarefa (na linguagem de sua escolha).
	Comando utilizado - [ $ find ./* -type f  ! -iname '*readme*' -exec grep -l 'testing/sre-test-1' {} \; -exec sed 's/testing\/sre-test-1/tfpmotta\/sre-1/' {} \; ]

- [X] Faça o commit e push da alteração para seu repositório.
	[ $ git add .
	  $ git commit -m "replace strings"
	  $ git push ]

## To fix

- [X] Aplicação não está realizando build da imagem Docker.
	Ajustes Dockerfile
		- Inclusão de variavel para github privado [ ENV GOPRIVATE=github.com/tfpmotta ]
		- Execução de git config utilizando **access tokens**
		- apt-get update e install de pacotes para troubleshooting e metricas
		- EXPOSE para as portas 8080 9090
		
	Build ok - [ $ docker build -t sre-1 . ]
	Start Container ok - [ $ docker container run -p 8080:8080 -p 9090:9090 -d sre-1:latest ]

- [ ] Não temos logs no pipeline ou alertas indicando sucesso do teste funcional.

- [ ] Existe um step no pipeline em que realizamos um teste funcional realizando o request para http://localhost:8080/random-number e validamos a resposta, verificar se o teste feito aqui realmente garante que o endpoint está respondendo devidamente.
- [ ] Criar o mesmo teste funcional para a rota `/metrics` da porta **9090**.

## To do

- [X] Realizar testes de performance na geração de números randômicos.
	- Criado script [performance.sh] teste de performance utilizando ApacheBench

- [X] Trazer relatórios sobre estatísticas e métricas dos testes de performance.
	- O script [performance.sh] está salvado relatório gerado pelo ApacheBench em [performance/logs/]
	
- [ ] Diminuir tempo de geração de número randômico.
- [ ] Criar documentação para outros colaboradores contribuírem com o projeto.
- [ ] Implementar métricas sobre o serviço http que responde na rota `/get-random-number` (dicas https://www.robustperception.io/prometheus-middleware-for-gorilla-mux e para uma implementação mais simples, utilize o arquivo [internal/router/router.go](../../internal/router/router.go)) expondo através da rota `/metrics` as métricas adicionais.
- [ ] Reduzir tempo de execução do workflow (GitHub Action).
