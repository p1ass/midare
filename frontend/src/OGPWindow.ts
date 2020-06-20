import { Period } from './api/client'

declare global {
  interface Window {
    getPeriods: () => Period[]
  }
}
