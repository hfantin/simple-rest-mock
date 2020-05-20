FROM alpine:latest
ENV TZ Brazil/East
WORKDIR /opt/app
ADD bin/simple-rest-mock-arm64 simple-rest-mock
ADD .files/* .files/
RUN chmod +x simple-rest-mock
CMD [ "/opt/app/simple-rest-mock" ]