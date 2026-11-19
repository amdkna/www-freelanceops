import { useState, type FormEvent } from 'react'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { LockKeyhole } from 'lucide-react'
import { Button } from '../components/ui/button'
import { api, ApiError } from '../lib/api/client'
export default function Login(){
 const [password,setPassword]=useState('')
 const qc=useQueryClient()
 const mutation=useMutation({mutationFn:api.login,onSuccess:async()=>{setPassword('');await qc.invalidateQueries({queryKey:['session']})}})
 const submit=(event:FormEvent<HTMLFormElement>)=>{event.preventDefault();mutation.mutate(password)}
 return <main className="mx-auto flex min-h-[70vh] max-w-md items-center px-5"><form onSubmit={submit} className="w-full space-y-5 rounded-2xl border border-slate-800 bg-slate-900 p-8"><LockKeyhole className="text-sky-400" size={30}/><div><h1 className="text-2xl font-bold">Admin sign in</h1><p className="mt-2 text-sm text-slate-400">Sign in to your private FreelanceOps workspace.</p></div><label className="block space-y-2 text-sm"><span>Password</span><input type="password" required minLength={1} autoComplete="current-password" value={password} onChange={e=>setPassword(e.target.value)} className="h-11 w-full rounded-lg border border-slate-700 bg-slate-950 px-3 text-white focus:border-sky-400 focus:outline-none"/></label>{mutation.isError&&<p role="alert" className="text-sm text-red-300">{mutation.error instanceof ApiError&&mutation.error.code==='SETUP_REQUIRED'?'Admin setup is required on the server.':mutation.error instanceof ApiError&&mutation.error.status===429?'Too many attempts. Try again later.':'Unable to sign in. Check your credentials and connection.'}</p>}<Button type="submit" disabled={mutation.isPending} className="w-full">{mutation.isPending?'Signing in…':'Sign in'}</Button></form></main>
}
