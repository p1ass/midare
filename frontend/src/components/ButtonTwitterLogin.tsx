import styled from 'styled-components'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { getLoginUrl } from '../api/client'

import { ButtonBase } from '../atom/ButtonBase'

const TwitterButton = styled(ButtonBase)`
  background-color: rgb(27, 149, 224);
  color: white;
  margin-bottom: 2rem;
`

export const ButtonTwitterLogin = () => {
  return (
    <TwitterButton href={getLoginUrl()}>
      <FontAwesomeIcon icon={['fab', 'twitter']} style={{ paddingRight: '0.5rem' }} />
      Twitterでログイン
    </TwitterButton>
  )
}
