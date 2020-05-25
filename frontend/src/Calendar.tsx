import React from 'react'
import styled from 'styled-components'
import { rangeTimes } from './Time'
import { Area } from './Area'
import { Borders } from './Borders'

// TODO : APIで取得したデータを使って範囲を計算する
const rows = Array.from({ length: 30 }, (_, i) => {
  return i + 1
}).map((i) => `May${`0${i}`.slice(-2)}`)

const timesPerHalfHour = rangeTimes()

const ScheduleBlock = styled(Area)`
  background: #429bf4;
  border-radius: 10px;
  font-weight: bold;
  padding: 1em;
  margin: 0.1em 0.5em;
  color: #eee;
  font-size: 0.5rem;
`

interface ScheduleProps {
  start: string
  end: string
  name: string
}

const Schedule = ({ start, end, name }: ScheduleProps) => {
  return <ScheduleBlock colStart={start} colEnd={end} row={name}></ScheduleBlock>
}

const HeaderCell = styled(Area)`
  text-align: center;
  height: 100%;
  padding: 0 8px;
  border-top: solid 1px #ccc;
  margin-top: -1px;
`

const DateText = styled.p`
  margin: 0.2rem;
  font-size: 0.9rem;
`

const DateHeaders = () => {
  return (
    <>
      {rows.map((row) => {
        return (
          <HeaderCell row={row} colStart={'header'} key={row}>
            <DateText>{row}</DateText>
          </HeaderCell>
        )
      })}
    </>
  )
}

const columnTemplate =
  ['[t-header]']
    .concat(timesPerHalfHour.map((time) => `[t-${time.hour}${time.min}]`))
    .join(' 0.3fr ') + ' 0.3fr '
const rowTemplate = rows.map((row) => `[${row}] 0.5fr `)

const Grid = styled.div`
  display: grid;
  background: white;
  box-sizing: border-box;
  margin: 16px;
  grid-template-rows: ${rowTemplate};
  grid-template-columns: ${columnTemplate};
  border: 1px solid #ccc;
`

export const Calendar = () => {
  return (
    <Grid>
      <Borders dateTexts={rows} timesPerHalfHour={timesPerHalfHour} />
      <DateHeaders />

      <Schedule name="May21" start={'1000'} end={'2600'}></Schedule>
      <Schedule name="May22" start={'1130'} end={'2630'}></Schedule>
      <Schedule name="May23" start={'1300'} end={'2630'}></Schedule>
    </Grid>
  )
}
