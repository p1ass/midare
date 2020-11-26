import styled from 'styled-components'
import { ButtonGitHub } from './ButtonGitHub'
import { Share } from '../atom/Share'

const WhatIsThis = styled.p`
  text-align: center;
`

const Ul = styled.ul`
  margin: 0 0 2rem 0;
`

const Li = styled.li`
  padding: 0.3rem 0;
`

export const Description = () => {
  return (
    <>
      <h2>これは何</h2>
      <WhatIsThis>
        ツイートを使って生活習慣の乱れを可視化するWebアプリです。
        <br />
        カレンダーUIで直感的に起床・就寝時間の変化を見ることができます。
        <br />
        ツイート数が多ければ多いほど精度が高くなります。
      </WhatIsThis>
      <h2>仕組み</h2>
      <ol>
        <li>Twitter APIを使って直近のツイートを収集します。</li>
        <li>ツイートの間隔が3.5時間以内であれば、その時間帯は起きているとみなします。</li>
      </ol>
      <h2>シェア</h2>
      <Share></Share>
      <h2>ソースコードはこちら↓</h2>
      <ButtonGitHub></ButtonGitHub>
      <h2>作成者</h2>
      <Ul>
        <Li>
          Twitter : <a href="https://twitter.com/p1ass">@p1ass</a>
        </Li>
        <Li>
          GitHub : <a href="https://github.com/p1ass">@p1ass</a>
        </Li>
      </Ul>
      <p>利用状況の計測のためにGoogle Analyticsを利用しています。</p>
    </>
  )
}
