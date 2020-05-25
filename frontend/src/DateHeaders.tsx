import React from 'react'
import styled from 'styled-components'
import { Area } from './Area'

const HeaderCell = styled(Area)`
  text-align: center;
  height: 100%;
  padding: 0 8px;
  border-top: solid 1px #ccc;
  margin-top: -1px;
`

const DateText = styled.p`
  margin: 0.2rem;
  font-size: 0.9rem;
`

interface DateHeadersProps {
  dateTexts: string[]
}

export const DateHeaders = ({ dateTexts }: DateHeadersProps) => {
  return (
    <>
      {dateTexts.map((dateText) => {
        return (
          <HeaderCell row={dateText} colStart={'header'} key={dateText}>
            <DateText>{dateText}</DateText>
          </HeaderCell>
        )
      })}
    </>
  )
}
