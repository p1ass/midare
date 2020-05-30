import React from 'react'
import styled from 'styled-components'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { getLoginUrl } from './api/client'

import { ButtonBase } from './ButtonBase'

const TwitterButton = styled(ButtonBase)`
  background-color: #00acee;
  color: white;
`

export const ButtonTwitterLogin = () => {
  return (
    <TwitterButton href={getLoginUrl()}>
      <FontAwesomeIcon icon={['fab', 'twitter']} style={{ paddingRight: '0.5rem' }} />
      Twitterでログイン
    </TwitterButton>
  )
}
