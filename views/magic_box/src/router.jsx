import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import React from 'react'
import Index from './pages/index'
import Login from './pages/login'

export const LoginUrlPath = '/login'

export default function AppRouter () {
  return (
      <Router>
          <Routes>
              <Route index element={<Index/>}/>
              <Route path={LoginUrlPath} element={<Login/>}/>
          </Routes>
    </Router>
  )
}
