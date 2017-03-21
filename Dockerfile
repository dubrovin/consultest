FROM scratch

ENV CONSUL_ADDR=consul.service.consul:8500 \
	SERVICE_TAGS=default \
	APP_PORT=:8000 \
	SERVICE_8000_NAME=app1

COPY ./consultest /consultest

CMD ["/consultest"]

EXPOSE 8000
