import axios from 'axios'

var instance = axios.create({
    baseURL: process.env.VUE_APP_BASE_URL
})

export default function(method, url, data=null){
    method = method.toLowerCase();
    
    if(method == 'get'){
        return instance.get(url, {params: data});
    }
    else if(method == 'post'){
        return instance.post(url, data)
    }
    else if(method == 'put'){
        return instance.put(url, data);
    }
    else if(method == 'delete'){
        return instance.delete(url, {params: data});
    }
    else{
        console.error('method undefined');
        return false;
    }
}