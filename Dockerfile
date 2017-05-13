FROM alpine
EXPOSE 10000
COPY ./bin/gotcp-linux-amd64 /app/gotcp
CMD /app/gotcp server
