import req from "./https";

export const getTradingSettings = () => req("get", "/tradings");
export const putTradingSettings = (params) => req("put", "/tradings", params);
