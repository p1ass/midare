import Head from 'next/head'
import { AppContext, AppInitialProps } from 'next/app'
import { IsProd } from '../common/env'
import '../index.css'

function MyApp({ Component, pageProps }: AppContext & AppInitialProps) {
  const description =
    'ツイートを使って生活習慣の乱れを可視化するWebアプリです。カレンダーUIで直感的に起床・就寝時間の変化を見ることが出来ます。'
  return (
    <>
      <Head>
        <script async src="https://www.googletagmanager.com/gtag/js?id=UA-127036212-9"></script>
        {IsProd() ? (
          <script
            dangerouslySetInnerHTML={{
              __html: `
        window.dataLayer = window.dataLayer || [];
        function gtag() {
            dataLayer.push(arguments);
        }
        gtag("js", new Date());
        gtag("config", "UA-127036212-9");
        `,
            }}
          ></script>
        ) : null}
        <meta charSet="utf-8" />
        <link rel="icon" href={`${process.env.PUBLIC_URL}/favicon.ico`} />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <meta name="theme-color" content="#000000" />
        <meta name="description" content={description} />
        <meta property="og:type" content="website" />
        <meta property="og:description" content={description} />
        <meta property="og:image" content="https://midare.p1ass.com/ogp.jpg" />
        <meta property="og:url" content="https://midare.p1ass.com/" />
        <meta name="twitter:card" content="summary_large_image" />
        <meta name="twitter:site" content="@p1ass" />
        <meta name="twitter:creator" content="@p1ass" />
        <meta property="og:title" content="生活習慣の乱れを可視化するやつ" />
        <link rel="apple-touch-icon" href={`${process.env.PUBLIC_URL}/icon.png`} />
        <link rel="manifest" href={`${process.env.PUBLIC_URL}/manifest.json`} />
        <title>生活習慣の乱れを可視化するやつ</title>
        <script
          data-ad-client="ca-pub-4978327687969784"
          async
          src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"
        ></script>
      </Head>
      <Component {...pageProps} />
    </>
  )
}

export default MyApp
