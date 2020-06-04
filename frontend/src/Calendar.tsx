import React, { useState, useEffect } from 'react'
import styled from 'styled-components'
import htmlToImage from 'html-to-image'
import download from 'downloadjs'

import { rangeTimes } from './Time'
import {
  convertPeriodsToAwakePeriods,
  getDatesBetweenLatestAndOldest,
  AwakePeriod,
} from './AwakePeriods'

import { Times } from './Times'
import { Borders } from './Borders'
import { AwakeSchedules } from './AwakeSchedule'
import { DateHeaders } from './DateHeaders'

import { getPeriods, uploadImage } from './api/client'

const timesPerHalfHour = rangeTimes()
const columnTemplate =
  '[t-header] 5fr ' +
  timesPerHalfHour.map((time) => `[t-${time.format('HHmm')}]`).join(' 0.5fr ') +
  ' 0.5fr '

const handleSave = async (dom: HTMLElement | null) => {
  if (!dom) {
    return
  }
  const dataUrl = await htmlToImage.toJpeg(dom, { quality: 0.95 })
  download(dataUrl, 'calendar.jpeg', 'image/jpeg')
}

const shareWithOGP = async (dom: HTMLElement | null) => {
  if (!dom) {
    return
  }
  const blob = await htmlToImage.toBlob(dom, { quality: 0.95 })
  if (blob) {
    const res = await uploadImage(blob)
    const a = document.createElement('a')
    a.href = `http://twitter.com/intent/tweet?url=${res.shareUrl}&text=あなたも生活習慣の乱れを可視化してみませんか？&hashtags=生活習慣の乱れを可視化するやつ`
    a.click()
  }
}

const sleep = (msec: number) => new Promise((resolve) => setTimeout(resolve, msec))

const Tips = () => {
  return (
    <p>
      <span role="img" aria-label="Tips">
        💡
      </span>
      クリックすることで起床後・就寝前のツイートを見ることができます。
    </p>
  )
}

const Button = styled.button`
  margin: 0.5rem 0;
`

const Grid = styled.div<{ rowTemplate: string[]; generatingImage: boolean }>`
  display: grid;
  background: white;
  box-sizing: border-box;
  padding: 0.5rem;
  margin-bottom: 1rem;
  grid-template-rows: ${({ rowTemplate }) => rowTemplate};
  grid-template-columns: ${columnTemplate};
  border: 1px solid #ccc;

  @media (max-width: 40rem) {
    padding: ${({ generatingImage }) => (generatingImage ? '1rem' : '0rem')};
  }
`

export const Calendar = () => {
  const [awakePeriods, setAwakePeriods] = useState(new Array<AwakePeriod>())
  const [dateTexts, setDateTexts] = useState(new Array<string>())
  const [dateLabels, setDateLabels] = useState(new Array<string>())
  const [rowTemplate, setRowTemplate] = useState(new Array<string>())
  const [gridDom, setGridDom] = useState<HTMLElement | null>(null)

  const [infoMsg, setInfoMsg] = useState('Now Loading...')

  const [generatingImage, setGeneratingImage] = useState(false)
  const [generatingType] = useState<'SAVE' | 'TWITTER'>('SAVE')

  useEffect(() => {
    if (!gridDom) {
      return
    }
    const handleSaveAsync = async () => {
      await sleep(1000)
      switch (generatingType) {
        case 'SAVE':
        await handleSave(gridDom)
          break
        case 'TWITTER':
        await shareWithOGP(gridDom)
        break
      }
      setGeneratingImage(false)
    }
    handleSaveAsync()
  }, [gridDom, generatingType])

  useEffect(() => {
    const getPeriodsAsync = async () => {
      try {
        const res = await getPeriods()
        if (res.periods.length === 0) {
          setInfoMsg('直近のツイートが存在しません')
          return
        }
        const awakePeriods = convertPeriodsToAwakePeriods(res.periods)
        setAwakePeriods(awakePeriods)

        const dates = getDatesBetweenLatestAndOldest(
          awakePeriods[awakePeriods.length - 1].okiTime.createdAt,
          awakePeriods[0].neTime.createdAt
        )

        const dateLabels = dates.map((date) => {
          return date.format('MMMMDD')
        })
        setDateLabels(dateLabels)

        const daysOfTheWeek = ['日', '月', '火', '水', '木', '金', '土']
        const dateTexts = dates.map((date) => {
          return date.format(`MM/DD (${daysOfTheWeek[date.day()]})`)
        })
        setDateTexts(dateTexts)

        const rowTemplate = ['time-header']
          .concat(dateLabels)
          .concat('time-footer')
          .map((dateLabel) => `[${dateLabel}] 0.5fr `)
        setRowTemplate(rowTemplate)
      } catch (e) {
        setInfoMsg('ツイートの取得に失敗しました。時間を空けてもう一度お試しください。')
        return
      }
    }
    getPeriodsAsync()
  }, [])

  return (
    <>
      {rowTemplate.length !== 0 ? (
        <>
          <Tips />

          <Grid rowTemplate={rowTemplate} generatingImage={false}>
            <Borders dateLabels={dateLabels} timesPerHalfHour={timesPerHalfHour} />
            <DateHeaders generatingImage={false} dateTexts={dateTexts} />
            <AwakeSchedules awakePeriods={awakePeriods}></AwakeSchedules>
            <Times generatingImage={false} row="time-header"></Times>
            <Times generatingImage={false} row="time-footer"></Times>
          </Grid>
          {/* 画像生成用用DOM */}
          {generatingImage ? (
            <>
              <p>画像生成中...</p>
              <Grid
                rowTemplate={rowTemplate}
                generatingImage={generatingImage}
                ref={(dom) => setGridDom(dom)}
              >
                <Borders dateLabels={dateLabels} timesPerHalfHour={timesPerHalfHour} />
                <DateHeaders generatingImage={generatingImage} dateTexts={dateTexts} />
                <AwakeSchedules awakePeriods={awakePeriods}></AwakeSchedules>
                <Times generatingImage={generatingImage} row="time-header"></Times>
                <Times generatingImage={generatingImage} row="time-footer"></Times>
              </Grid>
            </>
          ) : null}
          {/* <Button
            onClick={async () => {
              setGeneratingType('TWITTER')
              setGeneratingImage(true)
            }}
          >
            結果をTwitterにシェア
          </Button> */}
          <Button
            onClick={async () => {
              setGeneratingImage(true)
            }}
          >
            画像ファイルとして保存
          </Button>
        </>
      ) : (
        <p>{infoMsg}</p>
      )}
    </>
  )
}
