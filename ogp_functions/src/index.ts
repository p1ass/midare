import { Request, Response } from 'express'
import puppeteer from 'puppeteer';
 

export async function ogpFunctions(_: Request, res: Response) {
    const browser = await puppeteer.launch();
    const page = await browser.newPage();
    await page.goto('https://example.com');
    await page.screenshot({path: 'example.png'});

    await browser.close();
    try {
        res.status(200)
        res.send('Hello World')
    } catch (err) {
        res.status(500)
        res.send(err)
    }
}
