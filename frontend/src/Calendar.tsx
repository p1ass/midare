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
        awakePeriods[awakePeriods.length - 1].neTime.createdAt,
        awakePeriods[0].okiTime.createdAt
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
    <Grid rowTemplate={rowTemplate}>
      <Borders dateLabels={dateLabels} timesPerHalfHour={timesPerHalfHour} />
      <DateHeaders dateTexts={dateTexts} />
      <AwakeSchedules awakePeriods={awakePeriods}></AwakeSchedules>
      <Times></Times>
    </Grid>
  )
}
