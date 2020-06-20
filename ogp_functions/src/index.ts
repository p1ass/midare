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
    page.setViewport(viewport)

    const filename = req.body.uuid + '.jpg'

    try{
        await page.goto(process.env.OGP_URL || 'http://localhost.local:3000/ogp');
        await page.exposeFunction('getPeriods', ()=> req.body.periods)
        await page.waitFor(800)
        await page.screenshot({path: filename});
        await browser.close();
    }catch(e){
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


    try{
        const uploadRes = await bucket.upload(filename,{gzip:true})
        await uploadRes[0].makePublic()
        console.log(uploadRes[0].metadata)
        return res.status(200)
    }catch(e){
        res.status(500)
        res.send(e)
        console.log(e)
        return
    }
}


