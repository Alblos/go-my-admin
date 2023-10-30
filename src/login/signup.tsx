import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { motion } from 'framer-motion'
import React from 'react'
import { NavLink } from 'react-router-dom'

type Props = {}

export default function Signup({ }: Props) {
    return (
        <div className='flex flex-row overflow-x-hidden w-full h-full font-poppins min-h-screen bg-gradient-to-t from-back-200 via-back-300 to-25% to-back-100 justify-center items-center text-white'>
            <motion.div
                className='absolute top-0 right-0 w-1/2 min-h-screen bg-back-200 border-l border-l-back-300 flex flex-col justify-center items-center'>
                <div className='f font-jost text-4xl'>Log In</div>
                <div className='mt-3 w-full max-w-md flex justify-center items-center flex-col'>
                    <Label className='text-left w-full text-base'>Username</Label>
                    <Input placeholder='Username' />
                </div>
                <div className='mt-3 w-full max-w-md flex justify-center items-center flex-col'>
                    <Label className='text-left w-full max-w-md text-base'>Password</Label>
                    <Input placeholder='*****' name='password' type='password' className='max-w-md w-full mx-auto' />
                </div>
                <div className='w-full flex justify-between items-center max-w-md mt-3'>
                    <div className='inline-flex items-center'>
                        <input type="checkbox" name="remember" id="remember" className='accent-main-200' />
                        <label htmlFor="remember" className='ml-2 select-none'>Remember me</label>
                    </div>
                    <NavLink to='/forgot-password' className='text-sm'>
                        <Button variant={"link"} className='text-main-200 px-0'>
                            Forgot Password?
                        </Button>
                    </NavLink>
                </div>
                <Button variant={"main"} className='mt-3 w-full max-w-md font-medium py-2 text-base'>Log In</Button>
                <NavLink to={"/login"}><Button>Log In</Button></NavLink>
            </motion.div>
        </div>
    )
}