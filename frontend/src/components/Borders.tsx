import React from 'react'
import styled from 'styled-components'
import dayjs from 'dayjs'
import { Area } from '../atom/Area'

const Border = styled(Area)<{ time: dayjs.Dayjs }>`
  border-top: solid 1px #ccc;
  border-bottom: solid 1px #ccc;
  min-height: 0.2rem;
  border-left: ${({ time }) => {
    return time.minute() === 0 ? `1px solid #ccc` : `none`
  }};
  margin-top: -1px;
`

interface BordersProps {
  dateLabels: string[]
  timesPerHalfHour: dayjs.Dayjs[]
}

export const Borders = ({ dateLabels, timesPerHalfHour }: BordersProps) => {
  return (
    <>
      {dateLabels.map((dateText) => {
        return timesPerHalfHour.map((time, i) => (
          <Border
            row={dateText}
            colStart={time.format('HHmm')}
            time={time}
            key={`${dateText}-${i}`}
          />
        ))
      })}
    </>
  )
}
