import React, { useState, useEffect } from 'react'
import styled from 'styled-components'

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

import { getPeriods } from './api/client'

const timesPerHalfHour = rangeTimes()
const columnTemplate =
  '[t-header] 5fr ' +
  timesPerHalfHour.map((time) => `[t-${time.format('HHmm')}]`).join(' 0.5fr ') +
  ' 0.5fr '

const Grid = styled.div<{ rowTemplate: string[] }>`
  display: grid;
  background: white;
  box-sizing: border-box;
  margin: 16px;
  grid-template-rows: ${({ rowTemplate }) => rowTemplate};
  grid-template-columns: ${columnTemplate};
  border: 1px solid #ccc;
`

export const Calendar = () => {
  const [awakePeriods, setAwakePeriods] = useState(new Array<AwakePeriod>())
  const [dateTexts, setDateTexts] = useState(new Array<string>())
  const [dateLabels, setDateLabels] = useState(new Array<string>())
  const [rowTemplate, setRowTemplate] = useState(new Array<string>())

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
          <p>クリックすることで起床後・就寝前のツイートを見ることができます。</p>
          <Grid rowTemplate={rowTemplate}>
            <Borders dateLabels={dateLabels} timesPerHalfHour={timesPerHalfHour} />
            <DateHeaders dateTexts={dateTexts} />
            <AwakeSchedules awakePeriods={awakePeriods}></AwakeSchedules>
            <Times row="time-header"></Times>
            <Times row="time-footer"></Times>
          </Grid>
        </>
      ) : (
        <p>Now Loading...</p>
      )}
    </>
  )
}
