FROM scratch
RUN mkdir -p /app
WORKDIR /app
ADD "./dist/gomotics" "/app/"
EXPOSE 8081
ENTRYPOINT ["/app/gomotics"]