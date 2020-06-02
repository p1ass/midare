import { RouteComponentProps } from 'react-router-dom'

const description =
  'ツイートを使って生活習慣の乱れを可視化するWebアプリです。カレンダーUIで直感的に起床・就寝時間の変化を見ることが出来ます。'

export const ShareRouter = ({ match }: RouteComponentProps<{ id: string }>) => {
  document.title = '生活習慣の乱れを可視化するやつ'

  const headData = document.head.children
  for (let i = 0; i < headData.length; i++) {
    const nameVal = headData[i].getAttribute('name')
    if (nameVal?.indexOf('description') !== -1) {
      headData[i].setAttribute('content', description)
    }
    // OGP(twitter)の設定
    if (nameVal?.indexOf('twitter:image') !== -1) {
      headData[i].setAttribute(
        'content',
        `https://storage.googleapis.com/midare-share/${match.params.id}.jpeg`
      )
    }
    if (nameVal?.indexOf('twitter:description') !== -1) {
      headData[i].setAttribute('content', description)
    }
  }
  return null
}
