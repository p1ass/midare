import styled from 'styled-components'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { getLoginUrl } from '../api/client'
import { IsProd } from '../lib/env'
import { faTwitter } from '@fortawesome/free-brands-svg-icons'

import { ButtonBase } from '../atom/ButtonBase'

const TwitterButton = styled(ButtonBase)`
  background-color: rgb(27, 149, 224);
  color: white;
  margin-bottom: 2rem;
`

export const ButtonTwitterLogin = () => {
  return (
    <TwitterButton
      href={getLoginUrl()}
      onClick={() => {
        if (IsProd()) {
          window.gtag('event', 'login', {
            event_category: 'login',
            event_label: 'twitter',
            value: 1,
          })
        }
      }}
    >
      <FontAwesomeIcon icon={faTwitter} style={{ paddingRight: '0.5rem' }} />
      起きている時間の変化を見る
    </TwitterButton>
  )
}
