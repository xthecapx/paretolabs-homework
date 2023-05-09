import './App.css'

import { QueryClient, QueryClientProvider } from 'react-query'
import Home from './page/Home'

const queryClient = new QueryClient()

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <Home />
    </QueryClientProvider>
  )
}

export default App
