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

import { getLoginUrl, getPeriods } from './api/client'

const timesPerHalfHour = rangeTimes()
const columnTemplate =
  '[t-header] 5fr ' +
  timesPerHalfHour.map((time) => `[t-${time.hour}${time.min}]`).join(' 0.5fr ') +
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

// TODO : 時間の四捨五入をしっかりする
// TODO : 3時間半時間が空いていないところがあるのでチェック

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

      const dateTexts = dates.map((date) => {
        return date.format('MM/DD (dd)')
      })
      setDateTexts(dateTexts)

      const rowTemplate = ['time'].concat(dateLabels).map((dateLabel) => `[${dateLabel}] 0.5fr `)
      setRowTemplate(rowTemplate)
    }
    getPeriodsAsync()
  }, [])

  return (
    <>
      <a href={getLoginUrl()}>ログイン</a>
      <Grid rowTemplate={rowTemplate}>
        <Borders dateLabels={dateLabels} timesPerHalfHour={timesPerHalfHour} />
        <DateHeaders dateTexts={dateTexts} />
        <AwakeSchedules awakePeriods={awakePeriods}></AwakeSchedules>
        <Times></Times>
      </Grid>
    </>
  )
}
