import styled from 'styled-components'

const HeaderWrapper = styled.header`
  width: 100%;
  height: 2rem;
  line-height: 2rem;
  background-color: rgb(88, 149, 98);
  color: white;
  font-weight: bold;
`

const Title = styled.span`
  padding-left: 1rem;
`

export const Header = () => {
  return (
    <HeaderWrapper>
      <Title>生活習慣の乱れを可視化するやつ</Title>
    </HeaderWrapper>
  )
}
