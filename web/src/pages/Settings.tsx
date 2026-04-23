/**
 * Settings / Admin page.
 *
 * TODO: Implement the admin panel with:
 *   - Mount point management (add / edit / remove storage backends)
 *   - User management (create, assign roles)
 *   - Per-mount permission overrides
 */
export default function Settings() {
  return (
    <main className="flex flex-col items-center justify-center min-h-screen bg-gray-50 dark:bg-gray-900 text-gray-700 dark:text-gray-300">
      <div className="text-center space-y-4 p-8">
        <h1 className="text-4xl font-bold tracking-tight text-gray-900 dark:text-white">
          Settings
        </h1>
        <p className="text-lg text-gray-500 dark:text-gray-400">
          Admin panel coming soon.
        </p>
      </div>
    </main>
  )
}
