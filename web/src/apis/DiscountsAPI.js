import req from './https'

export const getDiscounts = (params) => req('get', '/deliveryOrders/' + params.DeliveryOrderID + '/discounts')
export const postDiscount = (params) => req('post', '/discounts', params)
export const putDiscount = (params) => req('put', '/discounts/' + params.ID, params)
export const deleteDiscount = (params) => req('delete', '/discounts/' + params.ID)