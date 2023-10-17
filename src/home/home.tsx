import React from 'react'
import Golang from '../components/golang'
import { motion } from 'framer-motion'

type Props = {}

export default function Home({ }: Props) {
    return (
        <div className='flex w-full font-poppins min-h-screen bg-gradient-to-b from-cod-gray-950 to-cod-gray-900 justify-center items-center text-white'>
            <motion.div initial={{ translateY: 200, opacity: 0 }} animate={{ translateY: 0, opacity: 1 }} transition={{ duration: 2, type: "tween" }} className='flex flex-row items-center'>
                <Golang size={120} className='stroke-1 stroke-viking-200' />
                <div className='text-6xl italic text-transparent bg-gradient-to-b from-viking-100 to-viking-300 bg-clip-text'>
                    myadmin
                </div>
            </motion.div>
        </div>
    )
}