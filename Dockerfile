FROM ubuntu:18.04
COPY bin/service .
RUN ./service 
