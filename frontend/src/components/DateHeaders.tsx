import styled from 'styled-components'

import { Area } from '../atom/Area'

const HeaderCell = styled(Area)`
  padding: 0px;
  border-bottom: solid 1px #ccc;
  margin-top: -1px;
`

const DateText = styled.p<{ generatingImage: boolean }>`
  margin: 0.2rem;
  font-size: 0.9rem;
  @media (max-width: 60rem) {
    font-size: 0.7rem;
  }
  @media (max-width: 40rem) {
    font-size: ${({ generatingImage }) => (generatingImage ? '0.1rem' : '0.4rem')};
    width: 3.4rem;
  }
`

interface DateHeadersProps {
  dateTexts: string[]
  generatingImage: boolean
}

export const DateHeaders = ({ dateTexts, generatingImage }: DateHeadersProps) => {
  return (
    <>
      {[''].concat(dateTexts).map((dateText) => {
        return (
          <HeaderCell row={dateText} colStart={'header'} key={dateText}>
            <DateText generatingImage={generatingImage}>{dateText}</DateText>
          </HeaderCell>
        )
      })}
    </>
  )
}
