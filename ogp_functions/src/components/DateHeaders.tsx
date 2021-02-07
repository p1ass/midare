import styled from 'styled-components'

import { Area } from '../atom/Area'

const HeaderCell = styled(Area)`
  padding: 0px;
  border-bottom: solid 1px #ccc;
  margin-top: -1px;
`

const DateText = styled.p`
  margin: 0.2rem;
  font-size: 0.8rem;
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
