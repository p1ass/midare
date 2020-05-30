import React, { useState, useEffect } from 'react'
import styled from 'styled-components'
import { Calendar } from './Calendar'
import { Header } from './Header'
import { ButtonTwitterLogin } from './ButtonTwitterLogin'
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
`

export function App() {
  const [user, setUser] = useState<User | null>({ id: '', name: '', screenName: '', imageUrl: '' })

  useEffect(() => {
    const getUserAsync = async () => {
      try {
        const res = await getMe()
        setUser(res)
      } catch (e) {
        setUser(null)
      }
    }
    getUserAsync()
  }, [])
  return (
    <>
      <Header></Header>
      <FlexContainer>
        {user === null ? <ButtonTwitterLogin></ButtonTwitterLogin> : <Calendar />}
      </FlexContainer>
    </>
  )
}
