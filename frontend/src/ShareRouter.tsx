import { RouteComponentProps } from 'react-router-dom'

const title = '生活習慣の乱れを可視化するやつ'

export const ShareRouter = ({ match }: RouteComponentProps<{ id: string }>) => {
  document.title = title
  const headData = document.head.children
  for (let i = 0; i < headData.length; i++) {
    const nameVal = headData[i].getAttribute('property')
    if (!nameVal) {
      continue
    }

    if (nameVal.indexOf('og:image') !== -1) {
      headData[i].setAttribute(
        'content',
        `https://storage.googleapis.com/midare-share/${match.params.id}.jpeg`
      )
    }
  }
  return null
}
