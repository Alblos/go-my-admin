import ProjectLayout from '@/layouts/project-layout'
import React from 'react'

type Props = {}

function SettingTab({ title, selected, onClick }: { title: string, selected: boolean, onClick: () => void }) {
    return (
        <div onClick={onClick} className={`py-4 text-lg hover:text-main-100 w-full hover:border-r-2 hover:border-r-main-100 ${selected ? 'text-main-100 border-r-2 border-r-main-100' : ''}`}>
            {title}
        </div>
    )
}

const items = [
    'Account',
    'Suscription',
    'Billing',
    'Security',
    'Team',
    'Notifications',
    'Integrations',
]

export default function Settings({ }: Props) {
    const [selected, setSelected] = React.useState(items[0])
    return (
        <ProjectLayout>
            <div className='w-full flex items-center justify-between'>
                <div className='text-4xl font-bold'>
                    Settings
                </div>
            </div>
            <div className='w-full flex flex-row mt-4'>
                <div className='w-1/5 h-full flex flex-col border-r border-r-back-300'>
                    {
                        items.map((item, index) => (
                            <SettingTab key={index} title={item} selected={selected === item} onClick={() => setSelected(item)} />
                        ))
                    }
                </div>
                <div className='w-4/5 h-full px-8'>
                    <div className='text-3xl'>{selected}</div>
                </div>
            </div>
        </ProjectLayout>
    )
}