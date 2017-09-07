FROM alpine:3.6
LABEL maintainer "mch1307@gmail.com"

# Set environment
ENV SERVICE_NAME=gomotics \
    SERVICE_HOME=/gomotics \
    SERVICE_VERSION=v1.2.0 

# Download and install gomotics
RUN mkdir -p ${SERVICE_HOME}/etc ${SERVICE_HOME}/log 
ADD "./dist/gomotics" "/bin/gomotics"
EXPOSE 8081
WORKDIR $SERVICE_HOME/log
ENTRYPOINT ["/bin/gomotics"]
