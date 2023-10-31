import Providers from '@/components/providers'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Separator } from '@/components/ui/separator'
import LoginLayout from '@/layouts/login-layout'
import { NavLink } from 'react-router-dom'

type Props = {}

export default function Login({ }: Props) {
    return (
        <LoginLayout>                <div className='f font-jost text-4xl'>Log In</div>
            <Providers />
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
            <Separator className='bg-back-300 max-w-md w-full my-4' />
            <NavLink to={"/signup"}><Button variant={"link"}>Sign Up</Button></NavLink>
        </LoginLayout>
    )
}