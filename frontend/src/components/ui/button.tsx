import * as React from 'react'
import { Slot } from '@radix-ui/react-slot'
import { cva, type VariantProps } from 'class-variance-authority'
import { cn } from '../../lib/utils'
const buttonVariants = cva('inline-flex items-center justify-center gap-2 rounded-lg text-sm font-medium transition-colors focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-sky-400 disabled:pointer-events-none disabled:opacity-50', { variants: { variant: { default: 'bg-sky-500 px-4 py-2 text-slate-950 hover:bg-sky-400', outline: 'border border-slate-600 px-4 py-2 text-slate-100 hover:bg-slate-800', ghost: 'px-4 py-2 text-slate-300 hover:bg-slate-800' }, size: { default: 'h-10', sm: 'h-9', lg: 'h-12' } }, defaultVariants: { variant: 'default', size: 'default' } })
type Props = React.ComponentProps<'button'> & VariantProps<typeof buttonVariants> & { asChild?: boolean }
export function Button({ asChild = false, className, variant, size, ...props }: Props) { const Comp = asChild ? Slot : 'button'; return <Comp className={cn(buttonVariants({variant,size,className}))} {...props} /> }
