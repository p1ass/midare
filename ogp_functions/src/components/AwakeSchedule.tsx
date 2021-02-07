import styled from 'styled-components'
import dayjs from 'dayjs'

import { Area } from '../atom/Area'
import { AwakePeriod } from '../entity/AwakePeriod'

const ScheduleBlock = styled(Area)`
  background: rgb(88, 149, 98);
  border-radius: 4px;
  font-weight: bold;
  margin: 0.1rem 0;
  color: #eee;
  font-size: 0.5rem;
  cursor: pointer;
`

interface AwakeScheduleProps {
  awakePeriod: AwakePeriod
}

const AwakeSchedule = ({ awakePeriod }: AwakeScheduleProps) => {
  const okiTime = awakePeriod.okiTime.splitDate
    ? awakePeriod.okiTime.splitDate
    : awakePeriod.okiTime.createdAt
  const neTime = awakePeriod.neTime.splitDate
    ? awakePeriod.neTime.splitDate
    : awakePeriod.neTime.createdAt

  const okiTimeTruncate = truncateDate(okiTime)
  const neTimeTruncate = truncateDate(neTime)
  return (
    <>
      <ScheduleBlock
        colStart={okiTimeTruncate.format('HHmm')}
        colEnd={
          !okiTimeTruncate.isSame(neTimeTruncate, 'date') && neTimeTruncate.hour() === 0
            ? '2400'
            : neTimeTruncate.format('HH') + neTimeTruncate.format('mm')
        }
        row={okiTimeTruncate.format('MMMMDD')}
      ></ScheduleBlock>
    </>
  )
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
        return <AwakeSchedule key={idx} awakePeriod={awakePeriod}></AwakeSchedule>
      })}
    </>
  )
}
