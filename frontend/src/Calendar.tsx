import React, { useState, useEffect } from 'react'
import styled from 'styled-components'
import dayjs from 'dayjs'

import { rangeTimes } from './Time'
import { Times } from './Times'
import { Borders } from './Borders'
import { AwakeSchedule } from './AwakeSchedule'
import { DateHeaders } from './DateHeaders'

import { getPeriods, Period } from './api/client'

const getDaysBetweenLatestAndOldest = (oldestDate: dayjs.Dayjs, latestDate: dayjs.Dayjs) => {
  const truncateOldestDate = oldestDate.startOf('date')
  const daysBetweenLatestAndOldest: dayjs.Dayjs[] = [truncateOldestDate]

  let truncateDate = truncateOldestDate
  while (!truncateDate.isSame(latestDate, 'date')) {
    truncateDate = truncateDate.add(1, 'day')

    daysBetweenLatestAndOldest.push(truncateDate)
  }

  return daysBetweenLatestAndOldest
}

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

type Tweet = {
  text: string
  createdAt: dayjs.Dayjs
}

type AwakePeriod = {
  okiTime: Tweet
  neTime: Tweet
}

const splidPeriodAtMidnight = (
  period: Period,
  okiCreated: dayjs.Dayjs,
  netaCreated: dayjs.Dayjs
) => {
  const awakePeriods: AwakePeriod[] = []
  let dividedTime = okiCreated.add(1, 'date').startOf('date')
  while (!netaCreated.isSame(dividedTime, 'date')) {
    awakePeriods.push({
      okiTime: { text: period.okiTime.text, createdAt: okiCreated },
      neTime: { text: period.neTime.text, createdAt: dividedTime },
    })
    okiCreated = dividedTime
    dividedTime = dividedTime.add(1, 'day')
  }
  awakePeriods.push({
    okiTime: { text: period.okiTime.text, createdAt: dividedTime },
    neTime: { text: period.neTime.text, createdAt: netaCreated },
  })
  return awakePeriods
}

const splitPeriodsAtMidnight = (periods: Period[]) => {
  let awakePeriods: AwakePeriod[] = []
  for (const period of periods) {
    const okiCreated = dayjs(period.okiTime.createdAt)
    const netaCreated = dayjs(period.neTime.createdAt)
    if (okiCreated.isSame(netaCreated, 'day')) {
      awakePeriods.push({
        okiTime: { text: period.okiTime.text, createdAt: okiCreated },
        neTime: { text: period.neTime.text, createdAt: netaCreated },
      })
    } else {
      const divided = splidPeriodAtMidnight(period, okiCreated, netaCreated)
      awakePeriods = awakePeriods.concat(divided)
    }
  }

  return awakePeriods
}

export const Calendar = () => {
  const [awakePeriods, setAwakePeriods] = useState(new Array<AwakePeriod>())
  const [dateTexts, setDateTexts] = useState(new Array<string>())
  const [dateLabels, setDateLabels] = useState(new Array<string>())
  const [rowTemplate, setRowTemplate] = useState(new Array<string>())

  useEffect(() => {
    const getPeriodsAsync = async () => {
      const res = await getPeriods()
      const awakePeriods = splitPeriodsAtMidnight(res.periods)
      setAwakePeriods(awakePeriods)

      const dates = getDaysBetweenLatestAndOldest(
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
      <Times></Times>
      {awakePeriods.map((awakePeriod, idx) => {
        const neTimeTruncate = awakePeriod.neTime.createdAt.startOf('hour')
        return (
          <AwakeSchedule
            name={awakePeriod.okiTime.createdAt.format('MMMMDD')}
            start={awakePeriod.okiTime.createdAt.startOf('hour').format('HHmm')}
            end={
              neTimeTruncate.hour() === 0
                ? '24'
                : neTimeTruncate.format('HH') + neTimeTruncate.format('mm')
            }
            key={idx}
          ></AwakeSchedule>
        )
      })}
    </Grid>
  )
}
