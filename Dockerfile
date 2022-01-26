FROM golang:latest as builder
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
WORKDIR /go/release
COPY ./ /go/release/
RUN CGO_ENABLED=0 go build ./cmd/action-cppcheck

FROM scratch as prod
COPY --from=builder /go/release/action-cppcheck /
CMD ["/action-cppcheck"]