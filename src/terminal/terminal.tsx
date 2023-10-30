import ProjectLayout from '@/layouts/project-layout';
import React from 'react';
import ControlledEditor, { Monaco, useMonaco } from '@monaco-editor/react';
import { loader } from '@monaco-editor/react';

import * as monaco from 'monaco-editor';
import editorWorker from 'monaco-editor/esm/vs/editor/editor.worker?worker';
import jsonWorker from 'monaco-editor/esm/vs/language/json/json.worker?worker';
import cssWorker from 'monaco-editor/esm/vs/language/css/css.worker?worker';
import htmlWorker from 'monaco-editor/esm/vs/language/html/html.worker?worker';
import tsWorker from 'monaco-editor/esm/vs/language/typescript/ts.worker?worker';

type Props = {};

self.MonacoEnvironment = {
    getWorker(_, label) {
        if (label === 'json') {
            return new jsonWorker();
        }
        if (label === 'css' || label === 'scss' || label === 'less') {
            return new cssWorker();
        }
        if (label === 'html' || label === 'handlebars' || label === 'razor') {
            return new htmlWorker();
        }
        if (label === 'typescript' || label === 'javascript') {
            return new tsWorker();
        }
        return new editorWorker();
    },
};

loader.config({ monaco });

loader.init().then(/* ... */);


export default function Terminal({ }: Props) {
    const autocompletions = [
        { value: "SELECT", score: 1000, meta: "keyword" },
    ]

    function handleEditorWillMount(monaco: Monaco) {
        monaco.languages
    }

    function handleEditorDidMount(editor: any, monaco: Monaco) {
        editor.focus();
    }

    return (
        <ProjectLayout>
            <div className='w-full flex items-center justify-between'>
                <div className='text-4xl font-bold'>
                    Query
                </div>
            </div>
            <ControlledEditor beforeMount={handleEditorWillMount} onMount={handleEditorDidMount} className='mt-4 w-full bg-white' theme='vs-dark' height="45vh" width={"full"} defaultLanguage="sql" defaultValue="// some comment" />
        </ProjectLayout>
    )
}
