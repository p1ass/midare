import Image from 'next/image'
import styled from 'styled-components'

import { User } from '../entity/User'

const UserInfoWrapper = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 0.5rem;
  font-weight: 600;
`

const UserDescription = styled.p`
  margin-left: 8px;
`

const UserIcon = styled(Image)`
  border-radius: 48px;
`

type Props = {
  user: User
}

export const CalendarUser = ({ user }: Props) => {
  return (
    <UserInfoWrapper>
      <UserIcon src={user.imageUrl} width="48px" height="48px" />
      <UserDescription>{user.name}さんの生活習慣はこちら！</UserDescription>
    </UserInfoWrapper>
  )
}
