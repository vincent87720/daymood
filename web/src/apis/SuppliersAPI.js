import req from './https'

export const getSuppliers = () => req('get', '/suppliers')
export const postSupplier = (params) => req('post', '/suppliers', params)
export const putSupplier = (params) => req('put', '/suppliers/' + params.ID, params)
export const deleteSupplier = (params) => req('delete', '/suppliers/' + params.ID)