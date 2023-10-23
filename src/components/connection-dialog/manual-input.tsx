import { Button } from '../ui/button'
import { motion } from 'framer-motion'
import { ChevronLeft } from 'lucide-react'
import { useConnectionDialog } from '@/store/connection-dialog'
import { useForm } from 'react-hook-form'
import { Label } from '../ui/label'
import { Input } from '../ui/input'
import { toast } from 'sonner'
import { Dialog, DialogContent, DialogTitle, DialogTrigger } from '../ui/dialog'
import React from 'react'

type Props = {}

export default function ManualInput({ }: Props) {
    const { setSource } = useConnectionDialog()
    const { register, handleSubmit } = useForm()

    const onSubmit = (data: any) => {
        console.log(data)
        const result = fakePromise()
        toast.promise(result, {
            loading: 'Testing connection...',
            success: 'Connection successful',
            error: 'Connection failed'
        })

        result.then(() => {
            setSource(null)
        }).catch(() => {
            console.log('error')
        })


    }

    const fakePromise = (): Promise<void> => {
        return new Promise((resolve, reject) => {
            setTimeout(() => {
                const random = Math.random();
                if (random < 0.5) {
                    reject();
                } else {
                    resolve();
                }
            }, 3000);
        });
    };


    return (
        <motion.div
            initial={{
                x: 500,
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
            <div className='w-full flex justify-between items-center'>
                <Button variant={"sourceselect"} className="tracking-wide w-fit items-center" onClick={() => setSource(null)}>
                    <ChevronLeft size={20} />
                    Back
                </Button>

            </div>
            <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-3">
                <div className=''>
                    <Label className='text-base ml-1 mb-1'>Host</Label>
                    <Input {...register("host")} />
                </div>
                <div className='flex flex-row items-center gap-3'>
                    <div className='flex-1'>
                        <Label className='text-base ml-1 mb-1'>Username</Label>
                        <Input {...register("username")} />
                    </div>
                    <div className='flex-2'>
                        <Label className='text-base ml-1 mb-1'>Port</Label>
                        <Input {...register("port")} />
                    </div>
                </div>
                <div>
                    <Label className='text-base ml-1 mb-1'>Password</Label>
                    <Input type='password' {...register("password")} />
                </div>
                <div className='flex flex-row items-center justify-center gap-3'>
                    <Button variant={"destructive"} className='w-full' onClick={() => setSource(null)}>Cancel</Button>
                    <Button variant={"sourceselect"} className='w-full'>Test Connection</Button>
                    <Button variant={"outline"} className='bg-main-200 text-black w-full'>Save</Button>
                </div>
            </form>
        </motion.div>
    )
}