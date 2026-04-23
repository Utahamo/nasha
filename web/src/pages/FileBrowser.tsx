/**
 * FileBrowser page – the main file manager view.
 *
 * TODO: Implement a dual-pane Finder-style browser with:
 *   - Breadcrumb navigation
 *   - File/folder list with icons, size, and modified-date columns
 *   - Drag-and-drop upload with progress bar
 *   - Context menu (rename, delete, download, share)
 *   - Inline image / video / PDF preview panel
 */
export default function FileBrowser() {
  return (
    <main className="flex flex-col items-center justify-center min-h-screen bg-gray-50 dark:bg-gray-900 text-gray-700 dark:text-gray-300">
      <div className="text-center space-y-4 p-8">
        <h1 className="text-4xl font-bold tracking-tight text-gray-900 dark:text-white">
          File Browser
        </h1>
        <p className="text-lg text-gray-500 dark:text-gray-400">
          File browser implementation coming soon.
        </p>
      </div>
    </main>
  )
}
