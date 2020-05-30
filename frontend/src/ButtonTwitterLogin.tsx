import React from 'react'
import styled from 'styled-components'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { getLoginUrl } from './api/client'

const TwitterButton = styled.a`
  background-color: #00acee;
  color: white;
  padding: 1rem;
  margin-top: 1rem;
  border-radius: 0.5rem;
  text-decoration: none;
  font-weight: bold;
  &:visited {
    color: white;
  }
`

export const ButtonTwitterLogin = () => {
  return (
    <TwitterButton href={getLoginUrl()}>
      <FontAwesomeIcon icon={['fab', 'twitter']} style={{ paddingRight: '0.5rem' }} />
      Twitterでログイン
    </TwitterButton>
  )
}
