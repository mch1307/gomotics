FROM alpine:3.6
LABEL maintainer "mch1307@gmail.com"

# Set environment
ENV SERVICE_HOME=/gomotics

RUN mkdir -p ${SERVICE_HOME}/etc ${SERVICE_HOME}/log 
ADD "./dist/gomotics" "/bin/gomotics"
EXPOSE 8081
WORKDIR $SERVICE_HOME/log
VOLUME /gomotics/log
ENTRYPOINT ["/bin/gomotics"]
