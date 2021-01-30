import { useRouter } from 'next/router'

const Callback = () => {
  const router = useRouter()
  if (typeof window !== 'undefined') {
    router.push('/')
  }
  return null
}

export default Callback
