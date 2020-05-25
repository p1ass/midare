import React, { SFC } from 'react'
import styled from 'styled-components'
type Time = {
  hour: string
  min: string
}
const members = ['taro', 'jiro', 'hanako']

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
const rowTemplate = ['[t-header]']
  .concat(times.map((time) => `[t-${time.hour}${time.min}]`))
  .join(' 0.3fr ')

const Grid = styled.div`
  display: grid;
  background: white;
  box-sizing: border-box;
  grid-template-rows: [time] 0.5fr [taro] 1fr [jiro] 1fr [hanako] 1fr;
  grid-template-columns: ${rowTemplate};
`

const Area = styled.div<{ column: string; rowStart: string; rowEnd?: string }>`
  grid-row: ${({ column }) => column};
  grid-column: ${({ rowStart, rowEnd }) => {
    if (rowEnd) {
      return `t-${rowStart} / t-${rowEnd}`
    }
    return `t-${rowStart}`
  }};
`

// @ts-ignore
const flatten = (item) => item.reduce((a, b) => a.concat(b), [])
// 全部のエリアにborderを撒き散らす

const Border = styled(Area)`
  border-top: solid 1px #ccc;
  min-height: 0.2rem;
  border-left: 1px solid #ccc;
`

const Borders = () => {
  const elms = rows.map((column) => {
    return times.map((time, i) => (
      <Border column={column} rowStart={`${time.hour}${time.min}`} key={`${column}-${i}`} />
    ))
  })
  return flatten(elms)
}

const Time = styled(Area)`
  font-size: 0.2rem;
`

const ScheduleBlock = styled(Area)`
  background: #429bf4;
  /* border: 1px solid #2b293f; */
  border-radius: 10px;
  font-weight: bold;
  padding: 1em;
  margin: 0.1em 0.5em;
  color: #eee;
  font-size: 0.5rem;
`

const Schedule: SFC<{ start: string; end: string; name: string }> = ({ start, end, name }) => {
  return <ScheduleBlock rowStart={start} rowEnd={end} column={name}></ScheduleBlock>
}

const HeaderCell = styled(Area)`
  text-align: center;
  height: 100%;
`
const Headers = () => {
  return (
    <>
      {members.map((member) => {
        return (
          <HeaderCell column={member} rowStart={'header'}>
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

      <Schedule name="taro" start={'0830'} end={'0900'}></Schedule>
      <Schedule name="taro" start={'1300'} end={'2700'}></Schedule>
      <Schedule name="hanako" start={'1130'} end={'1500'}></Schedule>
      <Schedule name="jiro" start={'1230'} end={'1600'}></Schedule>
      <Schedule name="hanako" start={'1030'} end={'1400'}></Schedule>
    </Grid>
  )
}
