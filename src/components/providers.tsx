import { GithubIcon } from 'lucide-react'
import React from 'react'
import Icon from './ui/icon'

type Props = {}

export default function Providers({ }: Props) {
    return (
        <div className='flex flex-row justify-around mt-3 w-full max-w-md'>
            <div className='cursor-pointer p-3 border border-back-300 bg-back-100 rounded-md hover:bg-back-300'>
                <Icon name='google' className='w-8 h-8' />
            </div>
            <div className='cursor-pointer p-3 border border-back-300 bg-back-100 rounded-md hover:bg-back-300'>
                <GithubIcon className='w-8 h-8 stroke-white' />
            </div>
            <div className='cursor-pointer p-3 border border-back-300 bg-back-100 rounded-md hover:bg-back-300'>
                <Icon name='azure' className='w-8 h-8 stroke-white' />
            </div>
        </div>
    )
}