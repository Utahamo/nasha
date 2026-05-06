import { Link } from 'react-router-dom'

export default function NotFound() {
  return (
    <main className="flex flex-col items-center justify-center min-h-screen bg-gray-50 dark:bg-gray-900 text-gray-700 dark:text-gray-300">
      <div className="text-center space-y-4 p-8">
        <h1 className="text-6xl font-bold text-gray-900 dark:text-white">404</h1>
        <p className="text-lg text-gray-500 dark:text-gray-400">页面未找到</p>
        <Link to="/files" className="inline-block mt-4 text-sm text-indigo-600 hover:text-indigo-700">
          返回文件浏览器
        </Link>
      </div>
    </main>
  )
}
