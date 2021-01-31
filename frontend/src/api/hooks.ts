import { useState, useEffect } from 'react'

import { Period } from '../entity/Period'
import { User } from '../entity/User'

import { getMe, getPeriods } from './client'

export const useMe = (): [User | undefined, unknown | undefined, boolean] => {
  const [user, setUser] = useState<User | undefined>(undefined)
  const [error, setError] = useState<unknown | undefined>(undefined)
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    const getUserAsync = async () => {
      try {
        setIsLoading(true)
        const user = await getMe()
        setUser(user)
      } catch (e) {
        setError(e)
      } finally {
        setIsLoading(false)
      }
    }
    getUserAsync()
  }, [])
  return [user, error, isLoading]
}

export const usePeriods = (): [Period[] | undefined, string, unknown] => {
  const [periods, setPeriods] = useState<Period[] | undefined>(undefined)
  const [shareUrl, setShareUrl] = useState('')
  const [error, setError] = useState<unknown>(undefined)
  useEffect(() => {
    const getPeriodsAsync = async () => {
      try {
        const res = await getPeriods()
        if (res.periods.length === 0) {
          return
        }
        setPeriods(res.periods)
        setShareUrl(res.shareUrl)
      } catch (e) {
        setError(e)
      }
    }
    getPeriodsAsync()
  }, [])

  return [periods, shareUrl, error]
}
