import { Request, Response } from 'express'
import puppeteer from 'puppeteer'
import { Storage } from '@google-cloud/storage'
import ReactDOMServer from 'react-dom/server'
import styled, { ServerStyleSheet, StyleSheetManager } from 'styled-components'

import { Calendar } from './components/Calendar'
import { Period } from './entity/Period'

interface Body {
  name: string
  iconUrl: string
  uuid: string
  periods: Period[]
}

const Flex = styled.div`
  display: flex;
  height: 100vh;
  width: 100vw;
  align-items: center;
  justify-content: flex-start;
`

const LeftContainer = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 0 32px;
  width: 250px;
  height: 100vh;
  background-color: rgb(88, 149, 98);
`

const IconImage = styled.img`
  width: 200px;
  border-radius: 200px;
`

const LeftHeadline = styled.p`
  font-size: 24px;
  font-weight: bold;
  text-align: center;
  color: white;
`

const WhiteText = styled.p`
  color: white;
`

const globalStyle = `
<style>
@import url('https://fonts.googleapis.com/css2?family=Noto+Sans+JP:wght@400;700&display=swap');
body {
  margin:0;
  font-family: 'Noto Sans JP', sans-serif;
</style>`

export async function ogpFunctions(req: Request<unknown, unknown, Body>, res: Response) {
  const sheet = new ServerStyleSheet()
  const calendar =
    req.body.periods && req.body.periods.length > 0 ? (
      <StyleSheetManager sheet={sheet.instance}>
        <Flex className="ogp-calendar-flex">
          <LeftContainer>
            <IconImage src={req.body.iconUrl.replace('normal', '400x400')}></IconImage>
            <LeftHeadline>{req.body.name}さんの生活習慣</LeftHeadline>
            <WhiteText>#生活習慣の乱れを可視化するやつ</WhiteText>
          </LeftContainer>
          <Calendar periods={req.body.periods}></Calendar>
        </Flex>
      </StyleSheetManager>
    ) : null
  if (!calendar) {
    console.error('calender is null')
    res.status(500).send({})
    return
  }

  const html = ReactDOMServer.renderToStaticMarkup(calendar)
  const styleTags = sheet.getStyleTags()
  sheet.seal()

  const viewport = {
    width: 1200,
    height: 640
  }

  const browser = await puppeteer.launch({
    args: [
      '--no-sandbox',
      '--disable-setuid-sandbox',
      // puppeteer高速化のためのオプション
      // https://github.com/puppeteer/puppeteer/issues/3120#issuecomment-415553869
      '--disable-gpu',
      '--disable-dev-shm-usage',
      '--no-first-run',
      '--no-sandbox',
      '--no-zygote',
      '--single-process' // <- this one doesn't works in Windows
    ],
    headless: process.env.NODE_ENV === 'production'
  })
  const page = (await browser.pages())[0]
  await page.emulateTimezone('Asia/Tokyo')
  await page.setViewport(viewport)
  await page.setContent(html + styleTags + globalStyle)
  const binary = await page.screenshot({ encoding: 'binary' })
  if (process.env.NODE_ENV === 'production') {
    // デバッグしやすいようにローカルではブラウザを閉じない
    await browser.close()
  }

  const filename = req.body.uuid + '.jpg'

  const storage = new Storage()

  const bucketName = process.env.BUCKET_NAME || 'midare-share'

  if (!bucketName) {
    res.status(500)
    res.send('bucket name not found')
    return
  }

  const bucket = storage.bucket(bucketName)

  const blob = bucket.file(filename)

  try {
    console.log('before save')
    if (binary instanceof Buffer) {
      await blob.save(binary)

      console.log('after save')
      res.status(200).send({})
      return
    } else {
      res.status(500).send({ message: 'binary is not Buffer' })
    }
  } catch (e) {
    res.status(500).send(e)
    console.log(e)
    return
  }
}
