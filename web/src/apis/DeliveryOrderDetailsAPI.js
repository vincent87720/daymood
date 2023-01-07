import req from './https'

export const getAllDeliveryOrderDetails = () => req('get', '/deliveryOrderDetails')
export const getDeliveryOrderDetails = (params) => req('get', '/deliveryOrders/' + params.DeliveryOrderID + '/deliveryOrderDetails')
export const postDeliveryOrderDetails = (params) => req('post', '/deliveryOrderDetails/multiple', params)
export const postDeliveryOrderDetail = (params) => req('post', '/deliveryOrderDetails', params)
export const putDeliveryOrderDetail = (params) => req('put', '/deliveryOrderDetails/' + params.ID, params)
export const deleteDeliveryOrderDetail = (params) => req('delete', '/deliveryOrderDetails/' + params.ID)