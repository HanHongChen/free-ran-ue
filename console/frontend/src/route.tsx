import { Routes, Route, Navigate } from 'react-router-dom'
import Login from './pages/login'
import NotFound from './pages/not-found'
import Dashboard from './pages/dashboard'
import Gnb from './pages/gnb'
import Ue from './pages/ue'
import { GnbProvider } from './context/GnbContext'

export default function AppRoutes() {
  return (
    <GnbProvider>
      <Routes>
      <Route path="/" element={<Navigate to="/login" replace />} />
      <Route path="/login" element={<Login />} />

      <Route path="dashboard" element={<Dashboard />} />
      <Route path="gnb" element={<Gnb />} />
      <Route path="ue" element={<Ue />} />

      <Route path="/404" element={<NotFound />} />
      <Route path="*" element={<Navigate to="/404" replace />} />
          </Routes>
    </GnbProvider>
  )
}


