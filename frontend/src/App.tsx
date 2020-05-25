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

const template = `
"avatar name name .    .    menu"
".      body body body body body"
".      com  like rt   .    ."
`

const GridLayout = styled.div`
  display: grid;
  grid-template: ${template};
  padding: 0.5em;
  margin: 1em;
  grid-gap: 0.2em;
  border: 1px solid #ccc;
  border-radius: 8px;
  align-content: center;
  font-size: 1em;
`

const Area = styled.div<{ area: string }>`
  grid-column: ${(props) => props.area};
  align-self: center;
`

const Avatar = styled.img`
  width: 48px;
`

const Name = styled.div`
  margin: 0;
`

const Body = styled.p``

export function App() {
  return (
    <Wrapper>
      <Timetable></Timetable>
    </Wrapper>
  )
}
