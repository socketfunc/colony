# colony

WebAssemblyのオーケストレーションシステム

wasm runtimeには[wasmer](https://github.com/wasmerio/wasmer)を使用する

Golang実装でgRPCで通信する

## Feature

#### Namespace

マルチテナントを行えるようにする

#### Authorization

基本はJWTで内部にScopeを入れて確認させる

- invoke api
- wasm file read

#### RateLimit

#### Versioning

#### Timeout

#### Metrics

Prometheusでログが取れる
- Istioで適切なルーティングができるようにする

#### Cashing

wasmファイルキャッシュ

- 実行後一定時間はキャッシュさせる
- インメモリでLRUキャッシュをしてデータを保持しておく

#### Deploy

- blue-greenデプロイ
- カナリアリリース

### Architecture

#### colony-runtime

wasmerを使用したwasm runtimeエンジン

- gRPCでリクエストを受け付け、指定の関数を実行させる
- 一定時間実行で終わらない場合はタイムアウトさせる(設定によって無限もできる)
- 実行速度を早めるためにインメモリでのwasmファイルキャッシュを用意する

gRPCのMetadataを使用して認証を行う

```json
{
  "authorization": "token"
}
```

gRPCのサービス設定

request/responseデータは動的に変わるのでMessagePackでシリアライズを行う

```protobuf
message InvokeRequest {
  string namespace = 1;
  string      name = 2;
  bytes       args = 3; // msgpack
}

message InvokeResponse {
  bytes result = 1; // msgpack
}

service Runtime {
  rpc Invoke(InvokeRequest) returns (InvokeResponse) {}
}
```

cliでdeployをする

```
$ colony-cli deploy config.yaml
```

config.yamlでwasmのファイルを指定する

```yaml
version: v1
metadata:
  namespace: default
  name: colony
spec:
  functions:
    - repository: ${http://repository.com/sample.wasm}
      name: ${function-name}
      resources:
        requests:
          timeout: 10s
        limits:
          requests: 1000
```

#### colony-registry

wasmファイルを管理する
バージョン管理を行う
暗号化して保存する(Storageを利用する)
