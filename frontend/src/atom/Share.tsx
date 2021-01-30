import styled from 'styled-components'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faTwitter } from '@fortawesome/free-brands-svg-icons'
import { faFacebook } from '@fortawesome/free-brands-svg-icons'

const ShareWrapper = styled.section`
  display: inline-flex;
  flex-direction: row;
  margin-bottom: 1rem;
  justify-content: flex-end;
  align-items: center;
`

const ShareButton = styled.a`
  margin: 0 0.5rem;
  text-decoration: none;
`

const Hatena = styled.i`
color: #4BA3D9;
font-style: normal;
font-variant: normal;
text-rendering: auto;
display: block;
margin:0;
 &:before {
     content: "B!";
     font-family: Verdana;
     font-weight: bold;
     font-size: 28px;
 }
}
`

export const Share = () => {
  return (
    <ShareWrapper>
      <ShareButton
        href="http://twitter.com/intent/tweet?url=https%3A%2F%2Fmidare.p1ass.com&text=生活習慣の乱れを可視化するやつ"
        target="_blank"
        rel="noopener noreferrer"
        title="Tweet"
      >
        <FontAwesomeIcon icon={faTwitter} style={{ color: '#1B95E0', fontSize: '26px' }} />
      </ShareButton>

      <ShareButton
        href="http://www.facebook.com/sharer/sharer.php?u=https%3A%2F%2Fmidare.p1ass.com&t=生活習慣の乱れを可視化するやつ"
        target="_blank"
        rel="noopener noreferrer"
        title="Facebook"
      >
        <FontAwesomeIcon icon={faFacebook} style={{ color: '#3B5999', fontSize: '26px' }} />
      </ShareButton>
      <ShareButton
        href="http://b.hatena.ne.jp/add?mode=confirm&url=https%3A%2F%2Fmidare.p1ass.com&title=生活習慣の乱れを可視化するやつ"
        target="_blank"
        rel="noopener noreferrer"
        title="はてな"
      >
        <Hatena></Hatena>
      </ShareButton>
    </ShareWrapper>
  )
}
