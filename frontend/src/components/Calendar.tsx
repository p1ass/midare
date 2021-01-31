import React, { forwardRef, useState, useEffect } from 'react'
import styled from 'styled-components'

import { rangeTimes } from '../lib/time'
import {
  convertPeriodsToAwakePeriods,
  getDatesBetweenLatestAndOldest,
  AwakePeriod,
} from '../entity/AwakePeriod'
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

interface CalendarProps {
  periods: Period[]
  generatingImage: boolean
}

export const Calendar = forwardRef(function _Calendar(
  { periods, generatingImage }: CalendarProps,
  ref: React.Ref<HTMLDivElement>
) {
  const [awakePeriods, setAwakePeriods] = useState(new Array<AwakePeriod>())
  const [dateTexts, setDateTexts] = useState(new Array<string>())
  const [dateLabels, setDateLabels] = useState(new Array<string>())
  const [rowTemplate, setRowTemplate] = useState(new Array<string>())

  useEffect(() => {
    const awakePeriods = convertPeriodsToAwakePeriods(periods)
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
  }, [periods])

  return (
    <>
      {rowTemplate.length !== 0 ? (
        <>
          <Grid rowTemplate={rowTemplate} generatingImage={generatingImage} ref={ref}>
            <Borders dateLabels={dateLabels} timesPerHalfHour={timesPerHalfHour} />
            <DateHeaders generatingImage={generatingImage} dateTexts={dateTexts} />
            <AwakeSchedules awakePeriods={awakePeriods} />
            <Times generatingImage={generatingImage} row="time-header" />
            <Times generatingImage={generatingImage} row="time-footer" />
          </Grid>
        </>
      ) : null}
    </>
  )
})
