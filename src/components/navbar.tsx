import React, { useRef } from 'react'
import { BarChart3Icon, Database, DatabaseIcon, PanelLeftCloseIcon, SettingsIcon, TablePropertiesIcon, TerminalSquareIcon } from 'lucide-react'
import { Button } from './ui/button'
import { animate, motion, useAnimate } from 'framer-motion'
import { useNavbarToggle } from '@/store/navbar-toggle'
import SelectProject from './select-project'
import NavbarLink from './navbar-links'
import ProfileDropdown from './profile-dropdown'
import { useLocation } from 'react-router-dom'

type Props = {}

const navigationItems = [
    {
        "to": "/dashboard",
        "title": "Connections",
        "icon": <DatabaseIcon size={25} />
    },
    {
        "to": "/metrics",
        "title": "Metrics",
        "icon": <BarChart3Icon size={25} />
    },
    {
        "to": "/tables",
        "title": "Tables",
        "icon": <TablePropertiesIcon size={25} />
    },
    {
        "to": "/terminal",
        "title": "Query",
        "icon": <TerminalSquareIcon size={25} />
    },
    {
        "to": "/settings",
        "title": "Settings",
        "icon": <SettingsIcon size={25} />
    }
]

export default function Navbar({ }: Props) {
    const { wide, toggleWide } = useNavbarToggle()
    const navRef = useRef<HTMLDivElement>(null)
    const navToggle = useRef<HTMLButtonElement>(null)
    const toggleNavbar = () => {
        if (navRef.current) {
            if (wide) {
                animate(navRef.current, {
                    width: 100,
                }).then(() => {
                    toggleWide()
                })
            } else {
                animate(navRef.current, {
                    width: 320,
                }).then(() => {
                    toggleWide()
                })
            }
        }
        if (navToggle.current) {
            animate(navToggle.current, {
                rotate: wide ? 180 : 0
            })
        }
    }

    const route = useLocation()
    // const [scope, animate] = useAnimate()


    // React.useEffect(() => {
    //     const currentRouteIndex = navigationItems.findIndex((item) => item.to === route.pathname)
    //     console.log(scope.current.style)
    //     const start = isNaN(parseFloat(scope.current.style.top.replace("px", ""))) ? 0 : parseFloat(scope.current.style.top.replace("px", ""))
    //     const end = 12 + (currentRouteIndex * 50)
    //     console.log(start, end)
    //     animate(scope.current, {
    //         top: [start, end],
    //     })
    // }, [route])

    return (
        <nav ref={navRef} className={`fixed inset-0 z-30 ${wide ? "w-[320px] px-4" : "w-[100px] px-2"} py-4 flex flex-col justify-between gap-3 border-r border-r-back-300 shadow-lg`}>
            <div className='flex flex-col h-full gap-3'>
                <div className={`flex ${wide ? 'flex-row items-center justify-end' : "flex-col-reverse"} gap-2`}>
                    <SelectProject />
                    <Button variant={"outline"} size={wide ? "icon" : "default"} ref={navToggle} onClick={toggleNavbar}><PanelLeftCloseIcon size={25} className={wide ? "" : "w-full"} /></Button>
                </div>
                {/* <div ref={scope} className='absolute top-12 right-2 bg-main-200 h-10 w-1'></div> */}
                {
                    navigationItems.map((item, index) => {
                        return (
                            <NavbarLink key={index} to={item.to} title={item.title} className={wide ? "" : "w-full"}>{item.icon}</NavbarLink>
                        )
                    })
                }
            </div>
            <ProfileDropdown />
        </nav>
    )
}