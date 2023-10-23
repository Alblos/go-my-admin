import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import { QueryClient, QueryClientProvider } from 'react-query'
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import Home from './home/home.tsx'
import Terminal from './terminal/terminal.tsx'
import Connections from './connections/connections.tsx'
import Metrics from './metrics/metrics.tsx'

const queryClient = new QueryClient()

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
  },
  {
    path: "/dashboard",
    element: <Connections />,
  },
  {
    path: "/terminal",
    element: <Terminal />
  },
  {
    path: "/metrics",
    element: <Metrics />
  },
  {
    path: "/*",
    element: <App />,
  }
])

ReactDOM.createRoot(document.getElementById('root')!).render(
  <QueryClientProvider client={queryClient}>
    <React.StrictMode>
      <RouterProvider router={router} />
    </React.StrictMode>
  </QueryClientProvider>,
)
