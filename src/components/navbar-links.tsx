import React from 'react'
import { NavLink } from 'react-router-dom'
import { Button, buttonVariants } from './ui/button'
import { VariantProps } from 'class-variance-authority'
import { useNavbarToggle } from '@/store/navbar-toggle'
import { twMerge } from 'tailwind-merge'

type Props = {
    to: string,
    title: string,
} & React.HTMLAttributes<HTMLButtonElement> & VariantProps<typeof buttonVariants>

export default function NavbarLink({ to, ...props }: Props) {
    const { wide } = useNavbarToggle()
    return (
        <NavLink to={to} className={({ isActive }) => isActive ? "[&>button]:bg-slate-950 [&>button]:hover:bg-slate-900 [&>button]:hover:text-white [&>button]:text-white" : ""} >
            <Button variant={"outline"} {...props} className={twMerge(props.className, `w-full flex items-center py-2 ${wide ? "justify-start gap-3" : "justify-center"}`)}>
                {props.children}
                {wide && <div>{props.title}</div>}
            </Button>
        </NavLink>
    )
}