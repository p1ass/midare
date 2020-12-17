import { Request, Response } from 'express'
import puppeteer from 'puppeteer';
import {Storage} from '@google-cloud/storage'

interface Body{
    uuid: string
    periods: any[]
}

export async function ogpFunctions(req: Request<any,any,Body>, res: Response) {
    const viewport = {
        width: 1280,
        height: 640,
    }

    const browser = await puppeteer.launch();
    const page = await browser.newPage();
    await page.emulateTimezone('Asia/Tokyo')
    page.setViewport(viewport)

    const filename = req.body.uuid + '.jpg'

    let binary : Buffer

    try{
        await page.goto(process.env.OGP_URL || 'http://localhost.local:3000/ogp');
        await page.exposeFunction('getPeriods', ()=> req.body.periods)
        await page.waitForTimeout(1200)
        binary = await page.screenshot({encoding: 'binary'});
        await browser.close();
    }catch(e){
        console.log(e)
        res.status(500).send(e)
        return
    }

    const storage = new Storage()

    const bucketName = process.env.BUCKET_NAME || 'midare-share'

    if (!bucketName){
        res.status(500)
        res.send('bucket name not found')
        return
    }

    const bucket = storage.bucket(bucketName)

    const blob = bucket.file(filename);

    try{
        await blob.save(binary)
        res.status(200).send({})
        return
    }catch(e){
        res.status(500).send(e)
        console.log(e)
        return
    }
}


