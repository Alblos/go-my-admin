import AddConnection from '@/components/connection-dialog/add-connection'
import ProjectLayout from '@/layouts/project-layout'

type Props = {}

export default function Connections({ }: Props) {
    return (
        <ProjectLayout>
            <div className='w-full flex items-center justify-between'>
                <div className='text-4xl font-bold'>
                    Connections
                </div>
                <AddConnection />
            </div>
            <div className='mt-8 w-full flex justify-start items-center'>
                You haven't added any connections yet. Add one to get started.
            </div>
        </ProjectLayout>
    )
}