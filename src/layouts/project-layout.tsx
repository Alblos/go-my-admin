import Navbar from '@/components/navbar'
import { useNavbarToggle } from '@/store/navbar-toggle'
import React from 'react'

type Props = {
    children: React.ReactNode
}

export default function ProjectLayout({ children }: Props) {
    const { wide } = useNavbarToggle()
    return (
        <div className='w-full h-full flex flex-row justify-start items-start bg-white'>
            <Navbar />
            <main className={`absolute top-3 px-4 ${wide ? "left-[320px] w-[calc(100%-320px)]" : "left-[100px] w-[calc(100%-100px)]"}`}>
                {children}
            </main>
        </div>
    )
}