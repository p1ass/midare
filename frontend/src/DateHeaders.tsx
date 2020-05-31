import React from 'react'
import styled from 'styled-components'
import { Area } from './Area'

const HeaderCell = styled(Area)`
  height: 100%;
  padding: 0 8px;
  border-bottom: solid 1px #ccc;
  margin-top: -1px;
`

const DateText = styled.p`
  margin: 0.2rem;
  font-size: 0.9rem;
  @media (max-width: 60rem) {
    font-size: 0.7rem;
  }
  @media (max-width: 40rem) {
    font-size: 0.4rem;
  }
`

interface DateHeadersProps {
  dateTexts: string[]
}

export const DateHeaders = ({ dateTexts }: DateHeadersProps) => {
  return (
    <>
      {[''].concat(dateTexts).map((dateText) => {
        return (
          <HeaderCell row={dateText} colStart={'header'} key={dateText}>
            <DateText>{dateText}</DateText>
          </HeaderCell>
        )
      })}
    </>
  )
}
