import { useState, useEffect } from 'react'

import { Period } from '../entity/Period'

import { getPeriods } from './client'

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
