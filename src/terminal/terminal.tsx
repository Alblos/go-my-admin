import ProjectLayout from '@/layouts/project-layout';
import React from 'react';
import ControlledEditor from '@monaco-editor/react';

type Props = {};

export default function Terminal({ }: Props) {
    const autocompletions = [
        { value: "SELECT", score: 1000, meta: "keyword" },
    ]
    return (
        <ProjectLayout>
            <ControlledEditor className='ml-12 w-full bg-white' theme='vs-dark' height="45vh" onMount={(editor, monaco) => {
                console.log(editor, monaco)
            }} width={"full"} defaultLanguage="sql" defaultValue="// some comment" />
        </ProjectLayout>
    )
}
