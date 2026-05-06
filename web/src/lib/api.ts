const BASE = '/api/v1'

function token(): string {
  return localStorage.getItem('token') || ''
}

function headers(extra: Record<string, string> = {}): Record<string, string> {
  return { Authorization: 'Bearer ' + token(), ...extra }
}

function url(path: string): string {
  return BASE + encodeURI(path)
}

async function handleResponse(res: Response) {
  if (res.status === 401) {
    localStorage.removeItem('token')
    window.location.href = '/login'
    throw new Error('unauthorized')
  }
  if (res.status === 413) {
    throw new Error('文件大小超过服务器限制')
  }
  if (!res.ok) {
    const body = await res.json().catch(() => ({ error: res.statusText }))
    throw new Error(body.error || 'request failed')
  }
  return res
}

export async function getJson(path: string) {
  const res = await fetch(url(path), { headers: headers() })
  await handleResponse(res)
  return res.json()
}

export async function postForm(path: string, form: FormData) {
  const res = await fetch(url(path), {
    method: 'POST',
    headers: headers(),
    body: form,
  })
  await handleResponse(res)
  return res.json()
}

export async function del(path: string) {
  const res = await fetch(url(path), { method: 'DELETE', headers: headers() })
  await handleResponse(res)
  return res.json()
}

export async function patchJson(path: string, body: unknown) {
  const res = await fetch(url(path), {
    method: 'PATCH',
    headers: headers({ 'Content-Type': 'application/json' }),
    body: JSON.stringify(body),
  })
  await handleResponse(res)
  return res.json()
}

export function apiUrl(path: string): string {
  return url(path) + '?token=' + encodeURIComponent(token())
}
