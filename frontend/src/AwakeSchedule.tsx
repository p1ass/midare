import React from 'react'
import styled from 'styled-components'
import { Area } from './Area'
import { AwakePeriod } from './AwakePeriods'

const ScheduleBlock = styled(Area)`
  background: rgb(88, 149, 98);
  border-radius: 4px;
  font-weight: bold;
  padding: 1em;
  margin: 0.1em 0.5em;
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

export const AwakeSchedules = ({ awakePeriods }: AwakeSchedulesProps) => {
  return (
    <>
      {awakePeriods.map((awakePeriod, idx) => {
        const neTimeTruncate = awakePeriod.neTime.createdAt.startOf('hour')
        const okiTimeTrunate = awakePeriod.okiTime.createdAt.startOf('hour')
        return (
          <AwakeSchedule
            name={awakePeriod.okiTime.createdAt.format('MMMMDD')}
            start={okiTimeTrunate.format('HHmm')}
            end={
              neTimeTruncate.hour() === 0 && okiTimeTrunate.hour() !== 0
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
