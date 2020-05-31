import React, { useEffect } from 'react'
import window from './AdWindow'
export default function AdsCard() {
  useEffect(() => {
    if (window.adsbygoogle && process.env.NODE_ENV !== 'development') {
      window.adsbygoogle.push({})
    }
  }, [])

  return (
    <ins
      className="adsbygoogle"
      style={{ display: 'block' }}
      data-ad-client="ca-pub-4978327687969784"
      data-ad-slot="6211274963"
      data-ad-format="auto"
      data-full-width-responsive="true"
    ></ins>
  )
}
