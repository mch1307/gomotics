FROM scratch
COPY dist/gomotics /gomotics
EXPOSE 8081
ENTRYPOINT ["/gomotics"]