import Head from 'next/head'
import { GetServerSideProps } from 'next'
import { useRouter } from 'next/router'

type Prop = {
  id?: string
}

const Share = ({ id }: Prop) => {
  const router = useRouter()
  if (typeof window !== 'undefined') {
    router.push('/')
  }
  return (
    <Head>
      {id ? (
        <meta
          property="og:image"
          content={`https://storage.googleapis.com/midare-share/${id}.jpg`}
          key="ogImage"
        />
      ) : null}
    </Head>
  )
}

export const getServerSideProps: GetServerSideProps = async (ctx) => {
  return { props: { id: ctx.params?.id } }
}

export default Share
