import styled from 'styled-components'

import { User } from '../entity/User'

const UserInfoWrapper = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 0.5rem;
  font-weight: 600;
`

type Props = {
  user: User
}

export const CalendarUser = ({ user }: Props) => {
  return (
    <UserInfoWrapper>
      <img src={user.imageUrl} width="48px" height="48px" />
      <p>👇 {user.screenName}さんの生活習慣はこちら👇</p>
    </UserInfoWrapper>
  )
}
