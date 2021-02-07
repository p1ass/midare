import styled from 'styled-components'

import { rangeTimes } from '../lib/time'
import { Area } from '../atom/Area'

const Hour = styled(Area)`
  margin: 4px 0;
  font-size: 1rem;
  min-width: 1rem;
`

export const Times = ({ row }: { row: string }) => {
  return (
    <>
      {rangeTimes().map((time, i) => (
        <Hour colStart={time.format('HHmm')} row={row} key={i.toString()}>
          {time.minute() === 0 ? time.hour() : ''}
        </Hour>
      ))}
    </>
  )
}
