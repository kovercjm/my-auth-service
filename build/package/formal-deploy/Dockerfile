FROM kovercheng/golang-dev:latest as builder

WORKDIR /workspace

COPY . .

RUN make go-build


FROM kovercheng/golang-prod:latest

WORKDIR /project

COPY --from=builder /workspace/dist/my-auth-service ./my-auth-service

ENTRYPOINT [ "./my-auth-service" ]
CMD [ "serve" ]
