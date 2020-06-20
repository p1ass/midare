import { Period } from './api/client'

declare global {
  interface Window {
    periods: Period[]
    getPeriods: () => Period[]
  }
}
