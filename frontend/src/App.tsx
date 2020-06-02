import React, { useState, useEffect } from 'react'
import { BrowserRouter, Route } from 'react-router-dom'
import styled from 'styled-components'

import { ShareRouter } from './ShareRouter'
import { Calendar } from './Calendar'
import { Header } from './Header'
import { Footer } from './Footer'
import { ButtonTwitterLogin } from './ButtonTwitterLogin'
import { Description } from './Description'
import { library } from '@fortawesome/fontawesome-svg-core'
import { fab } from '@fortawesome/free-brands-svg-icons'

import { getMe, User } from './api/client'

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
  max-width: 1000px;
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

  const SwitchWhetherLogin = !user ? ButtonTwitterLogin : Calendar

  const Main = () => {
    return (
      <>
        <Header></Header>
        <Container>
          <FlexContainer>
            <h1>生活習慣の乱れを可視化するやつ</h1>
            {!isFetchUser ? <SwitchWhetherLogin /> : null}
            <Description></Description>
          </FlexContainer>
        </Container>
        <Footer></Footer>
      </>
    )
  }

  return (
    <BrowserRouter>
      <Route path="/share/:id" component={ShareRouter} />
      <Route path="/" component={Main}></Route>
    </BrowserRouter>
  )
}
