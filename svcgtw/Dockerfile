FROM devopsfaith/krakend:latest

WORKDIR /app

ADD . ./

ENV SVCGTW_PORT :9090
EXPOSE 9090

ENTRYPOINT [ "krakend", "run", "-c", "krakend.json" ]

