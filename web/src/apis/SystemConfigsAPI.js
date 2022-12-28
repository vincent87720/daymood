import req from './https'

export const getSystemConfigs = () => req('get', '/systemConfigs')