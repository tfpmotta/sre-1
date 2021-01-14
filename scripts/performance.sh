#!/bin/bash

##########################################################################
###  Script para realizar teste de performance utilizando ApacheBench  ###
##########################################################################

DIR="/src/scripts"
DIR_LOG="$DIR/logs"

if [ -e "$DIR_LOG" ];then echo "ok";else mkdir $DIR_LOG;fi

## Arquivo de log ex: performance-2101140130 (Ano com dois digitos, mês, dia, hora e minuto)
LOG=`echo performance-$(date '+%y%m%d%H%M%S')`

echo "\nTeste de Performance /get-random-number \n`date`\n" > $DIR_LOG/$LOG

## Teste de carga utilizando um total de 200 requisições com 20 requisições simultâneas
## (Deixei baixo para o lab pq esgotei a cota do random.org algumas vezes) \0/

ab -n 200 -c 20 http://localhost:8080/get-random-number >> $DIR_LOG/$LOG
