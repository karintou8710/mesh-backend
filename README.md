# Mesh-backend

## 起動方法

`.env`ファイルを`.env.example`を参考に作る

```bash
$ docker compose up

# 初回はマイグレーションを実行する
$ docker compose exec app go run /app/scripts/migrate/main.go
```
