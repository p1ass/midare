import GoogleAds from 'react-google-ads'
import styled from 'styled-components'

import { ButtonTwitterLogin } from '../components/ButtonTwitterLogin'
import { Description } from '../components/Description'
import { Header } from '../atom/Header'
import { Footer } from '../atom/Footer'
import { CalendarContainer } from '../components/CalendarContainer'
import { useMe } from '../api/hooks'

const FlexContainer = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  color: #333;
  margin: 0 1rem;
`

const Container = styled.div`
  margin: 0 auto;
`

const Main = () => {
  const [user, , isLoading] = useMe()
  const SwitchWhetherLogin = user ? CalendarContainer : ButtonTwitterLogin

  return (
    <>
      <Header />
      <Container>
        <FlexContainer>
          <h1>生活習慣の乱れを可視化するやつ</h1>
          {!isLoading ? <SwitchWhetherLogin /> : null}
          <GoogleAds
            client="ca-pub-4978327687969784"
            slot="6211274963"
            className="adsbygoogle"
            format="auto"
            style={{ display: 'block' }}
          />
          <Description />
        </FlexContainer>
      </Container>
      <Footer />
    </>
  )
}

export default Main
