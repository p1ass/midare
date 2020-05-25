import React from 'react'
import styled from 'styled-components'
import { flatten } from './lib'
import { Area } from './Area'
import { Time } from './Time'

const Border = styled(Area)<{ min: string }>`
  border-top: solid 1px #ccc;
  min-height: 0.2rem;
  border-left: ${({ min }) => {
    return min === '00' ? `1px solid #ccc` : `none`
  }};
  margin-top: -1px;
`

interface BordersProps {
  dateTexts: string[]
  timesPerHalfHour: Time[]
}

export const Borders = ({ dateTexts, timesPerHalfHour }: BordersProps) => {
  const elms = dateTexts.map((dateText) => {
    return timesPerHalfHour.map((time, i) => (
      <Border
        row={dateText}
        colStart={`${time.hour}${time.min}`}
        min={time.min}
        key={`${dateText}-${i}`}
      />
    ))
  })
  return flatten(elms)
}
