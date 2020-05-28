import React from 'react'
import styled from 'styled-components'
import { Area } from './Area'

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

export const AwakeSchedule = ({ start, end, name }: AwakeScheduleProps) => {
  return <ScheduleBlock colStart={start} colEnd={end} row={name}></ScheduleBlock>
}
