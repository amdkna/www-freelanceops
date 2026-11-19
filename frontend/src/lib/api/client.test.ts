import { describe, expect, it, vi } from 'vitest'
import { ApiError, request, validateWebsiteInput } from './client'
const reply = (status: number, body: unknown) => vi.fn(async () => new Response(JSON.stringify(body), {status, headers:{'Content-Type':'application/json'}})) as unknown as typeof fetch
describe('typed API client', () => {
  it('parses successful 200 responses', async () => { const fetcher=reply(200,{status:'ok'}); await expect(request('/health/ready',{},fetcher)).resolves.toEqual({status:'ok'}); expect(fetcher).toHaveBeenCalledWith('/api/health/ready', expect.objectContaining({credentials:'same-origin'})) })
  it.each([400,401,404])('preserves %i and structured errors',async status => { await expect(request('/v1/websites',{},reply(status,{error:{code:'TEST_ERROR',message:'Denied'}}))).rejects.toMatchObject({status,code:'TEST_ERROR',message:'Denied'}) })
  it('handles network failure without leaking details',async () => { await expect(request('/health/ready',{},vi.fn(async()=>{throw Error('private host')}) as unknown as typeof fetch)).rejects.toMatchObject({status:0,code:'NETWORK_ERROR'}) })
  it('rejects malformed response',async()=> { await expect(request('/health/ready',{},vi.fn(async()=>new Response('<html>')) as unknown as typeof fetch)).rejects.toMatchObject({code:'INVALID_RESPONSE'}) })
  it('checks website input and UTC dates',()=> { expect(()=>validateWebsiteInput({name:'',url:'https://example.com'})).toThrow(ApiError); expect(()=>validateWebsiteInput({name:'ok',url:'file:///etc/passwd'})).toThrow(ApiError); expect(()=>validateWebsiteInput({name:'ok',url:'https://example.com',ssl_expires_at:'2026-12-20T09:00:00+04:00'})).toThrow(ApiError); expect(()=>validateWebsiteInput({name:'ok',url:'https://example.com',ssl_expires_at:'2026-12-20T09:00:00Z'})).not.toThrow() })
})
