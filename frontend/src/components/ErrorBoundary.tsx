import React from 'react'

interface State {
  hasError: Boolean
}

export class ErrorBoundary extends React.Component<{}, State> {
  constructor() {
    super({})
    this.state = { hasError: false }
  }

  static getDerivedStateFromError() {
    // Update state so the next render will show the fallback UI.
    return { hasError: true }
  }

  componentDidCatch() {}

  render() {
    if (this.state.hasError) {
      // You can render any custom fallback UI
      return null
    }

    return this.props.children
  }
}
