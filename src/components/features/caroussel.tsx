import { AnimatePresence, motion } from 'framer-motion';
import React from 'react'
import { Button } from '../ui/button';
import { ChevronDown } from 'lucide-react';

type Props = {}

const items = [
    {
        id: 1,
        title: 'Manage Your Connections',
        subtitle: "Manage multiple connections to your databases",
        src: "/undraw_manage.svg"
    },
    {
        id: 2,
        title: 'Work Together with your team',
        subtitle: "Manage all your projects in one place with your team",
        src: "/undraw_team.svg"
    },
    {
        id: 3,
        title: 'Create a complete, fast and secure backend in a few clicks',
        subtitle: "Use our auto CRUD generator to create a complete backend in a few clicks",
        src: "/undraw_connect.svg"
    },
    {
        id: 4,
        title: 'Create custom endpoints',
        subtitle: "Create custom endpoints to handle specific tasks",
        src: "/undraw_server.svg"
    },
    {
        id: 5,
        title: 'Use our GraphQL API',
        subtitle: "Use our GraphQL API to handle all your data",
        src: "/GraphQL_Logo.svg"
    }
    // Add more items as needed
];

export default function Caroussel({ }: Props) {
    const [currentIndex, setCurrentIndex] = React.useState(0);

    const nextItem = () => {
        setCurrentIndex((prevIndex) => (prevIndex < items.length - 1 ? prevIndex + 1 : 0));
    };

    // React.useEffect(() => {
    //     const timer = setTimeout(() => {
    //         nextItem()
    //     }, 5000)
    //     return () => clearTimeout(timer)
    // }, [currentIndex])

    return (
        <div className="relative h-[calc(100vh-10rem)] flex justify-between w-full items-center">
            <AnimatePresence initial={true} mode='wait'>
                <motion.div
                    key={currentIndex}
                    initial={{ opacity: 0, y: 100 }}
                    animate={{ opacity: 1, y: 0 }}
                    exit={{ opacity: 0, y: -100 }}
                    transition={{ type: 'tween', duration: 0.5 }}
                    className=" w-full max-w-6xl mx-auto text-2xl flex justify-around items-center px-8"
                >
                    <div className='w-full pr-12'>
                        <div className='text-2xl text-left'>
                            {items[currentIndex].title}
                        </div>
                        <div className='text-lg'>{items[currentIndex].subtitle}</div>
                    </div>
                    <div className='relative min-w-fit w-[400px] h-[400px] object-contain bg-transparent'>
                        <img src={items[currentIndex].src} alt={items[currentIndex].title} className='rounded-md w-[400px] h-[400px] object-center object-contain' />
                    </div>
                </motion.div>
            </AnimatePresence>
            <div className='absolute bottom-0 left-0 w-full flex justify-center items-center'>
                <div className='flex flex-col justify-center items-center gap-2'>
                    <div className='flex flex-row justify-center w-fit items-center gap-1'>
                        {
                            Array.from({ length: currentIndex + 1 }, (_, i) => (
                                <div key={i} className='bg-main-100 w-3 h-3 rounded-full'></div>
                            ))
                        }
                        {
                            Array.from({ length: items.length - currentIndex - 1 }, (_, i) => (
                                <div key={i} className='bg-back-100 w-3 h-3 rounded-full'></div>
                            ))
                        }
                    </div>
                    <Button variant='secondary' className='bg-back-100' onClick={nextItem}>
                        <ChevronDown className='w-6 h-6 stroke-main-100' />
                    </Button>
                </div>
            </div>
        </div>
    );
}