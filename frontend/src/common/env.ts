export const IsProd = () => {
  return process.env.NODE_ENV === 'production'
}
