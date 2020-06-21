import React, { useState, useEffect } from 'react'
import GoogleAds from 'react-google-ads'
import styled from 'styled-components'

import { CalendarContainer } from '../components/CalendarContainer'
import { ButtonTwitterLogin } from '../components/ButtonTwitterLogin'
import { Description } from '../components/Description'
import { Header } from '../atom/Header'
import { Footer } from '../atom/Footer'
import { User } from '../entity/User'

import { getMe } from '../api/client'

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

const NewRelease = styled.p`
  font-weight: bold;
`

export const Main = () => {
  const [user, setUser] = useState<User | null>(null)
  const [isFetchUser, setIsFetchUser] = useState(true)

  const SwitchWhetherLogin = !user ? ButtonTwitterLogin : CalendarContainer
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

  return (
    <>
      <Header></Header>
      <Container>
        <FlexContainer>
          <h1>生活習慣の乱れを可視化するやつ</h1>
          <NewRelease>
            新機能 : 画像付きでTwitterにシェア出来るようになりました！
            <span role="img" aria-label="クラッカー">
              🎉
            </span>
          </NewRelease>
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
