import './App.css'

import { QueryClient, QueryClientProvider } from 'react-query'
import {
  RouterProvider,
} from "react-router-dom";
import { router } from './Router';

const queryClient = new QueryClient()

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
    </QueryClientProvider>
  )
}

export default App
