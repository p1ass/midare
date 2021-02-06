import { Request, Response } from 'express'
import puppeteer from 'puppeteer'
import { Storage } from '@google-cloud/storage'

interface Body {
  uuid: string
  periods: any[]
}

const MAX_RETRY_COUNT = 3

export async function ogpFunctions(req: Request<any, any, Body>, res: Response) {
  const viewport = {
    width: 1280,
    height: 640
  }

  const browser = await puppeteer.launch({
    args: ['--no-sandbox', '--disable-setuid-sandbox'],
    headless: process.env.NODE_ENV === 'production'
  })
  const page = (await browser.pages())[0]
  await page.emulateTimezone('Asia/Tokyo')
  await page.setViewport(viewport)

  const filename = req.body.uuid + '.jpg'

  let binary = Buffer.from('')

  for (let retryCnt = 0; retryCnt < MAX_RETRY_COUNT; retryCnt++) {
    try {
      await page.goto(process.env.OGP_URL || 'http://localhost:3000/ogp')
      await page.exposeFunction('getPeriods', () => req.body.periods)
      await page.waitForSelector('.ogp-calendar-flex', { timeout: 5000 })
      binary = await page.screenshot({ encoding: 'binary' })
      break
    } catch (e) {
      if (retryCnt < MAX_RETRY_COUNT) {
        continue
      }
      console.error(e)
      res.status(500).send(e)
      return
    }
  }

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
  } catch (e) {
    res.status(500).send(e)
    console.log(e)

    return
  }

  try {
    console.log('before update ogp')
    await updateOGP(page, req.body.uuid)
    console.log('after update ogp')
    await browser.close()
    res.status(200).send({})
    return
  } catch (e) {
    res.status(500).send(e)
    console.log(e)
    return
  }
}

const updateOGP = async (page: puppeteer.Page, uuid: string) => {
  await page.goto('https://cards-dev.twitter.com/validator')
  await page.waitForSelector(
    '#react-root > div > div > div.css-1dbjc4n.r-13qz1uu.r-417010 > main > div > div > div.css-1dbjc4n.r-13qz1uu > form > div > div:nth-child(6) > label > div > div.css-1dbjc4n.r-18u37iz.r-16y2uox.r-1wbh5a2.r-19h5ruw.r-1udh08x.r-1inuy60.r-ou255f.r-m611by > div > input'
  )
  await page.type(
    '#react-root > div > div > div.css-1dbjc4n.r-13qz1uu.r-417010 > main > div > div > div.css-1dbjc4n.r-13qz1uu > form > div > div:nth-child(6) > label > div > div.css-1dbjc4n.r-18u37iz.r-16y2uox.r-1wbh5a2.r-19h5ruw.r-1udh08x.r-1inuy60.r-ou255f.r-m611by > div > input',
    process.env.TWITTER_USER || ''
  )
  await page.type(
    '#react-root > div > div > div.css-1dbjc4n.r-13qz1uu.r-417010 > main > div > div > div.css-1dbjc4n.r-13qz1uu > form > div > div:nth-child(7) > label > div > div.css-1dbjc4n.r-18u37iz.r-16y2uox.r-1wbh5a2.r-19h5ruw.r-1udh08x.r-1inuy60.r-ou255f.r-m611by > div > input',
    process.env.TWITTER_PASSWORD || ''
  )
  await page.click(
    '#react-root > div > div > div.css-1dbjc4n.r-13qz1uu.r-417010 > main > div > div > div.css-1dbjc4n.r-13qz1uu > form > div > div:nth-child(8) > div > div'
  )
  await page.waitForSelector('#ValidationForm > div > div:nth-child(1) > input.FormControl')
  await page.type(
    '#ValidationForm > div > div:nth-child(1) > input.FormControl',
    `https://midare.p1ass.com/share/${uuid}`
  )
  await page.click('#ValidationForm > div > div.Grid-cell.u-sizeFill.u-marginTm > input')
}
