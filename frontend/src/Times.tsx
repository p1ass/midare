import React from 'react'
import styled from 'styled-components'

import { rangeTimes } from './Time'
import { Area } from './Area'

const Hour = styled(Area)`
  margin: 4px 0;
  font-size: 1rem;
  min-width: 1rem;
`

export const Times = () => {
  return (
    <>
      {rangeTimes()
        // .filter((time) => time.min === '00')
        .map((time, i) => {
          return (
            <Hour colStart={`${time.hour}${time.min}`} row={'time'} key={i.toString()}>
              {time.min === '00' ? time.hour : ''}
            </Hour>
          )
        })}
    </>
  )
}
