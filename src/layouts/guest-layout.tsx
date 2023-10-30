import GuestNavbar from '@/components/guest-navbar'
import { AnimatePresence, motion } from 'framer-motion'
import React from 'react'
import { twMerge } from 'tailwind-merge'

type Props = {
    children: React.ReactNode;
} & React.HTMLAttributes<HTMLDivElement> & React.HTMLProps<HTMLDivElement>;

export default function GuestLayout({
    children, ...props
}: Props) {
    return (
        <>
            <GuestNavbar />
            <AnimatePresence>
                <motion.main
                    initial={{ opacity: 0 }}
                    animate={{ opacity: 1 }}
                    transition={{ duration: 0.3 }}
                    className={twMerge(props.className, `flex w-full font-poppins min-h-screen bg-gradient-to-t from-back-200 via-back-300 to-25% to-back-100 text-white`)}>
                    {children}
                </motion.main>
            </AnimatePresence>
        </>
    )
}
