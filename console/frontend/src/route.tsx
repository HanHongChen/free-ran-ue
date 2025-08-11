import { Routes, Route, Navigate } from 'react-router-dom'
import Login from './pages/login'
import NotFound from './pages/not-found'
import Dashboard from './pages/dashboard'

export default function AppRoutes() {
  return (
      <Routes>
          <Route path="/" element={<Navigate to="/login" replace />} />
          <Route path="/login" element={<Login />} />
          <Route path="/dashboard" element={<Dashboard />} />
          <Route path="/404" element={<NotFound />} />
          <Route path="*" element={<Navigate to="/404" replace />} />
      </Routes>
  )
}


