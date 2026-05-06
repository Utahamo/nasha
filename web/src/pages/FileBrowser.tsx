import { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'

interface FileEntry {
  Name: string
  Size: number
  IsDir: boolean
  ModTime: string
  Path: string
}

export default function FileBrowser() {
  const { '*': splat } = useParams()
  const currentPath = '/' + (splat || '')
  const [entries, setEntries] = useState<FileEntry[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  function token() {
    return localStorage.getItem('token') || ''
  }

  function apiUrl(path: string) {
    return '/api/v1/fs' + path + '?token=' + encodeURIComponent(token())
  }

  async function load(path: string) {
    setLoading(true)
    setError('')
    try {
      const res = await fetch('/api/v1/fs' + path, {
        headers: { Authorization: 'Bearer ' + token() },
      })
      if (res.status === 401) {
        localStorage.removeItem('token')
        window.location.href = '/login'
        return
      }
      if (!res.ok) throw new Error(await res.text())
      const data = await res.json()
      setEntries(data)
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : String(err))
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    load(currentPath)
  }, [currentPath])

  function parentPath(path: string) {
    const p = path.replace(/\/$/, '')
    const idx = p.lastIndexOf('/')
    return idx <= 0 ? '/' : p.slice(0, idx)
  }

  const breadcrumbs = currentPath.split('/').filter(Boolean)

  return (
    <main className="min-h-screen bg-gray-50 dark:bg-gray-900 text-gray-900 dark:text-gray-100">
      {/* Header */}
      <header className="border-b border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 px-6 py-4">
        <div className="max-w-6xl mx-auto flex items-center justify-between">
          <h1 className="text-xl font-bold">nasha</h1>
          <a
            href="/login"
            className="text-sm text-gray-500 hover:text-gray-700 dark:hover:text-gray-300"
            onClick={e => { e.preventDefault(); localStorage.removeItem('token'); window.location.href = '/login' }}
          >
            Sign out
          </a>
        </div>
      </header>

      {/* Breadcrumbs */}
      <div className="max-w-6xl mx-auto px-6 py-3">
        <nav className="flex items-center gap-1 text-sm text-gray-500 dark:text-gray-400">
          <a href="/files" className="hover:text-indigo-600">root</a>
          {breadcrumbs.map((crumb, i) => {
            const href = '/files/' + breadcrumbs.slice(0, i + 1).join('/')
            return (
              <span key={href}>
                <span className="mx-1">/</span>
                <a href={href} className="hover:text-indigo-600">{crumb}</a>
              </span>
            )
          })}
          {!loading && entries.length > 0 && (
            <span className="ml-auto text-xs text-gray-400">{entries.length} item{entries.length !== 1 ? 's' : ''}</span>
          )}
        </nav>
      </div>

      {/* File list */}
      <div className="max-w-6xl mx-auto px-6">
        {error && (
          <div className="rounded-lg bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 p-4 text-sm text-red-600 dark:text-red-400">
            {error}
          </div>
        )}

        {loading ? (
          <div className="text-center py-20 text-gray-400">Loading…</div>
        ) : (
          <div className="bg-white dark:bg-gray-800 rounded-xl border border-gray-200 dark:border-gray-700 overflow-hidden">
            {currentPath !== '/' && (
              <a
                href={'/files' + parentPath(currentPath)}
                className="flex items-center gap-3 px-4 py-3 text-sm text-gray-500 hover:bg-gray-50 dark:hover:bg-gray-700 border-b border-gray-100 dark:border-gray-700"
              >
                <span>📁</span>
                <span>..</span>
              </a>
            )}
            {entries.length === 0 && currentPath === '/' && (
              <div className="text-center py-20 text-gray-400">
                No files found. Configure a mount in config.yaml pointing to an existing directory.
              </div>
            )}
            {entries.map((entry) => (
              <a
                key={entry.Path}
                href={entry.IsDir ? '/files' + entry.Path : apiUrl(entry.Path)}
                target={entry.IsDir ? undefined : '_blank'}
                className="flex items-center gap-3 px-4 py-3 text-sm hover:bg-gray-50 dark:hover:bg-gray-700 border-b border-gray-100 dark:border-gray-700 last:border-0"
              >
                <span>{entry.IsDir ? '📁' : '📄'}</span>
                <span className="flex-1">{entry.Name}</span>
                {!entry.IsDir && (
                  <span className="text-xs text-gray-400">{formatSize(entry.Size)}</span>
                )}
                <span className="text-xs text-gray-400">{formatTime(entry.ModTime)}</span>
              </a>
            ))}
          </div>
        )}
      </div>
    </main>
  )
}

function formatSize(bytes: number): string {
  if (bytes === 0) return '—'
  const units = ['B', 'KB', 'MB', 'GB']
  const i = Math.min(Math.floor(Math.log(bytes) / Math.log(1024)), units.length - 1)
  return (bytes / Math.pow(1024, i)).toFixed(i === 0 ? 0 : 1) + ' ' + units[i]
}

function formatTime(iso: string): string {
  if (!iso) return ''
  const d = new Date(iso)
  return d.toLocaleDateString() + ' ' + d.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
}
