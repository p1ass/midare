export type Time = {
  hour: string
  min: string
}

// 時間を30分単位で出力
export const rangeTimes = (start = 6, hours = 24): Time[] => {
  return Array.from({ length: hours * 2 }, (_, i) => {
    const hr = Math.floor(i / 2) + start
    const min = (i % 2) * 30
    const minSprint = min === 0 ? '00' : min
    const hrSprint = hr < 10 ? `0${hr}` : hr
    return { hour: hrSprint.toString(), min: minSprint.toString() }
  })
}
