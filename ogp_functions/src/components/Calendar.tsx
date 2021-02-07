import React from 'react'
import styled from 'styled-components'

import { rangeTimes } from '../lib/time'
import { convertPeriodsToAwakePeriods, getDatesBetweenLatestAndOldest } from '../entity/AwakePeriod'
import { Period } from '../entity/Period'

import { Times } from './Times'
import { Borders } from './Borders'
import { AwakeSchedules } from './AwakeSchedule'
import { DateHeaders } from './DateHeaders'

const timesPerHalfHour = rangeTimes()
const columnTemplate =
  '[t-header] 5fr ' +
  timesPerHalfHour.map((time) => `[t-${time.format('HHmm')}]`).join(' 0.5fr ') +
  ' 0.5fr '

const Grid = styled.div<{ rowTemplate: string[] }>`
  display: grid;
  background: white;
  box-sizing: border-box;
  padding: 0.5rem;
  grid-template-rows: ${({ rowTemplate }) => rowTemplate};
  grid-template-columns: ${columnTemplate};
  border: 1px solid #ccc;
`

interface CalendarProps {
  periods: Period[]
}

export const Calendar = ({ periods }: CalendarProps) => {
  const awakePeriods = convertPeriodsToAwakePeriods(periods)

  const dates = getDatesBetweenLatestAndOldest(
    awakePeriods[awakePeriods.length - 1].okiTime.createdAt,
    awakePeriods[0].neTime.createdAt
  )

  const dateLabels = dates.map((date) => {
    return date.format('MMMMDD')
  })

  const daysOfTheWeek = ['日', '月', '火', '水', '木', '金', '土']
  const dateTexts = dates.map((date) => {
    return date.format(`MM/DD (${daysOfTheWeek[date.day()]})`)
  })

  const rowTemplate = ['time-header']
    .concat(dateLabels)
    .concat('time-footer')
    .map((dateLabel) => `[${dateLabel}] 0.5fr `)

  return (
    <>
      {rowTemplate.length !== 0 ? (
        <>
          <Grid rowTemplate={rowTemplate}>
            <Borders dateLabels={dateLabels} timesPerHalfHour={timesPerHalfHour} />
            <DateHeaders dateTexts={dateTexts} />
            <AwakeSchedules awakePeriods={awakePeriods} />
            <Times row="time-header" />
            <Times row="time-footer" />
          </Grid>
        </>
      ) : null}
    </>
  )
}
