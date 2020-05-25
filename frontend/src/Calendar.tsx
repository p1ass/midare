import React from 'react'
import styled from 'styled-components'

import { rangeTimes } from './Time'
import { Borders } from './Borders'
import { AwakeSchedule } from './AwakeSchedule'
import { DateHeaders } from './DateHeaders'

// TODO : APIで取得したデータを使って範囲を計算する
const dateTexts = Array.from({ length: 30 }, (_, i) => {
  return i + 1
}).map((i) => `May${`0${i}`.slice(-2)}`)

const timesPerHalfHour = rangeTimes()

const columnTemplate =
  ['[t-header]']
    .concat(timesPerHalfHour.map((time) => `[t-${time.hour}${time.min}]`))
    .join(' 0.3fr ') + ' 0.3fr '
const rowTemplate = dateTexts.map((dateText) => `[${dateText}] 0.5fr `)

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
      <Borders dateTexts={dateTexts} timesPerHalfHour={timesPerHalfHour} />
      <DateHeaders dateTexts={dateTexts} />

      <AwakeSchedule name="May21" start={'1000'} end={'2600'}></AwakeSchedule>
      <AwakeSchedule name="May22" start={'1130'} end={'2630'}></AwakeSchedule>
      <AwakeSchedule name="May23" start={'1300'} end={'2630'}></AwakeSchedule>
    </Grid>
  )
}
