import axios from "axios";
export const API_URL = {}
const $api = axios.create({
    baseURL: API_URL,
    withCredentials: true
})

// TODO:  ${localStorage.getItem('token')} => заменить на наш токен
$api.interceptors.request.use((config) => {
    config.headers.Authorization = `Bearer ${localStorage.getItem('token')}`
    return config;
})

$api.interceptors.response.use((config) => {
    return config;
},async (error) => {
    const originalRequest = error.config;
    if (error.response.status === 401 && error.config && !error.config._isRetry) {
        originalRequest._isRetry = true;
        try {
            // TODO: рефреш токен, access token взять
            const response = await axios.get(`${API_URL}/auth/refresh`, {withCredentials: true})
            localStorage.setItem('token', response.data.accessToken);
            return $api.request(originalRequest);
        } catch (e) {
            console.log('NOT AUTHORIZED')
        }
    }
    throw error;
})
export default $api;