import React, { useState } from 'react'
import styled from 'styled-components'
import dayjs from 'dayjs'
import Modal from 'react-modal'

import { Area } from './Area'
import { AwakePeriod } from './AwakePeriods'

const ScheduleBlock = styled(Area)`
  background: rgb(88, 149, 98);
  border-radius: 4px;
  font-weight: bold;
  margin: 0.1rem 0;
  color: #eee;
  font-size: 0.5rem;
  cursor: pointer;
`

Modal.setAppElement('#root')
const customModalStyles = {
  content: {
    top: '50%',
    left: '50%',
    right: 'auto',
    bottom: 'auto',
    marginRight: '-40%',
    transform: 'translate(-50%, -50%)',
  },
}

interface AwakeScheduleProps {
  awakePeriod: AwakePeriod
}

const AwakeSchedule = ({ awakePeriod }: AwakeScheduleProps) => {
  const [isOpen, setIsOpen] = useState(false)

  const okiTime = awakePeriod.okiTime.splitDate
    ? awakePeriod.okiTime.splitDate
    : awakePeriod.okiTime.createdAt
  const neTime = awakePeriod.neTime.splitDate
    ? awakePeriod.neTime.splitDate
    : awakePeriod.neTime.createdAt

  const okiTimeTrunate = truncateDate(okiTime)
  const neTimeTruncate = truncateDate(neTime)
  return (
    <>
      <ScheduleBlock
        colStart={okiTimeTrunate.format('HHmm')}
        colEnd={
          okiTimeTrunate.hour() !== 0 && neTimeTruncate.hour() === 0
            ? '2400'
            : neTimeTruncate.format('HH') + neTimeTruncate.format('mm')
        }
        row={okiTimeTrunate.format('MMMMDD')}
        onClick={() => {
          setIsOpen(true)
        }}
      ></ScheduleBlock>
      <Modal
        isOpen={isOpen}
        shouldCloseOnOverlayClick={true}
        onRequestClose={() => setIsOpen(false)}
        contentLabel="ツイート詳細"
        style={customModalStyles}
      >
        <h3>起床直後のツイート</h3>
        <span>{awakePeriod.okiTime.createdAt.format('MM/DD HH:mm')}</span>
        <p>{awakePeriod.okiTime.text}</p>
        <h3>就寝直前のツイート</h3>
        <span>{awakePeriod.neTime.createdAt.format('MM/DD HH:mm')}</span>
        <p>{awakePeriod.neTime.text}</p>
      </Modal>
    </>
  )
}

interface AwakeSchedulesProps {
  awakePeriods: AwakePeriod[]
}

const truncateDate = (date: dayjs.Dayjs) => {
  if (date.minute() < 15) {
    return date.startOf('hour')
  }
  if (date.minute() >= 15 && date.minute() < 45) {
    return date.startOf('hour').add(30, 'minute')
  }
  return date.add(1, 'hour').startOf('hour')
}

export const AwakeSchedules = ({ awakePeriods }: AwakeSchedulesProps) => {
  return (
    <>
      {awakePeriods.map((awakePeriod, idx) => {
        return <AwakeSchedule key={idx} awakePeriod={awakePeriod}></AwakeSchedule>
      })}
    </>
  )
}
