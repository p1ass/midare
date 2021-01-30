import styled from 'styled-components'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faGithub } from '@fortawesome/free-brands-svg-icons'

import { ButtonBase } from '../atom/ButtonBase'

const GitHubButton = styled(ButtonBase)`
  background-color: #171515;
  color: white;
  margin: 1rem;
`

export const ButtonGitHub = () => {
  return (
    <GitHubButton href="https://github.com/p1ass/midare">
      <FontAwesomeIcon icon={faGithub} style={{ paddingRight: '0.5rem' }} />
      GitHubを開く
    </GitHubButton>
  )
}
