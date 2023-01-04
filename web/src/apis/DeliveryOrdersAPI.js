import req from './https'

export const getDeliveryOrders = () => req('get', '/deliveryOrders')
export const getDeliveryOrder = (params) => req('get', '/deliveryOrders/' + params.ID)
export const postDeliveryOrder = (params) => req('post', '/deliveryOrders', params)
export const putDeliveryOrder = (params) => req('put', '/deliveryOrders/' + params.ID, params)
export const deleteDeliveryOrder = (params) => req('delete', '/deliveryOrders/' + params.ID)