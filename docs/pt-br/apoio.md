
# Doc de Apoio

## Configurações Docker

- Ajustes Dockerfile

		- Inclusão de variavel para github privado [ ENV GOPRIVATE=github.com/tfpmotta ]
		- Execução de git config utilizando **access tokens**
		- apt-get update e install de pacotes para troubleshooting e metricas
		- EXPOSE para as portas 8080 9090
		
 - Build 
	
		$ docker build -t sre-1:v1.2 .
		Sending build context to Docker daemon  332.3kB
		Step 1/10 : FROM golang:1.13
		 ---> d6f3656320fe
		Step 2/10 : RUN apt-get update && apt-get install -y     net-tools     tcpdump     vim     apache2-utils
		 ---> Using cache
		 ---> 142cf48afae4
		 
		...
		
		Step 9/10 : EXPOSE 8080 9090
		 ---> Running in a5beb7482876
		Removing intermediate container a5beb7482876
		 ---> ee7176c6e3e3
		Step 10/10 : ENTRYPOINT ["/bin/cmd"]
		 ---> Running in 1de716f3fcb2
		Removing intermediate container 1de716f3fcb2
		 ---> 672268f50de3
		Successfully built 672268f50de3

 - Start Container

		$ docker container run -p 8080:8080 -p 9090:9090 -d sre-1:v1.2
		f5bfa289737e1e6e1fcd83f33297dc70b63e1465241350981a8ac936a12fe01d
		
		$ docker container ls
		CONTAINER ID   IMAGE        COMMAND      CREATED          STATUS          PORTS                                            NAMES
		512171cdea7e   sre-1:v1.2   "/bin/cmd"   24 seconds ago   Up 22 seconds   0.0.0.0:8080->8080/tcp, 0.0.0.0:9090->9090/tcp   cool_swanson
		
		$ docker container exec -it 512171cdea7e /bin/bash
		root@512171cdea7e:/src# /bin/cmd
			2021/01/14 16:04:13 service server started on: http://0.0.0.0:8080
			2021/01/14 16:04:13 metrics server started on: http://0.0.0.0:9090
			2021/01/14 16:04:13 awaiting signal
			2021/01/14 16:04:13 listen tcp 0.0.0.0:8080: bind: address already in use
			2021/01/14 16:04:13 listen tcp 0.0.0.0:9090: bind: address already in use

- Criado script [performance.sh] para realizar testes de performance

		#!/bin/bash
		##########################################################################
		###  Script para realizar teste de performance utilizando ApacheBench  ###
		##########################################################################

		DIR="/src/performance"
		DIR_LOG="$DIR/logs"

		if [ -e "$DIR_LOG" ];then echo "ok" > /dev/null ;else mkdir $DIR_LOG;fi

		## Arquivo de log ex: performance-210114013002.txt (Ano com dois digitos, mês, dia, hora, minuto, segundo)
		LOG=`echo performance-$(date '+%y%m%d%H%M%S').txt`

		echo "\nTeste de Performance /get-random-number \n`date`\n"
		echo "\nTeste de Performance /get-random-number \n`date`\n" > $DIR_LOG/$LOG

		## Teste de carga utilizando um total de 200 requisições com 20 requisições simultâneas
		## (Deixei baixo para o lab pq esgotei a cota do random.org algumas vezes) \0/

		ab -n 200 -c 20 http://localhost:8080/get-random-number >> $DIR_LOG/$LOG
		echo "\nRelatório salvo em $DIR_LOG/$LOG"


- Relatórios sobre estatísticas e métricas dos testes de performance.
	- Relatório gerado pelo ApacheBench em [performance/logs/]

			Teste de Performance /get-random-number 
			Thu Jan 14 01:05:14 -03 2021

			This is ApacheBench, Version 2.3 <$Revision: 1843412 $>
			Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
			Licensed to The Apache Software Foundation, http://www.apache.org/

			Benchmarking localhost (be patient)

			Server Software:        
			Server Hostname:        localhost
			Server Port:            8080

			Document Path:          /get-random-number
			Document Length:        37 bytes

			Concurrency Level:      20
			Time taken for tests:   0.910 seconds
			Complete requests:      200
			Failed requests:        0
			Total transferred:      29000 bytes
			HTML transferred:       7400 bytes
			Requests per second:    219.83 [#/sec] (mean)
			Time per request:       90.980 [ms] (mean)
			Time per request:       4.549 [ms] (mean, across all concurrent requests)
			Transfer rate:          31.13 [Kbytes/sec] received

			Connection Times (ms)
						  min  mean[+/-sd] median   max
			Connect:        0   12  14.2      6      63
			Processing:    23   56  23.9     49     207
			Waiting:       23   44  20.1     39     196
			Total:         30   68  24.0     66     207

			Percentage of the requests served within a certain time (ms)
			  50%     66
			  66%     82
			  75%     84
			  80%     91
			  90%     95
			  95%     99
			  98%    104
			  99%    107
			 100%    207 (longest request)
