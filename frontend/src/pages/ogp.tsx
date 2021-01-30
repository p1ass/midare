import { useState, useEffect } from 'react'
import styled from 'styled-components'
import { Calendar } from '../components/Calendar'
import { Period } from '../entity/Period'

declare global {
  interface Window {
    getPeriods: () => Period[]
  }
}

const sleep = (msec: number) => new Promise((resolve) => setTimeout(resolve, msec))

const Flex = styled.div`
  display: flex;
  height: 100vh;
  width: 100vw;
  align-items: center;
  justify-content: center;
`

export const OGP = () => {
  const [periods, setPeriods] = useState<Period[]>([])
  useEffect(() => {
    const getPeriodsAsync = async () => {
      while (!window.getPeriods) {
        await sleep(100)
      }
      setPeriods(await window.getPeriods())
    }
    getPeriodsAsync()
  }, [])

  return periods.length !== 0 ? (
    <Flex className="ogp-calendar-flex">
      <Calendar periods={periods} generatingImage={true}></Calendar>
    </Flex>
  ) : null
}

export default OGP
