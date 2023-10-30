import { Separator } from '@radix-ui/react-select'
import { CheckCircle2Icon } from 'lucide-react'
import React from 'react'
import { Button } from '../ui/button'

type Props = {
    tier: {
        tier: string
        price: number
        yearPrice?: number
        benefits: string[]
    },
    billing: 'monthly' | 'yearly'
}

export default function PricingCard({ tier, billing }: Props) {
    return (
        <div className='w-96 rounded-md border border-back-300 p-4 h-[75vh] bg-gradient-to-t from-back-200 to-back-300 flex flex-col justify-between gap-4'>
            <div className='w-full flex flex-col gap-4'>
                <div className='text-3xl font-bold'>{tier.tier}</div>
                <div className='text-right w-full font-figtree flex justify-end gap-4'>
                    {billing == 'yearly' && tier.yearPrice && <div className='text-main-100'>{tier.yearPrice}€/month</div>}
                    <span className={billing == 'yearly' && tier.yearPrice ? 'line-through text-back-600' : ''}>{tier.price}€/month</span>
                </div>
                <Separator className='w-full bg-back-100 h-px' />
                <div className='flex flex-col gap-6 text-lg'>
                    {
                        tier.benefits.map((benefit, index) => (
                            <div key={index}>
                                <CheckCircle2Icon className='inline-block w-6 h-6 stroke-main-100 mr-2' />
                                {benefit}
                            </div>
                        ))
                    }
                </div>
            </div>
            <Button variant={"main"}>Get Started</Button>
        </div>
    )
}