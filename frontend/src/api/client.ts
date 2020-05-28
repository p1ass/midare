import axios from 'axios'

const baseURL = process.env.REACT_APP_API_BASE_URL || 'http://localhost.local:8080'

const instance = axios.create({
  baseURL: baseURL,
  withCredentials: true,
})

interface Tweet {
  id: string
  text: string
  createdAt: string
}

export interface Period {
  okiTime: Tweet
  neTime: Tweet
}

interface GetMeResponse {
  id: string
  name: string
  screenName: string
  imageUrl: string
}

interface GetPeriodsResponse {
  periods: Period[]
}

export const getLoginUrl = () => {
  return baseURL + '/login'
}

export const getMe = async () => {
  const res = await instance.get<GetMeResponse>('/me')
  return res.data
}

export const getPeriods = async () => {
  const res = await instance.get<GetPeriodsResponse>('/periods')
  return res.data
}
