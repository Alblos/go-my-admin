import React, { useEffect } from 'react'
import { DialogContent, DialogHeader, DialogTitle } from '../ui/dialog'
import { Button } from '../ui/button'
import { Separator } from '../ui/separator'
import Icon from '../ui/icon'
import { motion, useAnimate } from 'framer-motion'
import { useConnectionDialog } from '@/store/connection-dialog'

type Props = {}

const sources = [
    {
        name: "Railway",
        icon: <Icon name="railway" size={40} />
    },
    {
        name: "Neon",
        icon: <Icon name="neon" className="w-32 h-10" />
    },
    {
        name: "AWS",
        icon: <Icon name="aws" size={40} className="stroke-white" />
    },
    {
        name: "Azure",
        icon: <Icon name="azure" size={40} className="stroke-white" />
    }
]

export default function SelectSource({ }: Props) {
    const [scope, animate] = useAnimate()
    const { setSource } = useConnectionDialog()

    return (
        <motion.div
            initial={{
                x: -500,
            }}
            animate={{
                x: 0,
            }}
            exit={{
                x: -500
            }}
            transition={{
                duration: 0.2,
                ease: "easeInOut"
            }}
            className='w-full gap-2 flex flex-col'>
            <DialogHeader>Connect to sources</DialogHeader>
            <div className="grid grid-cols-2 gap-2 w-full">
                {
                    sources.map((source) => (
                        <Button variant={"sourceselect"} className="flex py-3 h-full min-h-fit flex-col justify-center items-center">
                            <div>{source.icon}</div>
                            <div className="text-lg mt-2 font-light">{source.name}</div>
                        </Button>
                    ))
                }
            </div>
            <Separator className='bg-back-100' />
            <Button onClick={() => setSource("manual")} variant={"sourceselect"} className="tracking-wide w-full">... or connect manually</Button>
        </motion.div>
    )
}