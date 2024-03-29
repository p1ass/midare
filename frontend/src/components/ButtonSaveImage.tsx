import styled from 'styled-components'

import { isProd } from '../lib/env'
import { ButtonBase } from '../atom/ButtonBase'

const Button = styled(ButtonBase)`
  background-color: #7f8c8d;
  color: white;
  margin: 1rem;
  border: none;
`

export const ButtonSaveImage = ({ onClick }: { onClick: () => Promise<void> }) => {
  return (
    <Button
      as="button"
      onClick={() => {
        if (isProd()) {
          window.gtag('event', 'image_saved', {
            event_category: 'image',
            value: 1,
          })
        }
        onClick()
      }}
    >
      画像ファイルとして保存
    </Button>
  )
}
