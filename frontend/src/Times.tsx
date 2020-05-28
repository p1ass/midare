import React from 'react'

import { rangeTimes } from './Time'
import { Area } from './Area'

export const Times = () => {
  return (
    <>
      {rangeTimes()
        .filter((time) => time.min === '00')
        .map((time, i) => {
          return (
            <Area colStart={`${time.hour}${time.min}`} row={'time'} key={i.toString()}>
              {time.hour}
            </Area>
          )
        })}
    </>
  )
}
