import { createRef, useState, useEffect } from 'react'
import { toJpeg } from 'html-to-image'
import download from 'downloadjs'

import { usePeriods } from '../api/hooks'

import { Calendar } from './Calendar'
import { ButtonSaveImage } from './ButtonSaveImage'
import { ButtonShareTwitter } from './ButtonShareTwitter'

const handleSave = async (dom: HTMLDivElement | null) => {
  if (!dom) {
    return
  }
  const dataUrl = await toJpeg(dom, { quality: 0.95 })
  download(dataUrl, 'calendar.jpeg', 'image/jpeg')
}

const sleep = (msec: number) => new Promise((resolve) => setTimeout(resolve, msec))

export const CalendarContainer = () => {
  const [infoMsg, setInfoMsg] = useState('Now Loading...')

  const [generatingImage, setGeneratingImage] = useState(false)

  const ref = createRef<HTMLDivElement>()

  useEffect(() => {
    const handleSaveAsync = async () => {
      await sleep(1000)
      await handleSave(ref.current)
      setGeneratingImage(false)
    }
    handleSaveAsync()
  }, [ref])

  const [periods, shareUrl, error] = usePeriods()
  useEffect(() => {
    if (periods && periods.length === 0) {
      setInfoMsg('直近のツイートが存在しません')
    }
    if (error) {
      console.error(error)
      setInfoMsg('ツイートの取得に失敗しました。時間を空けてもう一度お試しください。')
    }
  }, [periods, error])

  return (
    <>
      {periods && periods.length !== 0 ? (
        <>
          <Calendar periods={periods} generatingImage={false} />
          <ButtonShareTwitter shareUrl={shareUrl} />
          <ButtonSaveImage
            onClick={async () => {
              setGeneratingImage(true)
            }}
          />
          {generatingImage ? (
            <>
              <p>画像生成中...</p>
              <Calendar periods={periods} generatingImage={generatingImage} ref={ref} />
            </>
          ) : null}
        </>
      ) : (
        <p>{infoMsg}</p>
      )}
    </>
  )
}
