FROM golang:alpine AS builder

WORKDIR /build

# コピー先のディレクトリを指定しないとcmdないを全て展開してコピーしてしまう
COPY ../cmd ./cmd
COPY ../node ./node
COPY ../go.mod .

WORKDIR ./cmd

RUN go build -o p2pNode

# /をつけることでrootから新たに作成できる（WORKDIRは相対パスで行くので/buildから分けるために/distとした）
WORKDIR /dist

#バイナリファイルをdistにコピー
RUN cp /build/cmd/p2pNode .

COPY ../files/ ./files
COPY ../nodeInfo/ ./nodeInfo

# バイナリファイルとfile,nodeInfoをscratchに移動
FROM scratch

COPY --from=builder /dist/p2pNode /
COPY --from=builder /dist/files /files
COPY --from=builder /dist/nodeInfo ./nodeInfo
ENTRYPOINT ["/p2pNode"]