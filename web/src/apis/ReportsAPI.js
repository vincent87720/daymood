import req from './https'

export const getBalancesReports = () => req('get', '/reports/balances')