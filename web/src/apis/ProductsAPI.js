import req from './https'

export const getProducts = () => req('get', '/products')
export const postProducts = (params) => req('post', '/products/multiple', params)
export const postProduct = (params) => req('post', '/products', params)
export const putProduct = (params) => req('put', '/products/' + params.ID, params)
export const deleteProduct = (params) => req('delete', '/products/' + params.ID)
export const getProductPurchaseHistories = (params) => req('get', '/products/' + params.ID + '/purchaseHistories')
export const getProductDeliveryHistories = (params) => req('get', '/products/' + params.ID + '/deliveryHistories')