import Icon from '@/components/ui/icon'
import React from 'react'
import { NavLink } from 'react-router-dom'

type Props = {
    children: React.ReactNode
}

export default function LoginLayout({ children }: Props) {
    return (
        <div className='flex flex-row overflow-x-hidden w-full h-full font-poppins min-h-screen bg-gradient-to-t from-back-200 via-back-300 to-25% to-back-100 justify-center items-center text-white'>
            <NavLink to="/" className='absolute top-2 left-6'>
                <Icon name='golang' className='w-16 h-16 fill-white' />
            </NavLink>
            <div
                className='absolute top-0 right-0 w-1/2 min-h-screen bg-back-200 border-l border-l-back-300 flex flex-col justify-center items-center'>
                {children}
            </div>
        </div>
    )
}