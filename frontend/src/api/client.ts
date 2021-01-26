import axios from 'axios'

import { Period } from '../entity/Period'
import { User } from '../entity/User'

const baseURL = process.env.REACT_APP_API_BASE_URL || 'http://localhost.local:8080'

const instance = axios.create({
  baseURL: baseURL,
  withCredentials: true,
})

interface GetPeriodsResponse {
  periods: Period[]
  shareUrl: string
}

export const getLoginUrl = () => {
  return baseURL + '/login'
}

export const getMe = async () => {
  const res = await instance.get<User>('/me')
  window.gtag('event', 'login_succeed', {
    value: res.data.screenName,
  })
  return res.data
}

export const getPeriods = async () => {
  const res = await instance.get<GetPeriodsResponse>('/periods')
  window.gtag('event', 'periods_got', {
    share_url: res.data.shareUrl,
    value: res.data.periods.length,
  })
  return res.data
}
