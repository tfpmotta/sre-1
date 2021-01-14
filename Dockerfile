FROM golang:1.13

RUN apt-get update && apt-get install -y \
    net-tools \
    tcpdump \
    vim \
    apache2-utils

WORKDIR /src

COPY . .

ENV GO111MODULE=on
ENV GOPRIVATE=github.com/tfpmotta/sre-1

RUN git config --global url."https://tfpmotta:89791e5669c51575eb2312fea6c7f8f8a9cf5dca@github.com".insteadOf "https://github.com"

RUN cd cmd && go build -o /bin/cmd

EXPOSE 8080 9090

ENTRYPOINT ["/bin/cmd"]
