import React, { SFC } from 'react'
import styled from 'styled-components'
type Time = {
  hour: string
  min: string
}
const members = Array.from({ length: 30 }, (_, i) => {
  return i + 1
}).map((i) => `May${i}`)

const rows = members

// 時間を30分単位で出力
const rangeTimes = (start = 6, hours = 24): Time[] => {
  return Array.from({ length: hours * 2 + 1 }, (_, i) => {
    const hr = Math.floor(i / 2) + start
    const min = (i % 2) * 30
    const minSprint = min === 0 ? '00' : min
    const hrSprint = hr < 10 ? `0${hr}` : hr
    return { hour: hrSprint.toString(), min: minSprint.toString() }
  })
}

const times = rangeTimes()
const columnTemplate = ['[t-header]']
  .concat(times.map((time) => `[t-${time.hour}${time.min}]`))
  .join(' 0.3fr ')

const rowTemplate = rows.map((row) => `[${row}] 1fr `)

const Grid = styled.div`
  display: grid;
  background: white;
  box-sizing: border-box;
  margin: 16px;
  grid-template-rows: ${rowTemplate};
  grid-template-columns: ${columnTemplate};
`

interface AreaProps {
  row: string
  colStart: string
  colEnd?: string
}

const Area = styled.div.attrs<AreaProps>(({ row, colStart, colEnd }) => ({
  style: {
    gridRow: row,
    gridColumn: colEnd ? `t-${colStart} / t-${colEnd}` : `t-${colStart}`,
  },
}))<AreaProps>``

// @ts-ignore
const flatten = (item) => item.reduce((a, b) => a.concat(b), [])
// 全部のエリアにborderを撒き散らす

const Border = styled(Area)`
  border-top: solid 1px #ccc;
  min-height: 0.2rem;
  border-left: 1px solid #ccc;
`

const Borders = () => {
  const elms = rows.map((row) => {
    return times.map((time, i) => (
      <Border row={row} colStart={`${time.hour}${time.min}`} key={`${row}-${i}`} />
    ))
  })
  return flatten(elms)
}

const Time = styled(Area)`
  font-size: 0.2rem;
`

const ScheduleBlock = styled(Area)`
  background: #429bf4;
  border-radius: 10px;
  font-weight: bold;
  padding: 1em;
  margin: 0.1em 0.5em;
  color: #eee;
  font-size: 0.5rem;
`

const Schedule: SFC<{ start: string; end: string; name: string }> = ({ start, end, name }) => {
  return <ScheduleBlock colStart={start} colEnd={end} row={name}></ScheduleBlock>
}

const HeaderCell = styled(Area)`
  text-align: center;
  height: 100%;
  padding: 0 8px;
`
const Headers = () => {
  return (
    <>
      {members.map((member) => {
        return (
          <HeaderCell row={member} colStart={'header'} key={member}>
            <p>{member}</p>
          </HeaderCell>
        )
      })}
    </>
  )
}

export const Timetable = () => {
  return (
    <Grid>
      <Borders />
      <Headers />

      <Schedule name="May21" start={'1000'} end={'2600'}></Schedule>
      <Schedule name="May22" start={'1130'} end={'2630'}></Schedule>
      <Schedule name="May23" start={'1300'} end={'2630'}></Schedule>
    </Grid>
  )
}
