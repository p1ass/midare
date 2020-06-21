import React from 'react'
import { BrowserRouter, Route, Switch } from 'react-router-dom'
import { library } from '@fortawesome/fontawesome-svg-core'
import { fab } from '@fortawesome/free-brands-svg-icons'

import { OGPCalendar } from './pages/OGPCalendar'
import { Main } from './pages/Main'
import { ShareRouter } from './pages/ShareRouter'

library.add(fab)

export function App() {
  return (
    <BrowserRouter>
      <Switch>
        <Route path="/ogp" component={OGPCalendar}></Route>
        <Route path="/share/:id" component={ShareRouter}></Route>
        <Route path="/" component={Main}></Route>
      </Switch>
    </BrowserRouter>
  )
}
