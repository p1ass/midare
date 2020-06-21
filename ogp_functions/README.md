# ogp_functions

## ローカル開発時

Cloud Storageにアクセスサービスアカウントを準備しておく。

```bash
$ export GOOGLE_APPLICATION_CREDENTIALS=`pwd`/service-account.json
$ yarn local
```

## デプロイ

```bash
$ export GCP_PROJECT=hoge
$ yarn deploy
```
