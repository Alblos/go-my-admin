import { motion } from 'framer-motion'
import React from 'react'
import { twMerge } from 'tailwind-merge'

type Props = {
    value: string
    onChange: (value: string) => void | React.Dispatch<React.SetStateAction<string>>
} & React.HTMLAttributes<HTMLDivElement>

export default function Dropdown({ ...props }: Props) {
    const [open, setOpen] = React.useState(false)
    return (
        <div onClick={() => setOpen(o => !o)} role='select' {...props} className={twMerge('w-full cursor-pointer relative flex flex-row justify-between px-2 border border-gray-200 hover:bg-gray-100 rounded-lg py-1 items-center', props.className)}>
            {props.children}
            {open && <motion.div className='absolute w-full top-8 left-0 h-fit bg-white border border-gray-200 flex flex-col p-2'>
                <div className='py-1 w-full px-2 rounded-lg hover:bg-gray-100'>Project 1</div>
                <div className='py-1 w-full px-2 rounded-lg hover:bg-gray-100'>Project 2</div>
                <div className='py-1 w-full px-2 rounded-lg hover:bg-gray-100'>Project 3</div>
            </motion.div>}
        </div>
    )
}