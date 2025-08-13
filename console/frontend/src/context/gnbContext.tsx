import { createContext, useContext, useState } from 'react'
import type { ReactNode } from 'react'

import type { ApiConsoleGnbRegistrationPost200Response } from '../api'

interface GnbConnection {
  ip: string
  port: number
}

interface GnbWithConnection extends ApiConsoleGnbRegistrationPost200Response {
  connection?: GnbConnection
}

interface GnbContextType {
  gnbList: GnbWithConnection[]
  addGnb: (gnb: ApiConsoleGnbRegistrationPost200Response, connection: GnbConnection) => void
  removeGnb: (gnbId: string) => void
}

const GnbContext = createContext<GnbContextType | undefined>(undefined)

export function GnbProvider({ children }: { children: ReactNode }) {
  const [gnbList, setGnbList] = useState<GnbWithConnection[]>([])

  const addGnb = (gnb: ApiConsoleGnbRegistrationPost200Response, connection: GnbConnection) => {
    setGnbList(prev => [...prev, { ...gnb, connection }])
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
