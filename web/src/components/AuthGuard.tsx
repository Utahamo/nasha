import { useEffect, type ReactNode } from 'react'
import { useNavigate } from 'react-router-dom'

export default function AuthGuard({ children }: { children: ReactNode }) {
  const navigate = useNavigate()

  useEffect(() => {
    const t = localStorage.getItem('token')
    if (!t) {
      navigate('/login', { replace: true })
    }
  }, [navigate])

  const t = localStorage.getItem('token')
  if (!t) return null

  return <>{children}</>
}
