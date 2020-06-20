# cloud-functions-typescript-template
cloud-functions-typescript-templateはTypeScriptを使ったGoogle Cloud Functionsのベースプロジェクトです。このプロジェクトにはよく使いそうなツールや設定ファイル（例: Test toolやLint）が最初から入っています。もしオススメのツールや設定ファイルがあったら気軽にPull RequestやIssueを提出してください。

### Prerequisites
npm, tsc, gcloud コマンドがインストール済

## package.json の configを編集

```
"config": {
  "function_name": "関数の名前を定義,　この関数名はindex.tsの export function の名前と一致する必要があります",
  "region": "europe-west1 or us-east1 or us-central1 or asia-northeast1",
  "gcp_project": "デプロイ先のGCPプロジェクトのIDをいれます",
  "runtime": "nodejs8"
}
```

## Build

```
$npm run build
```

## Test

```
$npm run test
```

## Deploy

```
$npm run deploy --prefix functions/src/
```
