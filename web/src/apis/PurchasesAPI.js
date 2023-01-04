import req from './https'

export const getPurchases = () => req('get', '/purchases')
export const getPurchase = (params) => req('get', '/purchases/' + params.ID)
export const postPurchase = (params) => req('post', '/purchases', params)
export const putPurchase = (params) => req('put', '/purchases/' + params.ID, params)
export const deletePurchase = (params) => req('delete', '/purchases/' + params.ID)