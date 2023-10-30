import Icon from '@/components/ui/icon'
import { motion } from 'framer-motion'
import GuestLayout from '@/layouts/guest-layout'

type Props = {}

export default function Home({ }: Props) {

    return (
        <GuestLayout className='flex-col justify-center items-center'>
            <>
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
            </>
        </GuestLayout>
    )
}