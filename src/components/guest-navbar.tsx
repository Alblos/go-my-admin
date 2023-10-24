import React from 'react'
import Icon from './ui/icon'
import { NavLink } from 'react-router-dom'
import { Button } from './ui/button'

type Props = {}

export default function GuestNavbar({ }: Props) {
    return (
        <nav className='fixed items-center inset-0 w-full h-16 border-b px-12 border-b-back-300 flex justify-between'>
            <div className='flex-1'>

                <Icon name='golang' size={20} className='w-14 h-14  fill-white' />
            </div>
            <div className='flex flex-1 justify-around gap-4 items-center text-lg'>
                <NavLink to={"#"} className='py-2 w-[125px] text-center rounded-md hover:bg-back-200 font-medium hover:text-main-100'>About</NavLink>
                <NavLink to={"#"} className='py-2 w-[125px] text-center rounded-md hover:bg-back-200 font-medium hover:text-main-100'>Pricing</NavLink>
                <NavLink to={"#"} className='py-2 w-[125px] text-center rounded-md hover:bg-back-200 font-medium hover:text-main-100'>Docs</NavLink>
            </div>
            <div className='flex-1 flex justify-end gap-4'>
                <Button variant={"back"} className='px-5 w-[150px] text-base'>Log In</Button>
                <Button variant={"accent"} className='px-5 w-[150px] text-base'>Sign Up</Button>
            </div>
        </nav>
    )
}