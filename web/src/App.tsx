import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom'
import Login from './pages/Login'
import FileBrowser from './pages/FileBrowser'
import Settings from './pages/Settings'

/**
 * Root application component.
 *
 * Route layout:
 *   /login          – public login page
 *   /               – redirect to /files
 *   /files/*path    – main file browser (auth-protected in future)
 *   /settings       – admin / settings panel (auth-protected in future)
 *
 * TODO: wrap protected routes with an AuthGuard that checks for a valid JWT.
 */
export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/files/*" element={<FileBrowser />} />
        <Route path="/settings" element={<Settings />} />
        <Route path="/" element={<Navigate to="/files" replace />} />
      </Routes>
    </BrowserRouter>
  )
}

