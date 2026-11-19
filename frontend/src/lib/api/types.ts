/** Mirrors docs/openapi.yaml. All timestamps are ISO-8601 UTC strings. */
export type UtcDateTime = string
export type WebsiteStatus = 'up' | 'down' | 'unknown'
export interface Website { id: string; name: string; url: string; status: WebsiteStatus; domain_expires_at: UtcDateTime | null; ssl_expires_at: UtcDateTime | null; created_at: UtcDateTime }
export interface WebsiteInput { name: string; url: string; domain_expires_at?: UtcDateTime | null; ssl_expires_at?: UtcDateTime | null }
export interface Incident { id: string; website_id: string; started_at: UtcDateTime; resolved_at: UtcDateTime | null }
export interface Check { id: string; website_id: string; checked_at: UtcDateTime; status_code: number | null; success: boolean }
export interface Renewal { id: string; website_id: string; kind: 'domain' | 'ssl'; expires_at: UtcDateTime; confirmed_at: UtcDateTime | null }
export interface ApiErrorBody { error: { code: string; message: string; details?: Record<string, string> }; request_id?: string }
export interface Session { authenticated: boolean; csrf_token: string }
export interface Health { status: 'ok' | 'not_ready' }
export interface Page<T> { items: T[]; next_cursor: string | null }
