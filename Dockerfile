FROM alpine:3.6
COPY dist/gomotics /gomotics
EXPOSE 8081
ENTRYPOINT ["/gomotics"]