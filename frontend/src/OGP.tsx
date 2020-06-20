import React from 'react'

import { Calendar } from './Calendar'
import window from './OGPWindow'

export const OGPCalendar = () => {
  return <Calendar periods={window.periods} generatingImage={true}></Calendar>
}
