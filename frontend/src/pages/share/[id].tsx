import Head from 'next/head'
import { useRouter } from 'next/router'

const Share = () => {
  const router = useRouter()
  router.push('/')

  const { id } = router.query
  return (
    <Head>
      <meta property="og:image" content={`https://storage.googleapis.com/midare-share/${id}.jpg`} />
    </Head>
  )
}

export async function getServerSideProps() {
  return { props: {} }
}

export default Share
