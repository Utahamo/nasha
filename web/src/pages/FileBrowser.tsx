import { useEffect, useState, useRef } from 'react'
import { useParams } from 'react-router-dom'
import { Folder, File, Upload, Plus, Trash2, Pencil, LogOut } from 'lucide-react'
import { getJson, apiUrl, postForm, del, patchJson } from '../lib/api'

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
  const [renaming, setRenaming] = useState<string | null>(null)
  const [renameValue, setRenameValue] = useState('')
  const [showNewFolder, setShowNewFolder] = useState(false)
  const [newFolderName, setNewFolderName] = useState('')
  const fileInputRef = useRef<HTMLInputElement>(null)

  async function load(path: string) {
    setLoading(true)
    setError('')
    try {
      const data = await getJson('/fs' + path)
      setEntries(data)
    } catch (err: unknown) {
      if (err instanceof Error) setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => { load(currentPath) }, [currentPath])

  function parentPath(path: string) {
    const p = path.replace(/\/$/, '')
    const idx = p.lastIndexOf('/')
    return idx <= 0 ? '/' : p.slice(0, idx)
  }

  // Upload
  async function handleUpload(e: React.ChangeEvent<HTMLInputElement>) {
    const files = e.target.files
    if (!files?.length) return
    const form = new FormData()
    for (const f of files) form.append('file', f)
    try {
      await postForm('/fs' + currentPath, form)
      load(currentPath)
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : 'upload failed')
    }
    if (fileInputRef.current) fileInputRef.current.value = ''
  }

  // New folder
  async function handleNewFolder() {
    if (!newFolderName.trim()) return
    try {
      await postForm('/fs' + currentPath + '/' + newFolderName.trim() + '?mkdir=1', new FormData())
      setShowNewFolder(false)
      setNewFolderName('')
      load(currentPath)
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : 'create folder failed')
    }
  }

  // Delete
  async function handleDelete(entry: FileEntry) {
    if (!window.confirm(`确认删除「${entry.Name}」？`)) return
    try {
      await del('/fs' + entry.Path)
      load(currentPath)
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : 'delete failed')
    }
  }

  // Rename
  async function handleRenameSubmit(entry: FileEntry) {
    if (!renameValue.trim()) return
    const parent = parentPath(entry.Path)
    const dst = parent === '/' ? '/' + renameValue.trim() : parent + '/' + renameValue.trim()
    try {
      await patchJson('/fs' + entry.Path, { src: entry.Path, dst })
      setRenaming(null)
      load(currentPath)
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : 'rename failed')
    }
  }

  const breadcrumbs = currentPath.split('/').filter(Boolean)

  return (
    <main className="min-h-screen bg-gray-50 dark:bg-gray-900 text-gray-900 dark:text-gray-100">
      {/* Header */}
      <header className="border-b border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 px-6 py-4">
        <div className="max-w-6xl mx-auto flex items-center justify-between">
          <h1 className="text-xl font-bold">nasha</h1>
          <div className="flex items-center gap-3">
            {/* Upload button */}
            <button
              onClick={() => fileInputRef.current?.click()}
              className="flex items-center gap-1.5 text-sm text-gray-500 hover:text-gray-700 dark:hover:text-gray-300"
              title="上传文件"
            >
              <Upload className="w-4 h-4" />
              <span className="hidden sm:inline">上传</span>
            </button>
            <input
              ref={fileInputRef}
              type="file"
              multiple
              className="hidden"
              onChange={handleUpload}
            />
            {/* New folder button */}
            <button
              onClick={() => { setShowNewFolder(true); setNewFolderName('') }}
              className="flex items-center gap-1.5 text-sm text-gray-500 hover:text-gray-700 dark:hover:text-gray-300"
              title="新建文件夹"
            >
              <Plus className="w-4 h-4" />
              <span className="hidden sm:inline">新建</span>
            </button>
            {/* Sign out */}
            <button
              onClick={() => { localStorage.removeItem('token'); window.location.href = '/login' }}
              className="flex items-center gap-1.5 text-sm text-gray-500 hover:text-gray-700 dark:hover:text-gray-300"
              title="退出登录"
            >
              <LogOut className="w-4 h-4" />
            </button>
          </div>
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
            <span className="ml-auto text-xs text-gray-400">{entries.length} 项</span>
          )}
        </nav>
      </div>

      {/* File list */}
      <div className="max-w-6xl mx-auto px-6">
        {error && (
          <div className="mb-4 rounded-lg bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 p-4 text-sm text-red-600 dark:text-red-400">
            {error}
            <button onClick={() => setError('')} className="ml-2 font-bold">&times;</button>
          </div>
        )}

        {/* New folder inline form */}
        {showNewFolder && (
          <div className="mb-4 flex items-center gap-2">
            <input
              autoFocus
              value={newFolderName}
              onChange={e => setNewFolderName(e.target.value)}
              onKeyDown={e => e.key === 'Enter' && handleNewFolder()}
              placeholder="文件夹名称"
              className="rounded-lg border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-800 px-3 py-1.5 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500"
            />
            <button onClick={handleNewFolder} className="text-sm px-3 py-1.5 rounded-lg bg-indigo-600 text-white hover:bg-indigo-700">确定</button>
            <button onClick={() => setShowNewFolder(false)} className="text-sm px-3 py-1.5 rounded-lg border border-gray-300 dark:border-gray-600 hover:bg-gray-100 dark:hover:bg-gray-700">取消</button>
          </div>
        )}

        {loading ? (
          <div className="text-center py-20 text-gray-400">加载中…</div>
        ) : (
          <div className="bg-white dark:bg-gray-800 rounded-xl border border-gray-200 dark:border-gray-700 overflow-hidden">
            {currentPath !== '/' && (
              <a
                href={'/files' + parentPath(currentPath)}
                className="flex items-center gap-3 px-4 py-3 text-sm text-gray-500 hover:bg-gray-50 dark:hover:bg-gray-700 border-b border-gray-100 dark:border-gray-700"
              >
                <Folder className="w-4 h-4 text-blue-400" />
                <span>..</span>
              </a>
            )}
            {entries.length === 0 && currentPath === '/' && (
              <div className="text-center py-20 text-gray-400">
                暂无文件。请在 config.yaml 中配置一个指向已有目录的挂载点。
              </div>
            )}
            {entries.map((entry) => (
              <div
                key={entry.Path}
                className="group flex items-center gap-3 px-4 py-3 text-sm hover:bg-gray-50 dark:hover:bg-gray-700 border-b border-gray-100 dark:border-gray-700 last:border-0"
              >
                {/* Icon + name */}
                {entry.IsDir ? (
                  <a href={'/files' + entry.Path} className="flex items-center gap-3 flex-1 min-w-0">
                    <Folder className="w-4 h-4 text-blue-500 shrink-0" />
                    <span className="truncate">{entry.Name}</span>
                  </a>
                ) : (
                  <a href={apiUrl(entry.Path)} target="_blank" className="flex items-center gap-3 flex-1 min-w-0">
                    <File className="w-4 h-4 text-gray-400 shrink-0" />
                    <span className="truncate">{entry.Name}</span>
                  </a>
                )}

                {/* Size & time */}
                {!entry.IsDir && (
                  <span className="text-xs text-gray-400 shrink-0 w-16 text-right">{formatSize(entry.Size)}</span>
                )}
                <span className="text-xs text-gray-400 shrink-0 w-32 text-right hidden sm:block">{formatTime(entry.ModTime)}</span>

                {/* Actions */}
                <div className="flex items-center gap-1 opacity-0 group-hover:opacity-100 transition-opacity shrink-0">
                  {!entry.IsDir && entry.Name !== '..' && (
                    <a href={apiUrl(entry.Path)} target="_blank" title="下载" className="p-1 text-gray-400 hover:text-indigo-600">
                      <Upload className="w-3.5 h-3.5" />
                    </a>
                  )}
                  {entry.Name !== '..' && (
                    <>
                      <button
                        onClick={() => { setRenaming(entry.Path); setRenameValue(entry.Name) }}
                        title="重命名"
                        className="p-1 text-gray-400 hover:text-indigo-600"
                      >
                        <Pencil className="w-3.5 h-3.5" />
                      </button>
                      <button
                        onClick={() => handleDelete(entry)}
                        title="删除"
                        className="p-1 text-gray-400 hover:text-red-500"
                      >
                        <Trash2 className="w-3.5 h-3.5" />
                      </button>
                    </>
                  )}
                </div>
              </div>
            ))}
          </div>
        )}
      </div>

      {/* Rename dialog */}
      {renaming && (
        <div className="fixed inset-0 bg-black/40 flex items-center justify-center z-50" onClick={() => setRenaming(null)}>
          <div className="bg-white dark:bg-gray-800 rounded-xl p-6 w-full max-w-sm mx-4 shadow-xl" onClick={e => e.stopPropagation()}>
            <h3 className="text-sm font-medium mb-3">重命名</h3>
            <input
              autoFocus
              value={renameValue}
              onChange={e => setRenameValue(e.target.value)}
              onKeyDown={e => {
                if (e.key === 'Enter') {
                  const entry = entries.find(en => en.Path === renaming)
                  if (entry) handleRenameSubmit(entry)
                }
                if (e.key === 'Escape') setRenaming(null)
              }}
              className="w-full rounded-lg border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 mb-3"
            />
            <div className="flex justify-end gap-2">
              <button onClick={() => setRenaming(null)} className="px-3 py-1.5 text-sm rounded-lg border border-gray-300 dark:border-gray-600 hover:bg-gray-100 dark:hover:bg-gray-700">取消</button>
              <button
                onClick={() => {
                  const entry = entries.find(en => en.Path === renaming)
                  if (entry) handleRenameSubmit(entry)
                }}
                className="px-3 py-1.5 text-sm rounded-lg bg-indigo-600 text-white hover:bg-indigo-700"
              >
                确定
              </button>
            </div>
          </div>
        </div>
      )}
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
