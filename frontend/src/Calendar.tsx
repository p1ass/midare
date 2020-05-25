import React, { SFC } from 'react'
import styled from 'styled-components'
type Time = {
  hour: string
  min: string
}
const members = ['taro', 'jiro', 'hanako']

const rows = ['time'].concat(members)

const rangeTimes = (start = 6, hours = 24): Time[] => {
  return Array.from({ length: hours * 2 + 1 }, (_, i) => {
    const hr = Math.floor(i / 2) + start
    const min = (i % 2) * 30
    const minSprint = min === 0 ? '00' : min

    // return `${hr}:${minSprint}`;
    return { hour: hr.toString(), min: minSprint.toString() }
  })
}

const times = rangeTimes() // 時間を15分単位で排出
const rowTemplate = ['[t-header]']
  .concat(times.map((time) => `[t-${time.hour}${time.min}]`))
  .join(' 0.2fr ')
// rowTemplateは↓こんな感じの文字列に
// [t-1000] 1fr [t-1015] 1fr [t-1030] 1fr [t-1045] 1fr [t-1100] 1fr [t-1115] 1fr [t-1130] ..

const Grid = styled.div`
  display: grid;
  background: white;
  box-sizing: border-box;
  grid-template-columns: [time] 0.5fr [taro] 1fr [jiro] 1fr [hanako] 1fr;
  grid-template-rows: ${rowTemplate};
`

const Area = styled.div<{ column: string; rowStart: string; rowEnd?: string }>`
  grid-column: ${({ column }) => column};
  grid-row: ${({ rowStart, rowEnd }) => {
    if (rowEnd) {
      return `t-${rowStart} / t-${rowEnd}`
    }
    return `t-${rowStart}`
  }};
`

// @ts-ignore
const flatten = (item) => item.reduce((a, b) => a.concat(b), [])
// 全部のエリアにborderを撒き散らす

const Border = styled(Area)<{ color: string }>`
  border-left: solid 1px #ccc;
  min-height: 0.2rem;
  border-top: ${({ color = '#ccc' }) => `solid 1px ${color}`};
`

const Borders = () => {
  const elms = rows.map((column) => {
    return times.map((time, i) => (
      <Border
        color={time.min == '00' ? '#333' : '#ccc'}
        column={column}
        rowStart={`${time.hour}${time.min}`}
        key={`${column}-${i}`}
      />
    ))
  })
  return flatten(elms)
}

// @ts-ignore
const Times: SFC<{}> = () => {
  return rangeTimes().map((time, i) => {
    return (
      <Area rowStart={`${time.hour}${time.min}`} column={'time'} key={i.toString()}>
        {time.hour}:{time.min}
      </Area>
    )
  })
}

const ScheduleBlock = styled(Area)`
  background: #429bf4;
  /* border: 1px solid #2b293f; */
  border-radius: 10px;
  font-weight: bold;
  padding: 1em;
  margin: 0.1em 0.5em;
  color: #eee;
`

const Schedule: SFC<{ start: string; end: string; name: string }> = ({
  start,
  end,
  name,
  children,
}) => {
  const time = `${start.substr(0, 2)}:${start.substr(2, 4)}`
  return (
    <ScheduleBlock rowStart={start} rowEnd={end} column={name}>
      <div>{time}</div>
      <div>
        [{name}
        ]: {children}
      </div>
    </ScheduleBlock>
  )
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
            {member}
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
      <Times />

      <Schedule name="taro" start={'1030'} end={'1145'}>
        外出
      </Schedule>
      <Schedule name="taro" start={'1345'} end={'1600'}>
        外出
      </Schedule>
      <Schedule name="hanako" start={'1130'} end={'1200'}>
        お昼
      </Schedule>
      <Schedule name="jiro" start={'1230'} end={'1400'}>
        会議
      </Schedule>
      <Schedule name="hanako" start={'1230'} end={'1400'}>
        会議
      </Schedule>
    </Grid>
  )
}
