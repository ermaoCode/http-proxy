FROM hub.ark.jd.com/jvessel/jvessel-base:1.0
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
# jvessel
ENV JVES_HEALTH_CHECK_SCRIPT '/helloworld/healthcheck.sh'
ENV JVES_LOG_DIR '/helloworld/logs'

# user
RUN mkdir -p /helloworld/logs
COPY healthcheck.sh /helloworld/healthcheck.sh
COPY entrypoint.sh  /helloworld/entrypoint.sh

COPY http-proxy /usr/local/bin/http-proxy

ENTRYPOINT [ "/helloworld/entrypoint.sh" ]
