import PricingCard from '@/components/pricing/pricing-card'
import { Switch } from '@/components/ui/switch'
import GuestLayout from '@/layouts/guest-layout'
import React from 'react'

type Props = {}

const tiers = [
    {
        tier: "Free",
        price: 0.00,
        benefits: ["2 Active Projects", "1 Active Teams", "2 Active Connections", "GraphQL API", "Auto CRUD Generator"],
    },
    {
        tier: "Basic",
        price: 10.99,
        yearPrice: 7.99,
        benefits: ["5 Active Projects", "3 Active Teams", "Unlimited Connections", "GraphQL API", "Auto CRUD Generator", "5 Custom Endpoints"],
    },
    {
        tier: "Pro",
        price: 25.99,
        yearPrice: 20.99,
        benefits: ["Unlimited Projects", "Unlimited Teams", "Unlimited Connections", "GraphQL API", "Auto CRUD Generator", "Unlimited Custom Endpoints"]
    }
]

export default function Pricing({ }: Props) {
    const [billing, setBilling] = React.useState<'monthly' | 'yearly'>('monthly')
    return (
        <GuestLayout className='flex-col justify-start items-center'>
            <div className='text-5xl pt-24 font-bold'>Pricing</div>
            <div className='flex flex-row mt-4 items-center gap-2 text-lg'>
                <div className={billing == 'monthly' ? 'text-main-200' : 'text-back-300'}>Monthly</div>
                <Switch value={billing == 'monthly' ? 'on' : 'off'} onClick={(e) => { console.log(e); setBilling(billing == 'monthly' ? 'yearly' : 'monthly') }} />
                <div className={billing == 'yearly' ? 'text-main-200' : 'text-back-300'}>Yearly</div>
            </div>
            <div className='w-full mt-4 flex justify-center items-start gap-6'>
                {
                    tiers.map((tier, index) => (
                        <PricingCard key={index} tier={tier} billing={billing} />
                    ))
                }
            </div>
        </GuestLayout>
    )
}