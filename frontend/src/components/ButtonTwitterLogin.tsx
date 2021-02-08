import styled from 'styled-components'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faTwitter } from '@fortawesome/free-brands-svg-icons'

import { getLoginUrl } from '../api/client'
import { isProd } from '../lib/env'
import { ButtonBase } from '../atom/ButtonBase'

const TwitterButton = styled(ButtonBase)`
  background-color: rgb(27, 149, 224);
  color: white;
  margin-bottom: 1rem;
`

export const ButtonTwitterLogin = () => {
  return (
    <TwitterButton
      href={getLoginUrl()}
      onClick={() => {
        if (isProd()) {
          window.gtag('event', 'login', {
            event_category: 'login',
            event_label: 'twitter',
            value: 1,
          })
        }
      }}
    >
      <FontAwesomeIcon icon={faTwitter} style={{ paddingRight: '0.5rem' }} />
      乱れを可視化する
    </TwitterButton>
  )
}
