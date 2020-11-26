import styled from 'styled-components'

const FooterWrapper = styled.footer`
  width: 100%;
  height: 2rem;
  line-height: 2rem;
  background-color: rgb(88, 149, 98);
  color: white;
  text-align: center;
`

export const Footer = () => {
  return <FooterWrapper>© 2020 - p1ass All Rights Reserved.</FooterWrapper>
}
