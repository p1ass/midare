import { Request, Response } from 'express'
import puppeteer from 'puppeteer'
import { Storage } from '@google-cloud/storage'
import ReactDOMServer from 'react-dom/server'
import styled, { ServerStyleSheet, StyleSheetManager } from 'styled-components'

import { Calendar } from './components/Calendar'
import { Period } from './entity/Period'

interface Body {
  uuid: string
  periods: Period[]
}

const Flex = styled.div`
  display: flex;
  height: 100vh;
  width: 100vw;
  align-items: center;
  justify-content: center;
`

const globalStyle = `
<style>
body {
  margin:0;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue', sans-serif;
</style>`

export async function ogpFunctions(req: Request<unknown, unknown, Body>, res: Response) {
  const sheet = new ServerStyleSheet()
  const calendar = (
    <StyleSheetManager sheet={sheet.instance}>
      <Flex className="ogp-calendar-flex">
        <Calendar periods={req.body.periods}></Calendar>
      </Flex>
    </StyleSheetManager>
  )
  if (!calendar) {
    console.error('calender is null')
    res.status(500).send({})
    return
  }

  const html = ReactDOMServer.renderToStaticMarkup(calendar)
  const styleTags = sheet.getStyleTags()
  sheet.seal()

  const viewport = {
    width: 1280,
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

  // デバッグしやすいようにローカルではブラウザを閉じない
  if (process.env.NODE_ENV === 'production') {
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
    await blob.save(binary)
    console.log('after save')
    res.status(200).send({})
    return
  } catch (e) {
    res.status(500).send(e)
    console.log(e)
    return
  }
}
