import AddConnection from "@/components/connection-dialog/add-connection"
import { Button } from "@/components/ui/button"
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger } from "@/components/ui/dropdown-menu"
import { Separator } from "@/components/ui/separator"
import ProjectLayout from "@/layouts/project-layout"
import React, { useEffect } from "react"
import { AreaChart, Card, LineChart, Title } from "@tremor/react";
import { BarList } from "@tremor/react";
import { DatabaseZapIcon } from "lucide-react"

type Props = {}

const connections = [
    {
        "id": "1",
        "name": "Connection 1",
    },
    {
        "id": "2",
        "name": "Connection 2",
    }
]

const chartValues = [
    {
        time: "2023-01-01T00:00:00Z",
        value: 0.5,
    }
]

const queries = [
    {
        name: "SELECT * FROM Users",
        value: 456,
        icon: () => <DatabaseZapIcon size={20} className="mr-3" />,
        href: "#"
    },
    {
        name: "SELECT * FROM Users WHERE id = 1",
        value: 123,
        icon: () => <DatabaseZapIcon size={20} className="mr-3" />,
        href: "#"
    },
    {
        name: "SELECT * FROM Users INNER JOIN Posts ON Users.id = Posts.user_id",
        value: 789,
        icon: () => <DatabaseZapIcon size={20} className="mr-3" />,
        href: "#"
    },
]

function generateRandomMetric() {
    const value = Math.random() * 100
    const time = new Date().toISOString()
    return {
        time,
        value,
    }
}

const valueFormatter = (value: number) => `${Math.round(value)}%`

export default function Metrics({ }: Props) {
    const [connection, setConnection] = React.useState(connections[0]?.name)
    const [usage, setUsage] = React.useState(chartValues)

    useEffect(() => {
        const interval = setInterval(() => {
            setUsage([...usage, generateRandomMetric()].slice(-10))
        }, 1000)
        return () => clearInterval(interval)
    })

    return (
        <ProjectLayout>
            <div className='w-full flex items-center justify-between'>
                <div className='text-4xl font-bold'>
                    Metrics
                </div>
            </div>
            {connections.length == 0 && <div className='mt-8 w-full flex justify-start items-center'>
                You haven't added any connections yet. Add one to get started.
            </div>}
            <div className="mt-8">
                <div className="flex flex-row w-full gap-4 justify-between items-center">
                    <DropdownMenu>
                        <DropdownMenuTrigger value={connection} className="w-full max-w-3xl">
                            <Button variant={"outline"} className="w-full justify-start">
                                {connection || 'Select a connection'}
                            </Button>
                        </DropdownMenuTrigger>
                        <DropdownMenuContent className="w-full min-w-full">
                            {
                                connections.map((connection) => {
                                    return <DropdownMenuItem className="w-[48rem]" onSelect={() => setConnection(connection.name)} key={connection.id}>{connection.name}</DropdownMenuItem>
                                })
                            }
                        </DropdownMenuContent>
                    </DropdownMenu>
                </div>
            </div>
            <Separator className="bg-back-300 my-4" />
            <div className="w-full grid grid-cols-2 gap-3">
                <Card>
                    <Title>CPU Usage</Title>
                    <AreaChart
                        className="mt-6"
                        data={usage}
                        index="time"
                        categories={["value"]}
                        colors={["yellow"]}
                        valueFormatter={valueFormatter}
                        showXAxis={false}
                        curveType="step"
                        yAxisWidth={40}
                    />
                </Card>
                <Card>
                    <Title>RAM Usage</Title>
                    <AreaChart
                        className="mt-6"
                        data={usage}
                        index="time"
                        categories={["value"]}
                        colors={["yellow"]}
                        valueFormatter={valueFormatter}
                        showXAxis={false}
                        curveType="step"
                        yAxisWidth={40}
                    />
                </Card>
            </div>
            <div className="mt-4">
                <BarList
                    data={queries.sort((a, b) => b.value - a.value)}
                    valueFormatter={(value: number) => `${value}ms`}
                    color={"yellow"}
                    showAnimation={true}
                />
            </div>
        </ProjectLayout>
    )
}