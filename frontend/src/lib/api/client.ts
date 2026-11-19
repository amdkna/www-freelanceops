import type { ApiErrorBody, Check, Health, Incident, Page, Renewal, Session, Website, WebsiteInput } from './types'
export class ApiError extends Error { constructor(public readonly status: number, public readonly code: string, message: string, public readonly details?: Record<string,string>) { super(message); this.name = 'ApiError' } }
export function validateWebsiteInput(input: WebsiteInput): void {
  if (!input.name.trim() || input.name.length > 120) throw new ApiError(400, 'VALIDATION_ERROR', 'Website name must be 1–120 characters')
  let url: URL
  try { url = new URL(input.url) } catch { throw new ApiError(400, 'VALIDATION_ERROR', 'Invalid website URL') }
  if (!['http:', 'https:'].includes(url.protocol) || !url.hostname) throw new ApiError(400, 'VALIDATION_ERROR', 'Only HTTP and HTTPS URLs are supported')
  for (const date of [input.domain_expires_at, input.ssl_expires_at]) if (date != null && (!/^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(?:\.\d+)?Z$/.test(date) || Number.isNaN(Date.parse(date)))) throw new ApiError(400, 'VALIDATION_ERROR', 'Dates must be valid UTC timestamps')
}
export async function request<T>(path: string, init: RequestInit = {}, fetcher: typeof fetch = fetch): Promise<T> {
  if (!path.startsWith('/') || path.startsWith('//')) throw new Error('API path must be relative to /api')
  let response: Response
  try { response = await fetcher('/api' + path, { ...init, credentials: 'same-origin', headers: { Accept: 'application/json', ...(init.body != null ? {'Content-Type':'application/json'} : {}), ...init.headers } }) }
  catch { throw new ApiError(0, 'NETWORK_ERROR', 'Could not connect to the server') }
  const text = await response.text()
  let payload: unknown = null
  if (text) { try { payload = JSON.parse(text) } catch { throw new ApiError(response.status, 'INVALID_RESPONSE', 'Server returned invalid JSON') } }
  if (!response.ok) { const err = payload as Partial<ApiErrorBody> | null; throw new ApiError(response.status, err?.error?.code ?? 'HTTP_ERROR', err?.error?.message ?? 'Request failed', err?.error?.details) }
  return payload as T
}
const encode = encodeURIComponent
export const api = {
  health: () => request<Health>('/health/ready'),
  session: () => request<Session>('/v1/auth/session'),
  login: (password: string) => request<Session>('/v1/auth/login', {method:'POST',body:JSON.stringify({password})}),
  logout: (csrf: string) => request<void>('/v1/auth/logout',{method:'POST',headers:{'X-CSRF-Token':csrf}}),
  websites: (cursor?: string) => request<Page<Website>>('/v1/websites'+(cursor?'?cursor='+encode(cursor):'')),
  website: (id: string) => request<Website>('/v1/websites/'+encode(id)),
  createWebsite: (data: WebsiteInput, csrf: string) => { validateWebsiteInput(data); return request<Website>('/v1/websites',{method:'POST',headers:{'X-CSRF-Token':csrf},body:JSON.stringify(data)}) },
  incidents: (id: string) => request<Page<Incident>>('/v1/websites/'+encode(id)+'/incidents'),
  checks: (id: string) => request<Page<Check>>('/v1/websites/'+encode(id)+'/checks'),
  renewals: (id: string) => request<Page<Renewal>>('/v1/websites/'+encode(id)+'/renewals'),
  confirmRenewal: (id: string, expires_at: string, csrf: string) => request<Renewal>('/v1/renewals/'+encode(id)+'/confirm',{method:'POST',headers:{'X-CSRF-Token':csrf},body:JSON.stringify({expires_at})})
}
