# backend

## Getting Started

1. [Twitter Developer](https://developer.twitter.com/en) でアプリを作成して、Consumer Key と Consumer Secret を取得。

2. `.env.example` を参考に `.env` ファイルを作成して適宜環境変数を設定する。

```console
$ cp .env.example .env
$ $EDITOR .env
```

3. ローカルでもクッキーを使えるように `/etc/hosts` を編集

```consoel
$ sudo $EDITOR /etc/hosts

# 次の設定を追加
127.0.0.1 localhost.local
::1 localhost.local
```

4. サーバを起動

```console
$ make serve
```

5. http://localhost.local:8080 にアクセスして起動しているか確認。
