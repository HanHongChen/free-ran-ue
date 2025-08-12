import { createContext, useContext, useState } from 'react'
import type { ReactNode } from 'react'

import type { ApiConsoleGnbRegistrationPost200Response } from '../api'

interface GnbContextType {
  gnbList: ApiConsoleGnbRegistrationPost200Response[]
  addGnb: (gnb: ApiConsoleGnbRegistrationPost200Response) => void
  removeGnb: (gnbId: string) => void
}

const GnbContext = createContext<GnbContextType | undefined>(undefined)

export function GnbProvider({ children }: { children: ReactNode }) {
  const [gnbList, setGnbList] = useState<ApiConsoleGnbRegistrationPost200Response[]>([])

  const addGnb = (gnb: ApiConsoleGnbRegistrationPost200Response) => {
    setGnbList(prev => [...prev, gnb])
  }

  const removeGnb = (gnbId: string) => {
    setGnbList(prev => prev.filter(gnb => gnb.gnbInfo?.gnbId !== gnbId))
  }

  return (
    <GnbContext.Provider value={{ gnbList, addGnb, removeGnb }}>
      {children}
    </GnbContext.Provider>
  )
}

export function useGnb() {
  const context = useContext(GnbContext)
  if (context === undefined) {
    throw new Error('useGnb must be used within a GnbProvider')
  }
  return context
}
