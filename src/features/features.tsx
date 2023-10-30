import Caroussel from '@/components/features/caroussel'
import GuestLayout from '@/layouts/guest-layout'
import React from 'react'

type Props = {}

export default function Features({ }: Props) {
    return (
        <GuestLayout className='flex-col justify-start items-center overflow-hidden'>
            <div className='flex flex-col justify-start items-center w-full h-full max-h-[calc(100vh-rem)] overflow-hidden'></div>
            <div className='text-5xl pt-24 font-bold'>Features</div>
            <Caroussel />
        </GuestLayout>
    )
}