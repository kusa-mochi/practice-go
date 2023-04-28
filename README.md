# practice-go

## モジュール定義

```bash
# bash
cd <main.goのあるディレクトリ>
go mod init practice_go     # "practice_go" は、後述の「パッケージ参照」で使う識別子。
```

## パッケージ定義

```go
// sub1/sub1.go
package sub1
```
```go
// sub1/sub1-1.go
package sub1
```
```go
// sub2/sub2.go
package sub2
```

## パッケージ参照

```go
// main.go
import (
    "practice_go/sub1"  // sub1.go と sub1-1.go で定義されているものが使えるようになる。
    "practice_go/sub2"  // sub2.go で定義されているものが使えるようになる。
)
```

## main.goの実行

```bash
go run main.go
```

## ビルド

```bash
go build -o practice_go
```

## ビルド出力の実行

```bash
# bash
./practice_go
```
