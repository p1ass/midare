import React from 'react'
import styled from 'styled-components'
import dayjs from 'dayjs'
import { Area } from './Area'
import { AwakePeriod } from './AwakePeriods'

const ScheduleBlock = styled(Area)`
  background: rgb(88, 149, 98);
  border-radius: 4px;
  font-weight: bold;
  margin: 0.1rem 0;
  color: #eee;
  font-size: 0.5rem;
`

interface AwakeScheduleProps {
  start: string
  end: string
  name: string
}

const AwakeSchedule = ({ start, end, name }: AwakeScheduleProps) => {
  return <ScheduleBlock colStart={start} colEnd={end} row={name}></ScheduleBlock>
}

interface AwakeSchedulesProps {
  awakePeriods: AwakePeriod[]
}

const truncateDate = (date: dayjs.Dayjs) => {
  if (date.minute() < 15) {
    return date.startOf('hour')
  }
  if (date.minute() >= 15 && date.minute() < 45) {
    return date.startOf('hour').add(30, 'minute')
  }
  return date.add(1, 'hour').startOf('hour')
}

export const AwakeSchedules = ({ awakePeriods }: AwakeSchedulesProps) => {
  return (
    <>
      {awakePeriods.map((awakePeriod, idx) => {
        const okiTimeTrunate = truncateDate(awakePeriod.okiTime.createdAt)
        const neTimeTruncate = truncateDate(awakePeriod.neTime.createdAt)
        return (
          <AwakeSchedule
            name={okiTimeTrunate.format('MMMMDD')}
            start={okiTimeTrunate.format('HHmm')}
            end={
              okiTimeTrunate.hour() !== 0 && neTimeTruncate.hour() === 0
                ? '2400'
                : neTimeTruncate.format('HH') + neTimeTruncate.format('mm')
            }
            key={idx}
          ></AwakeSchedule>
        )
      })}
    </>
  )
}
