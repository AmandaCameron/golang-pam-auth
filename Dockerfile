FROM golang

WORKDIR /code

RUN go get -u github.com/skelterjohn/wgo
RUN apt-get update && apt-get install pamtester libpam-dev

ADD pam-d/pam-config /usr/share/pam-configs/go-pam-test
#ADD pam-d/login /etc/pam.d/login

ADD src/ /code/src
ADD .gocfg /code/.gocfg


RUN wgo build -o pam-test.so -buildmode c-shared pam-test
