import React from 'react'
import Head from 'next/head'
import { AppContext, AppInitialProps } from 'next/app'
import Script from 'next/script'

import { isProd } from '../lib/env'

import '../index.css'
import '@fortawesome/fontawesome-svg-core/styles.css'

function MyApp({ Component, pageProps }: AppContext & AppInitialProps) {
  const description =
    'ツイートを使って生活習慣の乱れを可視化するWebアプリです。カレンダーUIで直感的に起床・就寝時間の変化を見ることが出来ます。'
  return (
    <>
      {isProd() ? (
        <>
          <Script
            src="https://www.googletagmanager.com/gtag/js?id=UA-127036212-9"
            strategy="afterInteractive"
          />
          <Script id="google-analytics" strategy="afterInteractive">
            {`
        window.dataLayer = window.dataLayer || [];
        function gtag() {
            dataLayer.push(arguments);
        }
        gtag("js", new Date());
        gtag("config", "UA-127036212-9");
        `}
          </Script>
        </>
      ) : null}
      <Script
        data-ad-client="ca-pub-4978327687969784"
        async
        strategy="afterInteractive"
        src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"
      />
      <Head>
        <meta charSet="utf-8" />
        <link rel="icon" href={`/favicon.ico`} />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <meta name="theme-color" content="#000000" />
        <meta name="description" content={description} />
        <meta property="og:type" content="website" />
        <meta property="og:description" content={description} />
        <meta property="og:image" content="https://midare.p1ass.com/ogp.jpg" key="ogImage" />
        <meta property="og:url" content="https://midare.p1ass.com/" />
        <meta name="twitter:card" content="summary_large_image" />
        <meta name="twitter:site" content="@p1ass" />
        <meta name="twitter:creator" content="@p1ass" />
        <meta property="og:title" content="生活習慣の乱れを可視化するやつ" />
        <link rel="apple-touch-icon" href={`/icon.png`} />
        <link rel="manifest" href={`/manifest.json`} />
        <title>生活習慣の乱れを可視化するやつ</title>
      </Head>
      <Component {...pageProps} />
    </>
  )
}

export default MyApp
