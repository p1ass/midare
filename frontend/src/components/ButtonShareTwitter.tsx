import styled from 'styled-components'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faTwitter } from '@fortawesome/free-brands-svg-icons'

import { ButtonBase } from '../atom/ButtonBase'
import { isProd } from '../lib/env'

const Button = styled(ButtonBase)`
  background-color: #1b95e0;
  color: white;
  margin: 1rem;
  border: none;
  width: 20rem;
  @media (max-width: 40rem) {
    width: 90%;
  }
  text-align: center;
`

export const ButtonShareTwitter = ({ shareUrl }: { shareUrl: string }) => {
  return (
    <Button
      href={`https://twitter.com/intent/tweet?url=${shareUrl}&hashtags=生活習慣の乱れを可視化するやつ`}
      onClick={() => {
        if (isProd()) {
          window.gtag('event', 'share', { event_category: 'link', event_label: shareUrl, value: 1 })
        }
      }}
    >
      <FontAwesomeIcon icon={faTwitter} style={{ paddingRight: '0.5rem' }} />
      画像をシェアする
    </Button>
  )
}
