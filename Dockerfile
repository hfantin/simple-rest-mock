FROM alpine:latest
ENV TZ Brazil/East
WORKDIR /opt/app
ADD bin/srm-arm64 srm
ADD .files/* .files/
ADD .env .env
RUN chmod +x srm
CMD [ "/opt/app/srm" ]