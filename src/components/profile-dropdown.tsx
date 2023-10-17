import React from 'react'
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { Button } from './ui/button'
import { User2Icon } from 'lucide-react'
import { useNavbarToggle } from '@/store/navbar-toggle'


type Props = {}

export default function ProfileDropdown({ }: Props) {
    const { wide } = useNavbarToggle()
    return (
        <DropdownMenu>
            <DropdownMenuTrigger>
                <Button variant={"outline"} className={`py-2 w-full ${wide ? "justify-start" : "justify-center"} gap-3 items-center`}>
                    <User2Icon size={25} />
                    {wide && <div>Profile</div>}
                </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent className={wide ? "w-full" : "absolute bottom-1 left-10 w-[150px]"}>
                <DropdownMenuLabel>Profile</DropdownMenuLabel>
                <DropdownMenuSeparator />
                <DropdownMenuItem>Account</DropdownMenuItem>
                <DropdownMenuItem>Projects</DropdownMenuItem>
                <DropdownMenuItem>Security</DropdownMenuItem>
                <DropdownMenuSeparator />
                <DropdownMenuItem className='text-red-600 font-medium'>Log Out</DropdownMenuItem>
            </DropdownMenuContent>
        </DropdownMenu>
    )
}