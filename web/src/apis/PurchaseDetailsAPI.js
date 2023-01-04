import req from './https'

export const getAllPurchaseDetails = () => req('get', '/purchaseDetails')
export const getPurchaseDetails = (params) => req('get', '/purchases/' + params.PurchaseID + '/purchaseDetails')
export const postPurchaseDetails = (params) => req('post', '/purchaseDetails/multiple', params)
export const postPurchaseDetail = (params) => req('post', '/purchaseDetails', params)
export const putPurchaseDetail = (params) => req('put', '/purchaseDetails/' + params.ID, params)
export const deletePurchaseDetail = (params) => req('delete', '/purchaseDetails/' + params.ID)