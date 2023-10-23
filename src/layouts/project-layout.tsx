import Navbar from '@/components/navbar'
import { useNavbarToggle } from '@/store/navbar-toggle'
import { AnimatePresence, motion } from 'framer-motion'
import React from 'react'
import { Toaster } from 'sonner'


type Props = {
    children: React.ReactNode
}

export default function ProjectLayout({ children }: Props) {
    const { wide } = useNavbarToggle()
    return (
        <div

            className='w-full h-full flex flex-row justify-start items-start'>
            <Navbar />
            <motion.main
                initial={{
                    opacity: 0,
                }}
                animate={{
                    opacity: 1,
                }}
                transition={{
                    duration: 0.3,
                }
                }
                className={`absolute top-5 md:px-8 px-4 ${wide ? "left-[320px] w-[calc(100%-320px)]" : "left-[100px] w-[calc(100%-100px)]"}`}>
                <AnimatePresence>
                    {children}
                </AnimatePresence>
            </motion.main>
            <Toaster richColors expand />
        </div>
    )
}