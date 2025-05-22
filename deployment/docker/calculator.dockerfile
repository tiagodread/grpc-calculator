FROM gcr.io/distroless/static

USER nobody

WORKDIR /app
COPY bin/calculator /app/calculator

EXPOSE 50051

CMD ["/app/calculator"]