import styled from 'styled-components'

interface AreaProps {
  row: string
  colStart: string
  colEnd?: string
}
export const Area = styled.div.attrs<AreaProps>(({ row, colStart, colEnd }) => ({
  style: {
    gridRow: row,
    gridColumn: colEnd ? `t-${colStart} / t-${colEnd}` : `t-${colStart}`
  }
}))<AreaProps>``
