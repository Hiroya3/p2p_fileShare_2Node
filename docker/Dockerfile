FROM golang:alpine AS builder

WORKDIR /build

# コピー先のディレクトリを指定しないとcmdないを全て展開してコピーしてしまう
COPY ../cmd ./cmd
COPY ../node ./node

WORKDIR cmd

CMD "/bin/sh"

#RUN go build -o p2pNode

#WORKDIR dist

# バイナリファイルをdistにコピー
#RUN cp /build/cmd/p2pNode .

#COPY ../files/ .
#COPY ../nodeInfo/ .
#
## バイナリファイルとfile,nodeInfoをscratchに移動
#FROM scratch
#COPY --from=builder /dist/p2pNode .
#COPY --from=builder /dist/files .
#COPY --from=builder /dist/nodeInfo .
#ENTRYPOINT ["/p2pNode"]