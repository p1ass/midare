import styled from 'styled-components'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { ButtonBase } from '../atom/ButtonBase'

const Button = styled(ButtonBase)`
  background-color: #1b95e0;
  color: white;
  margin: 1rem;
  border: none;
`

export const ButtonShareTwitter = ({ href }: { href: string }) => {
  return (
    <Button
      href={href}
      onClick={() =>
        window.gtag('event', 'share', { event_category: 'link', event_label: href, value = 1 })
      }
    >
      <FontAwesomeIcon icon={['fab', 'twitter']} style={{ paddingRight: '0.5rem' }} />
      画像付きでTwitterにシェア
    </Button>
  )
}
