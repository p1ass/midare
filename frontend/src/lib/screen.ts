import { useEffect, useState } from 'react'

export const useHasTouchScreen = () => {
  const [state, setState] = useState(false)

  useEffect(() => {
    setState(hasTouchScreen())
  }, [])

  return {
    hasTouchScreen: state,
  } as const
}

const hasTouchScreen = () => {
  if (navigator.maxTouchPoints > 0) {
    return true
  }
  if (window.matchMedia('(pointer:coarse)').matches) {
    return true
  }
  if ('orientation' in window) {
    return true
  }

  return false
}
