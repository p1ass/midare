import React from 'react'
import styled from 'styled-components'

const HeaderWrapper = styled.header`
  width: 100vw;
  height: 2rem;
  line-height: 2rem;
  background-color: rgb(88, 149, 98);
  color: white;
  font-weight: bold;
  padding-left: 1rem;
`

export const Header = () => {
  return <HeaderWrapper>生活習慣の乱れを可視化するやつ for ツイ廃</HeaderWrapper>
}
