FROM ubuntu:20.04

EXPOSE 8080

ADD bin/HelloGoService HelloGoService/bin/HelloGoService

CMD ["HelloGoService/bin/HelloGoService"]