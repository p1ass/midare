import { Period } from '../entity/Period'

declare global {
  interface Window {
    getPeriods: () => Period[]
  }
}
