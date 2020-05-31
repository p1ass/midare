import React, { useState, useEffect } from 'react'
import styled from 'styled-components'
import AdSense from 'react-adsense'
import htmlToImage from 'html-to-image'

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
import { ErrorBoundary } from './ErrorBoundary'

import { getPeriods } from './api/client'

const timesPerHalfHour = rangeTimes()
const columnTemplate =
  '[t-header] 5fr ' +
  timesPerHalfHour.map((time) => `[t-${time.format('HHmm')}]`).join(' 0.5fr ') +
  ' 0.5fr '

const handleSave = ({ dom }: { dom: HTMLElement }) => {
  if (!dom) {
    return
  }
  htmlToImage.toJpeg(dom, { quality: 0.95 }).then(function (dataUrl) {
    const link = document.createElement('a')
    link.download = 'calendar.jpg'
    link.href = dataUrl
    link.click()
  })
}

const Grid = styled.div<{ rowTemplate: string[] }>`
  display: grid;
  background: white;
  box-sizing: border-box;
  padding: 0.5rem;
  margin-bottom: 1rem;
  grid-template-rows: ${({ rowTemplate }) => rowTemplate};
  grid-template-columns: ${columnTemplate};
  border: 1px solid #ccc;

  @media (max-width: 40rem) {
    padding: 0rem;
  }
`

export const Calendar = () => {
  const [awakePeriods, setAwakePeriods] = useState(new Array<AwakePeriod>())
  const [dateTexts, setDateTexts] = useState(new Array<string>())
  const [dateLabels, setDateLabels] = useState(new Array<string>())
  const [rowTemplate, setRowTemplate] = useState(new Array<string>())
  const [gridDom, setGridDom] = useState<HTMLElement | null>(null)

  useEffect(() => {
    const getPeriodsAsync = async () => {
      const res = await getPeriods()
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
    }
    getPeriodsAsync()
  }, [])

  return (
    <>
      {rowTemplate.length !== 0 ? (
        <>
          <p>
            <span role="img" aria-label="Tips">
              💡
            </span>
            クリックすることで起床後・就寝前のツイートを見ることができます。
          </p>
          <Grid rowTemplate={rowTemplate} ref={(dom) => setGridDom(dom)}>
            <Borders dateLabels={dateLabels} timesPerHalfHour={timesPerHalfHour} />
            <DateHeaders dateTexts={dateTexts} />
            <AwakeSchedules awakePeriods={awakePeriods}></AwakeSchedules>
            <Times row="time-header"></Times>
            <Times row="time-footer"></Times>
          </Grid>
          <button
            onClick={() => {
              if (gridDom) handleSave({ dom: gridDom })
            }}
          >
            画像として保存
          </button>
          <ErrorBoundary>
            <AdSense.Google client="ca-pub-4978327687969784" slot="6211274963" format="auto" />
          </ErrorBoundary>
        </>
      ) : (
        <p>Now Loading...</p>
      )}
    </>
  )
}
