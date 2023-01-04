import req from './https'

export const getTradingSettings = () => req('get', '/tradings')