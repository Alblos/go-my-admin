import React from 'react'
import Icon from '@/components/ui/icon'
import { motion } from 'framer-motion'
import GuestNavbar from '@/components/guest-navbar'

type Props = {}

export default function Home({ }: Props) {
    const circleRef = React.useRef(null)

    React.useEffect(() => {
        window.addEventListener('mousemove', (e) => {
            if (!circleRef.current) return
            const circle = circleRef.current as HTMLDivElement
            circle.style.left = `${e.clientX - 25}px`
            circle.style.top = `${e.clientY - 25}px`
        })
    }, [])

    return (
        <>
            <div className='w-full bg-black/10 backdrop-blur-2xl h-screen fixed inset-0 -z-10'></div>
            <div className='w-full bg-transparent h-screen fixed inset-0 -z-20'>
                <div className='bg-main-200 h-36 w-36 rounded-full absolute' ref={circleRef}></div>
            </div>
            <GuestNavbar />
            <div className='flex flex-col w-full font-poppins min-h-screen bg-gradient-to-t from-back-200 via-back-300 to-25% to-back-100 justify-center items-center text-white'>
                <motion.div
                    initial={{ opacity: 0, x: -600 }}
                    animate={{ opacity: 1, x: 0 }}
                    transition={{ duration: 0.5 }}
                    className='flex items-center justify-center'>
                    <Icon name='golang' size={20} className='w-48 h-48 fill-white' />
                    <div className=' font-figtree uppercase font-black italic text-8xl tracking-tighter'>MyAdmin</div>
                </motion.div>
                <motion.div
                    initial={{ opacity: 0, x: -300 }}
                    animate={{ opacity: 1, x: 0 }}
                    transition={{ delay: 0.5, duration: 0.3 }}
                    className='max-w-2xl text-center text-2xl font-extralight'>
                    A blazingly-fast AI-powered database dashboard built on top of web3 to manage your serverless edge database deployments
                </motion.div>
            </div>
        </>
    )
}