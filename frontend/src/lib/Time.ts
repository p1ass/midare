import dayjs from 'dayjs'

// 時間を30分単位で出力
export const rangeTimes = (start = 0, hours = 24): dayjs.Dayjs[] => {
  return Array.from({ length: hours * 2 }, (_, i) => {
    const hr = Math.floor(i / 2) + start
    const min = (i % 2) * 30
    return dayjs().set('hour', hr).set('minute', min)
  })
}
