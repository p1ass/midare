import React from 'react'
import styled from 'styled-components'
import { Timetable } from './Calendar'

const Wrapper = styled.div`
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  color: #333;
`

export function App() {
  return (
    <Wrapper>
      <Timetable></Timetable>
    </Wrapper>
  )
}
