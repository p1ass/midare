import styled from 'styled-components'

import { rangeTimes } from '../lib/Time'
import { Area } from '../atom/Area'

const Hour = styled(Area)<{ generatingImage: boolean }>`
  margin: 4px 0;
  font-size: 1rem;
  min-width: 1rem;
  @media (max-width: 60rem) {
    font-size: 0.7rem;
    min-width: 0.7rem;
  }
  @media (max-width: 40rem) {
    font-size: ${({ generatingImage }) => (generatingImage ? '1rem' : '0.2rem')};
    min-width: ${({ generatingImage }) => (generatingImage ? '1.5rem' : '0.2rem')};
  }
`

export const Times = ({ row, generatingImage }: { row: string; generatingImage: boolean }) => {
  return (
    <>
      {rangeTimes().map((time, i) => (
        <Hour
          generatingImage={generatingImage}
          colStart={time.format('HHmm')}
          row={row}
          key={i.toString()}
        >
          {time.minute() === 0 ? time.hour() : ''}
        </Hour>
      ))}
    </>
  )
}
