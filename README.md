# Mesh-backend

## 起動方法

`.env`ファイルを`.env.example`を参考に作る

```bash
docker compose up

# 初回はマイグレーションを実行する
docker compose exec app go run /app/scripts/migrate/main.go
```

ホットリロード付き（docker compose のバージョンを最新にする必要あり）

```bash
docker compose up --watch
```

## 型生成

```bash
protoc -I. --go_out=. --go-grpc_out=. proto/*.proto
```
