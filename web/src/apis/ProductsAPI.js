import req from './https'

export const getProducts = () => req('get', '/products')
export const postProduct = (params) => req('post', '/products', params)
export const putProduct = (params) => req('put', '/products/'+params.ID, params)
export const deleteProduct = (params) => req('delete', '/products/'+params.ID)