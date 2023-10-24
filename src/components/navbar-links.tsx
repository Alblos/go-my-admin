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
        <NavLink to={to} className={({ isActive }) => isActive ? "[&>button]:text-main-200 [&>button]:border-r-main-200 [&>button]:border-r-4" : ""} >
            <Button {...props} className={twMerge(props.className, `w-full bg-transparent hover:text-main-200 rounded-none flex items-center py-2 ${wide ? "justify-start gap-3" : "justify-center"}`)}>
                {props.children}
                {wide && <div>{props.title}</div>}
            </Button>
        </NavLink>
    )
}