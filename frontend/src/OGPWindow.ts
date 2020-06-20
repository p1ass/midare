import { Period } from './api/client'

interface OGPWindow extends Window {
  periods: Period[]
}
declare let window: OGPWindow
export default window
