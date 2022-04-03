import GoogleAds from 'react-google-ads'
import styled from 'styled-components'

import { ButtonTwitterLogin } from '../components/ButtonTwitterLogin'
import { Description } from '../components/Description'
import { Header } from '../atom/Header'
import { Footer } from '../atom/Footer'
import { CalendarContainer } from '../components/CalendarContainer'
import { useMe } from '../api/hooks'
import { useHasTouchScreen } from '../lib/screen'

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

const MessageBrowser = styled.p`
  font-size: 0.9rem;
  text-align: center;
  color: #ff0000;
  width: 80vw;
`

const Title = styled.h1`
  text-align: center;
`

const Main = () => {
  const [user, , isLoading] = useMe()
  const { hasTouchScreen } = useHasTouchScreen()
  const SwitchWhetherLogin = user ? (
    <CalendarContainer user={user}></CalendarContainer>
  ) : (
    <>
      {hasTouchScreen ? (
        <MessageBrowser>
          Twitterアプリ内で開いた場合は、 SafariやChromeで開き直してください
        </MessageBrowser>
      ) : null}
      <ButtonTwitterLogin></ButtonTwitterLogin>
    </>
  )

  return (
    <>
      <Header />
      <Container>
        <FlexContainer>
          <Title>
            <span className="ww">生活習慣</span>
            <span className="ww">の</span>
            <span className="ww">乱れを</span>
            <span className="ww">可視化</span>
            <span className="ww">する</span>
            <span className="ww">やつ</span>
          </Title>
          {/* <OneYearMessage>祝1周年🎉</OneYearMessage>
          <OneYearDescription>いつもお使いいただきありがとうございます！</OneYearDescription> */}

          {!isLoading ? SwitchWhetherLogin : null}
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
