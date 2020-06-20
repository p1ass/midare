import React, { useState, useEffect } from 'react'
import { BrowserRouter, Route, Switch } from 'react-router-dom'
import styled from 'styled-components'
import GoogleAds from 'react-google-ads'
import { CalendarContainer } from './CalendarContainer'
import { Header } from './Header'
import { Footer } from './Footer'
import { ButtonTwitterLogin } from './ButtonTwitterLogin'
import { Description } from './Description'
import { OGPCalendar } from './OGP'

import { library } from '@fortawesome/fontawesome-svg-core'
import { fab } from '@fortawesome/free-brands-svg-icons'

import { getMe, User } from './api/client'
import { ShareRouter } from './ShareRouter'

library.add(fab)

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

export function App() {
  const [user, setUser] = useState<User | null>(null)
  const [isFetchUser, setIsFetchUser] = useState(true)

  useEffect(() => {
    const getUserAsync = async () => {
      try {
        setIsFetchUser(true)
        const res = await getMe()
        setUser(res)
        setIsFetchUser(false)
      } catch (e) {
        setUser(null)
        setIsFetchUser(false)
      }
    }
    getUserAsync()
  }, [])

  const SwitchWhetherLogin = !user ? ButtonTwitterLogin : CalendarContainer

  const Main = () => {
    return (
      <>
        <Header></Header>
        <Container>
          <FlexContainer>
            <h1>生活習慣の乱れを可視化するやつ</h1>
            <p>
              新機能 : Twitterに画像付きでシェア出来るようになりました！
              <span role="img" aria-label="クラッカー">
                🎉
              </span>
            </p>
            {!isFetchUser ? <SwitchWhetherLogin /> : null}
            <GoogleAds
              client="ca-pub-4978327687969784"
              slot="6211274963"
              className="adsbygoogle"
              format="auto"
              style={{ display: 'block' }}
            />
            <Description></Description>
          </FlexContainer>
        </Container>
        <Footer></Footer>
      </>
    )
  }

  return (
    <BrowserRouter>
      <Switch>
        <Route path="/ogp" component={OGPCalendar}></Route>
        <Route path="/share/:id" component={ShareRouter}></Route>
        <Route path="/" component={Main}></Route>
      </Switch>
    </BrowserRouter>
  )
}
