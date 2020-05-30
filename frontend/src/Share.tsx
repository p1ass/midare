import React from 'react'
import styled from 'styled-components'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'

const ShareWrapper = styled.section`
  display: flex;
  flex-direction: row;
  margin-bottom: 1rem;
`

const ShareButton = styled.a`
  margin: 0 0.5rem;
`

export const Share = () => {
  return (
    <ShareWrapper>
      <ShareButton
        href="http://twitter.com/intent/tweet?url=https%3A%2F%2Fmidare.p1ass.com&text=生活習慣の乱れを可視化するやつ for ツイ廃"
        target="_blank"
        rel="noopener noreferrer"
        title="Tweet"
      >
        <FontAwesomeIcon icon={['fab', 'twitter']} style={{ color: '#1B95E0' }} size="2x" />
      </ShareButton>

      <ShareButton
        href="http://www.facebook.com/sharer/sharer.php?u=https%3A%2F%2Fmidare.p1ass.com&t=生活習慣の乱れを可視化するやつ for ツイ廃 - ぷらすのブログ"
        target="_blank"
        rel="noopener noreferrer"
        title="Facebook"
      >
        <FontAwesomeIcon icon={['fab', 'facebook']} style={{ color: '#3B5999' }} size="2x" />
      </ShareButton>
    </ShareWrapper>
  )
}
