import React, { useState, useEffect } from 'react'
import styled from 'styled-components'
import { Calendar } from '../components/Calendar'
import { Period } from '../entity/Period'

const sleep = (msec: number) => new Promise((resolve) => setTimeout(resolve, msec))

const Flex = styled.div`
  display: flex;
  height: 100vh;
  width: 100vw;
  align-items: center;
  justify-content: center;
`

export const OGPCalendar = () => {
  const [periods, setPeriods] = useState<Period[]>([])
  useEffect(() => {
    const getPeriodsAsync = async () => {
      await sleep(600)
      setPeriods(await window.getPeriods())
    }
    getPeriodsAsync()
  }, [])

  return periods.length !== 0 ? (
    <Flex>
      <Calendar periods={periods} generatingImage={true}></Calendar>
    </Flex>
  ) : null
}
