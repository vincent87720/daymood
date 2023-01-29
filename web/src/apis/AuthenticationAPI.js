import req from './https'

export const login = (params) => req('post', '/login', params)
export const logout = (params) => req('post', '/logout', params)