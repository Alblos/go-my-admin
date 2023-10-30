import React from 'react'

type Props = {
    name: string
    size?: number
} & React.SVGProps<SVGSVGElement>

export default function Icon({ name, size = 25, ...props }: Props) {
    switch (name.toLowerCase()) {
        case 'railway':
            return <svg {...props} xmlns="http://www.w3.org/2000/svg" width={size} height={size} viewBox="0 0 1024 1024" fill="none"><path d="M4.756 438.175A520.713 520.713 0 0 0 0 489.735h777.799c-2.716-5.306-6.365-10.09-10.045-14.772-132.97-171.791-204.498-156.896-306.819-161.26-34.114-1.403-57.249-1.967-193.037-1.967-72.677 0-151.688.185-228.628.39-9.96 26.884-19.566 52.942-24.243 74.14h398.571v51.909H4.756ZM783.93 541.696H.399c.82 13.851 2.112 27.517 3.978 40.999h723.39c32.248 0 50.299-18.297 56.162-40.999ZM45.017 724.306S164.941 1018.77 511.46 1024c207.112 0 385.071-123.006 465.907-299.694H45.017Z" fill="#fff" /><path d="M511.454 0C319.953 0 153.311 105.16 65.31 260.612c68.771-.144 202.704-.226 202.704-.226h.031v-.051c158.309 0 164.193.707 195.118 1.998l19.149.706c66.7 2.224 148.683 9.384 213.19 58.19 35.015 26.471 85.571 84.896 115.708 126.52 27.861 38.499 35.876 82.756 16.933 125.158-17.436 38.97-54.952 62.215-100.383 62.215H16.69s4.233 17.944 10.58 37.751h970.632A510.385 510.385 0 0 0 1024 512.218C1024.01 229.355 794.532 0 511.454 0Z" fill="#fff" /></svg>
        case 'neon':
            return (
                <svg {...props} xmlns="http://www.w3.org/2000/svg" width={size * 3.55} height={size} fill="none">
                    <path fill="#12FFF7" fill-rule="evenodd" d="M0 6.207A6.207 6.207 0 0 1 6.207 0h23.586A6.207 6.207 0 0 1 36 6.207v20.06c0 3.546-4.488 5.085-6.664 2.286l-6.805-8.754v10.615A5.586 5.586 0 0 1 16.945 36H6.207A6.207 6.207 0 0 1 0 29.793V6.207Zm6.207-1.241c-.686 0-1.241.555-1.241 1.24v23.587c0 .686.555 1.242 1.24 1.242h10.925c.343 0 .434-.278.434-.621V16.18c0-3.547 4.488-5.086 6.665-2.286l6.805 8.753V6.207c0-.686.064-1.241-.621-1.241H6.207Z" clip-rule="evenodd" />
                    <path fill="url(#a)" fill-rule="evenodd" d="M0 6.207A6.207 6.207 0 0 1 6.207 0h23.586A6.207 6.207 0 0 1 36 6.207v20.06c0 3.546-4.488 5.085-6.664 2.286l-6.805-8.754v10.615A5.586 5.586 0 0 1 16.945 36H6.207A6.207 6.207 0 0 1 0 29.793V6.207Zm6.207-1.241c-.686 0-1.241.555-1.241 1.24v23.587c0 .686.555 1.242 1.24 1.242h10.925c.343 0 .434-.278.434-.621V16.18c0-3.547 4.488-5.086 6.665-2.286l6.805 8.753V6.207c0-.686.064-1.241-.621-1.241H6.207Z" clip-rule="evenodd" />
                    <path fill="url(#b)" fill-rule="evenodd" d="M0 6.207A6.207 6.207 0 0 1 6.207 0h23.586A6.207 6.207 0 0 1 36 6.207v20.06c0 3.546-4.488 5.085-6.664 2.286l-6.805-8.754v10.615A5.586 5.586 0 0 1 16.945 36H6.207A6.207 6.207 0 0 1 0 29.793V6.207Zm6.207-1.241c-.686 0-1.241.555-1.241 1.24v23.587c0 .686.555 1.242 1.24 1.242h10.925c.343 0 .434-.278.434-.621V16.18c0-3.547 4.488-5.086 6.665-2.286l6.805 8.753V6.207c0-.686.064-1.241-.621-1.241H6.207Z" clip-rule="evenodd" />
                    <path fill="#B9FFB3" d="M29.793 0A6.207 6.207 0 0 1 36 6.207v20.06c0 3.546-4.488 5.085-6.664 2.286l-6.805-8.754v10.615A5.586 5.586 0 0 1 16.945 36a.62.62 0 0 0 .62-.62v-19.2c0-3.547 4.488-5.086 6.665-2.286l6.805 8.753V1.241C31.035.556 30.479 0 29.793 0Z" />
                    <path fill="#fff" d="M60.686 10.6v9.416l-9.13-9.416h-4.752V26h4.334V15.88L61.082 26h3.938V10.6h-4.334ZM73.054 22.568V19.84h9.702v-3.278h-9.702v-2.53h11.77V10.6H68.632V26h16.434v-3.432H73.054ZM96.583 26.506c6.094 0 10.054-2.992 10.054-8.206 0-5.214-3.96-8.206-10.054-8.206S86.551 13.086 86.551 18.3c0 5.214 3.938 8.206 10.032 8.206Zm0-3.652c-3.388 0-5.478-1.65-5.478-4.554 0-2.904 2.112-4.554 5.478-4.554 3.388 0 5.478 1.65 5.478 4.554 0 2.904-2.09 4.554-5.478 4.554ZM123.249 10.6v9.416l-9.13-9.416h-4.752V26h4.334V15.88L123.645 26h3.938V10.6h-4.334Z" />
                    <defs>
                        <linearGradient id="a" x1="36" x2="4.345" y1="36" y2="0" gradientUnits="userSpaceOnUse"><stop stop-color="#B9FFB3" /><stop offset="1" stop-color="#B9FFB3" stop-opacity="0" /></linearGradient>
                        <linearGradient id="b" x1="36" x2="14.617" y1="36" y2="27.683" gradientUnits="userSpaceOnUse"><stop stop-color="#1A1A1A" stop-opacity=".9" /><stop offset="1" stop-color="#1A1A1A" stop-opacity="0" /></linearGradient>
                    </defs>
                </svg>)
        case 'azure':
            return <svg {...props} xmlns="http://www.w3.org/2000/svg" width={size} height={size} viewBox="0 0 24 24" stroke-width="1.5" stroke="#2c3e50" fill="none" stroke-linecap="round" stroke-linejoin="round">
                <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                <path d="M6 7.5l-4 9.5h4l6 -15z" />
                <path d="M22 20l-7 -15l-3 7l4 5l-8 3z" />
            </svg>
        case 'aws':
            return (<svg {...props} width={size} height={size} xmlns="http://www.w3.org/2000/svg" id="integration-logos" viewBox="0 0 300 250"><g id="aws-logo"><path d="M109.864,136.2a18.523,18.523,0,0,0,.673,5.445,32.9,32.9,0,0,0,1.958,4.405,2.656,2.656,0,0,1,.428,1.407,2.422,2.422,0,0,1-1.162,1.835l-3.854,2.57a2.936,2.936,0,0,1-1.591.55,2.814,2.814,0,0,1-1.835-.856,18.926,18.926,0,0,1-2.2-2.876c-.612-1.04-1.223-2.2-1.9-3.609A22.6,22.6,0,0,1,82.4,153.509c-5.139,0-9.238-1.468-12.236-4.4s-4.527-6.852-4.527-11.747a15.75,15.75,0,0,1,5.567-12.6c3.732-3.181,8.688-4.772,14.989-4.772a48.431,48.431,0,0,1,6.485.49c2.264.306,4.589.8,7.036,1.346v-4.466c0-4.65-.979-7.892-2.876-9.789-1.957-1.9-5.261-2.814-9.972-2.814a27.81,27.81,0,0,0-6.607.8,48.9,48.9,0,0,0-6.608,2.08,17.482,17.482,0,0,1-2.141.8,3.727,3.727,0,0,1-.979.184c-.856,0-1.285-.612-1.285-1.9v-3a3.874,3.874,0,0,1,.429-2.142,4.577,4.577,0,0,1,1.713-1.284,35.178,35.178,0,0,1,7.708-2.753,37.063,37.063,0,0,1,9.544-1.163c7.281,0,12.6,1.652,16.029,4.956,3.365,3.3,5.078,8.32,5.078,15.05V136.2Zm-24.838,9.3a19.826,19.826,0,0,0,6.3-1.1,13.625,13.625,0,0,0,5.812-3.916,9.706,9.706,0,0,0,2.08-3.915,21.969,21.969,0,0,0,.612-5.323v-2.569a50.765,50.765,0,0,0-5.628-1.04,46.057,46.057,0,0,0-5.751-.367c-4.1,0-7.1.8-9.116,2.447a8.526,8.526,0,0,0-3,7.035c0,2.876.734,5.017,2.264,6.485C80.07,144.761,82.211,145.5,85.026,145.5Zm49.126,6.607a3.414,3.414,0,0,1-2.324-.611,4.988,4.988,0,0,1-1.285-2.386l-14.377-47.292a10.725,10.725,0,0,1-.551-2.447,1.343,1.343,0,0,1,1.469-1.53h5.995a3.393,3.393,0,0,1,2.386.612,5.01,5.01,0,0,1,1.224,2.386l10.278,40.5,9.544-40.5a4.3,4.3,0,0,1,1.162-2.386,4.21,4.21,0,0,1,2.447-.612h4.895a3.684,3.684,0,0,1,2.447.612,4.128,4.128,0,0,1,1.162,2.386l9.666,40.99,10.584-40.99a5.251,5.251,0,0,1,1.224-2.386,3.992,3.992,0,0,1,2.386-.612h5.69a1.36,1.36,0,0,1,1.529,1.53,6.116,6.116,0,0,1-.122.979,8.679,8.679,0,0,1-.429,1.529l-14.744,47.292a4.7,4.7,0,0,1-1.284,2.386,3.916,3.916,0,0,1-2.325.612h-5.262a3.684,3.684,0,0,1-2.447-.612,4.346,4.346,0,0,1-1.162-2.447l-9.483-39.461-9.422,39.4a4.776,4.776,0,0,1-1.162,2.448,3.793,3.793,0,0,1-2.447.611Zm78.616,1.652a40.493,40.493,0,0,1-9.422-1.1,27.862,27.862,0,0,1-7.036-2.447,4.419,4.419,0,0,1-1.9-1.713,4.317,4.317,0,0,1-.367-1.713v-3.12c0-1.285.489-1.9,1.407-1.9a3.446,3.446,0,0,1,1.1.184c.367.122.918.367,1.53.611a33.286,33.286,0,0,0,6.729,2.142,36.78,36.78,0,0,0,7.281.734c3.854,0,6.852-.673,8.932-2.019a6.588,6.588,0,0,0,3.181-5.812,5.957,5.957,0,0,0-1.652-4.283c-1.1-1.162-3.181-2.2-6.179-3.181l-8.871-2.753c-4.466-1.407-7.77-3.487-9.788-6.24a14.576,14.576,0,0,1-3.059-8.871,13.618,13.618,0,0,1,1.651-6.791,15.744,15.744,0,0,1,4.405-5.017,19.426,19.426,0,0,1,6.363-3.181,26.671,26.671,0,0,1,7.709-1.04,30.646,30.646,0,0,1,4.1.245c1.407.183,2.691.428,3.976.673,1.224.3,2.386.611,3.487.978a13.338,13.338,0,0,1,2.57,1.1,5.283,5.283,0,0,1,1.835,1.529,3.3,3.3,0,0,1,.551,2.019v2.875c0,1.285-.49,1.958-1.407,1.958a6.369,6.369,0,0,1-2.325-.734,27.983,27.983,0,0,0-11.746-2.386c-3.488,0-6.241.551-8.137,1.713a5.907,5.907,0,0,0-2.876,5.445,5.778,5.778,0,0,0,1.836,4.344c1.223,1.162,3.487,2.324,6.729,3.365l8.688,2.753c4.405,1.407,7.586,3.364,9.483,5.873a13.847,13.847,0,0,1,2.814,8.565,15.728,15.728,0,0,1-1.591,7.1,16.44,16.44,0,0,1-4.466,5.383,19.7,19.7,0,0,1-6.791,3.426A29.073,29.073,0,0,1,212.768,153.754Z" fill="#252f3e" /><path d="M224.33,183.487c-20.127,14.867-49.371,22.759-74.516,22.759a134.8,134.8,0,0,1-90.973-34.689c-1.9-1.713-.184-4.038,2.08-2.692a183.471,183.471,0,0,0,91.035,24.166,181.66,181.66,0,0,0,69.438-14.193C224.759,177.308,227.634,181.04,224.33,183.487Z" fill="#f90" fill-rule="evenodd" /><path d="M232.712,173.943c-2.57-3.3-17.008-1.59-23.554-.8-1.958.245-2.264-1.468-.489-2.753,11.5-8.076,30.406-5.751,32.608-3.059,2.2,2.753-.612,21.657-11.379,30.712-1.652,1.407-3.243.673-2.509-1.163C229.837,190.829,235.282,177.186,232.712,173.943Z" fill="#f90" fill-rule="evenodd" /></g></svg>)
        case 'golang':
            return <svg {...props} xmlns="http://www.w3.org/2000/svg" version="1.1" id="Layer_1" x="0px" y="0px" viewBox="0 0 205.4 76.7">
                <g>
                    <g>
                        <g>
                            <g>
                                <path d="M15.5,23.2c-0.4,0-0.5-0.2-0.3-0.5l2.1-2.7c0.2-0.3,0.7-0.5,1.1-0.5h35.7c0.4,0,0.5,0.3,0.3,0.6l-1.7,2.6      c-0.2,0.3-0.7,0.6-1,0.6L15.5,23.2z" />
                            </g>
                        </g>
                    </g>
                    <g>
                        <g>
                            <g>
                                <path d="M0.4,32.4c-0.4,0-0.5-0.2-0.3-0.5l2.1-2.7c0.2-0.3,0.7-0.5,1.1-0.5h45.6c0.4,0,0.6,0.3,0.5,0.6l-0.8,2.4      c-0.1,0.4-0.5,0.6-0.9,0.6L0.4,32.4z" />
                            </g>
                        </g>
                    </g>
                    <g>
                        <g>
                            <g>
                                <path d="M24.6,41.6c-0.4,0-0.5-0.3-0.3-0.6l1.4-2.5c0.2-0.3,0.6-0.6,1-0.6h20c0.4,0,0.6,0.3,0.6,0.7L47.1,41      c0,0.4-0.4,0.7-0.7,0.7L24.6,41.6z" />
                            </g>
                        </g>
                    </g>
                    <g>
                        <g id="CXHf1q_5_">
                            <g>
                                <g>
                                    <path d="M128.4,21.4c-6.3,1.6-10.6,2.8-16.8,4.4c-1.5,0.4-1.6,0.5-2.9-1c-1.5-1.7-2.6-2.8-4.7-3.8c-6.3-3.1-12.4-2.2-18.1,1.5       c-6.8,4.4-10.3,10.9-10.2,19c0.1,8,5.6,14.6,13.5,15.7c6.8,0.9,12.5-1.5,17-6.6c0.9-1.1,1.7-2.3,2.7-3.7c-3.6,0-8.1,0-19.3,0       c-2.1,0-2.6-1.3-1.9-3c1.3-3.1,3.7-8.3,5.1-10.9c0.3-0.6,1-1.6,2.5-1.6c5.1,0,23.9,0,36.4,0c-0.2,2.7-0.2,5.4-0.6,8.1       c-1.1,7.2-3.8,13.8-8.2,19.6c-7.2,9.5-16.6,15.4-28.5,17c-9.8,1.3-18.9-0.6-26.9-6.6c-7.4-5.6-11.6-13-12.7-22.2       c-1.3-10.9,1.9-20.7,8.5-29.3c7.1-9.3,16.5-15.2,28-17.3c9.4-1.7,18.4-0.6,26.5,4.9c5.3,3.5,9.1,8.3,11.6,14.1       C130,20.6,129.6,21.1,128.4,21.4z" />
                                </g>
                                <g>
                                    <path d="M161.5,76.7c-9.1-0.2-17.4-2.8-24.4-8.8c-5.9-5.1-9.6-11.6-10.8-19.3c-1.8-11.3,1.3-21.3,8.1-30.2       c7.3-9.6,16.1-14.6,28-16.7c10.2-1.8,19.8-0.8,28.5,5.1c7.9,5.4,12.8,12.7,14.1,22.3c1.7,13.5-2.2,24.5-11.5,33.9       c-6.6,6.7-14.7,10.9-24,12.8C166.8,76.3,164.1,76.4,161.5,76.7z M185.3,36.3c-0.1-1.3-0.1-2.3-0.3-3.3       c-1.8-9.9-10.9-15.5-20.4-13.3c-9.3,2.1-15.3,8-17.5,17.4c-1.8,7.8,2,15.7,9.2,18.9c5.5,2.4,11,2.1,16.3-0.6       C180.5,51.3,184.8,44.9,185.3,36.3z" />
                                </g>
                            </g>
                        </g>
                    </g>
                </g>
            </svg>
        case 'google':
            return <svg {...props} xmlns="http://www.w3.org/2000/svg" height={size} viewBox="0 0 24 24" width={size}><path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" fill="#4285F4" /><path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" fill="#34A853" /><path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z" fill="#FBBC05" /><path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" fill="#EA4335" /><path d="M1 1h22v22H1z" fill="none" /></svg>
    }
}