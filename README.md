# p1ass/midare

🕒 ツイートを使って生活習慣の乱れを可視化するWebアプリ  
https://midare.p1ass.com

![Build](https://github.com/p1ass/midare/workflows/Build/badge.svg)

## スクリーンショット

**あまり乱れてない人**
![tae](images/tae.png)

**乱れている人**
![tae](images/yabai.png)

## ニュース記事

- [ツイート時刻から「生活習慣の乱れ」を可視化してくれるアプリが登場　「ツイ廃判定アプリ」「なんて恐ろしいものを」 - ねとらぼ](https://nlab.itmedia.co.jp/nl/articles/2006/03/news042.html)
- [ツイ廃ぶりが一目瞭然！ 過去のツイートを分析して「生活の乱れを可視化」するアプリ【やじうまWatch】 - INTERNET Watch](https://internet.watch.impress.co.jp/docs/yajiuma/1257035.html) 

## 登壇資料

 [うじまる君の生活習慣の乱れを可視化したい！ / uzimaru birthday LT - Speaker Deck](https://speakerdeck.com/p1ass/uzimaru-birthday-lt )

## ブログ

TBD

## Getting Started

1. [Twitter Developer](https://developer.twitter.com/en) でアプリを作成して、Consumer KeyとConsumer Secretを取得。

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
