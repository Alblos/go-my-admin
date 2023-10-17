import React from 'react'
import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectLabel,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select"
import { useProjectSelect } from '@/store/project-select'
import { useNavbarToggle } from '@/store/navbar-toggle'
type Props = {}

export default function SelectProject({ }: Props) {
    const { project, setProject } = useProjectSelect()
    const { wide } = useNavbarToggle()
    return (
        <Select value={project} onValueChange={setProject}>
            <SelectTrigger>
                <SelectValue>{wide ? project : project.slice(0, 2).toUpperCase()}</SelectValue>
            </SelectTrigger>
            <SelectContent className={wide ? "" : "absolute left-[90px] min-w-[200px] -top-10"}>
                <SelectGroup>
                    <SelectItem value="Project 1">
                        <SelectLabel>Project 1</SelectLabel>
                    </SelectItem>
                    <SelectItem value="Project 2">
                        <SelectLabel>Project 2</SelectLabel>
                    </SelectItem>
                    <SelectItem value="Project 3">
                        <SelectLabel>Project 3</SelectLabel>
                    </SelectItem>
                </SelectGroup>
            </SelectContent>
        </Select>
    )
}