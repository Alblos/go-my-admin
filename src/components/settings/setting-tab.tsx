type Props = {
    title: string
    selected: boolean
    onClick: () => void
}

export function SettingTab({ title, selected, onClick }: Props) {
    return (
        <div onClick={onClick} className={`py-4 text-lg hover:text-main-100 w-full hover:border-r-2 hover:border-r-main-100 ${selected ? 'text-main-100 border-r-2 border-r-main-100' : ''}`}>
            {title}
        </div>
    )
}