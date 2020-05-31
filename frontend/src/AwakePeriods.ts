import dayjs from 'dayjs'
import { Period } from './api/client'

type Tweet = {
  id: string
  text: string
  createdAt: dayjs.Dayjs
  splitDate: dayjs.Dayjs | null
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
      okiTime: {
        id: period.okiTime.id,
        text: period.okiTime.text,
        createdAt: okiDate,
        splitDate: null,
      },
      neTime: {
        id: period.neTime.id,
        text: period.neTime.text,
        createdAt: netaDate,
        splitDate: dividedTime,
      },
    })
    dividedTime = dividedTime.add(1, 'day')
  }
  awakePeriods.push({
    okiTime: {
      id: period.okiTime.id,
      text: period.okiTime.text,
      createdAt: okiDate,
      splitDate: dividedTime,
    },
    neTime: {
      id: period.neTime.id,
      text: period.neTime.text,
      createdAt: netaDate,
      splitDate: null,
    },
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
        okiTime: {
          id: period.okiTime.id,
          text: period.okiTime.text,
          createdAt: okiDate,
          splitDate: null,
        },
        neTime: {
          id: period.neTime.id,
          text: period.neTime.text,
          createdAt: netaDate,
          splitDate: null,
        },
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
