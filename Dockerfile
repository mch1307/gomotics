FROM alpine:3.6
COPY gomotics /
EXPOSE 8081
ENTRYPOINT ["/gomotics"]