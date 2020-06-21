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
  return res.data
}

export const getPeriods = async () => {
  const res = await instance.get<GetPeriodsResponse>('/periods')
  return res.data
}

interface UploadImageResponse {
  shareUrl: string
}

export const uploadImage = async (image: Blob) => {
  const params = new FormData()

  params.append('file', image)
  const res = await instance.post<UploadImageResponse>('/images', params)

  return res.data
}
