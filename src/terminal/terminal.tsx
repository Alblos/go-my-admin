import ProjectLayout from '@/layouts/project-layout';
import React from 'react';
import Editor from '@monaco-editor/react';

type Props = {};

export default function Terminal({ }: Props) {
    const autocompletions = [
        { value: "SELECT", score: 1000, meta: "keyword" },
    ]
    return (
        <ProjectLayout>
            <Editor className='ml-12 w-full' theme='vs-light' height="45vh" onMount={(editor, monaco) => {
                console.log(editor, monaco)
            }} width={"full"} defaultLanguage="sql" defaultValue="// some comment" />
        </ProjectLayout>
    )
}
