import dayjs from 'dayjs'
import { Period } from './api/client'

type Tweet = {
  text: string
  createdAt: dayjs.Dayjs
}

export type AwakePeriod = {
  okiTime: Tweet
  neTime: Tweet
}

const splitPeriodAtMidnight = (period: Period, okiDate: dayjs.Dayjs, netaDate: dayjs.Dayjs) => {
  const awakePeriods: AwakePeriod[] = []
  let dividedTime = okiDate.add(1, 'date').startOf('date')
  while (!netaDate.isSame(dividedTime, 'date')) {
    awakePeriods.push({
      okiTime: { text: period.okiTime.text, createdAt: okiDate },
      neTime: { text: period.neTime.text, createdAt: dividedTime },
    })
    okiDate = dividedTime
    dividedTime = dividedTime.add(1, 'day')
  }
  awakePeriods.push({
    okiTime: { text: period.okiTime.text, createdAt: dividedTime },
    neTime: { text: period.neTime.text, createdAt: netaDate },
  })
  return awakePeriods
}

export const convertPeriodsToAwakePeriods = (periods: Period[]) => {
  let awakePeriods: AwakePeriod[] = []
  for (const period of periods) {
    const okiDate = dayjs(period.okiTime.createdAt)
    const netaDate = dayjs(period.neTime.createdAt)
    if (okiDate.isSame(netaDate, 'day')) {
      awakePeriods.push({
        okiTime: { text: period.okiTime.text, createdAt: okiDate },
        neTime: { text: period.neTime.text, createdAt: netaDate },
      })
    } else {
      const divided = splitPeriodAtMidnight(period, okiDate, netaDate)
      awakePeriods = awakePeriods.concat(divided.reverse())
    }
  }

  return awakePeriods
}

export const getDatesBetweenLatestAndOldest = (
  oldestDate: dayjs.Dayjs,
  latestDate: dayjs.Dayjs
) => {
  const truncateOldestDate = oldestDate.startOf('date')
  const daysBetweenLatestAndOldest: dayjs.Dayjs[] = [truncateOldestDate]

  let truncateDate = truncateOldestDate
  while (!truncateDate.isSame(latestDate, 'date')) {
    truncateDate = truncateDate.add(1, 'day')

    daysBetweenLatestAndOldest.push(truncateDate)
  }

  return daysBetweenLatestAndOldest
}
